package device

import (
	"github.com/god-jason/bucket/aggregate/aggregator"
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/product"
	"github.com/mochi-mqtt/server/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Aggregator struct {
	aggregator.Aggregator
	As string
}

type Device struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProductId primitive.ObjectID `json:"product_id" bson:"product_id"`
	ProjectId primitive.ObjectID `json:"project_id,omitempty" bson:"project_id"`
	SpaceId   primitive.ObjectID `json:"space_id,omitempty" bson:"space_id"`
	Name      string             `json:"name"`
	Disabled  bool               `json:"disabled"`

	running bool

	//产品
	productId primitive.ObjectID
	product   *product.Product

	//变量
	values map[string]any

	//聚合器
	aggregators map[string]*Aggregator

	//等待的操作响应
	pendingActions map[string]chan map[string]any

	//网关连接
	gatewayClient *mqtt.Client

	//监听
	watchers map[base.DeviceValuesWatcher]any
}

func (d *Device) Open() error {

	d.product = product.Get(d.productId.Hex())
	if d.product == nil {
		return errors.New("找不到产品" + d.productId.Hex())
	}

	d.values = make(map[string]any)

	d.aggregators = make(map[string]*Aggregator)
	for _, p := range d.product.Properties {
		for _, a := range p.Aggregators {
			agg, err := aggregator.New(a.Type)
			if err != nil {
				return err
			}
			d.aggregators[p.Name] = &Aggregator{
				Aggregator: agg,
				As:         a.As,
			}
		}
	}

	d.pendingActions = make(map[string]chan map[string]any)

	d.watchers = make(map[base.DeviceValuesWatcher]any)

	d.running = true

	return nil
}

func (d *Device) Close() error {
	d.running = false
	d.pendingActions = nil
	d.watchers = nil
	d.aggregators = nil
	return nil
}

func (d *Device) snap() {
	if !d.running {
		return
	}
	for _, agg := range d.aggregators {
		agg.Snap()
	}
}

func (d *Device) aggregate(now time.Time) {
	if !d.running {
		return
	}

	if len(d.aggregators) > 0 {
		values := make(map[string]any)
		for _, a := range d.aggregators {
			val := a.Pop()
			if val != nil {
				values[a.As] = val
			}
		}

		if len(values) > 0 {

			values["device_id"] = d.Id
			values["date"] = now
			//写入数据库，batch
			aggregateStore.InsertOne(values)
		}
	}
}

func (d *Device) PatchValues(values map[string]any) {
	if !d.running {
		return
	}

	his := make(map[string]any)

	for k, v := range values {
		d.values[k] = v

		//检查字段
		p := d.product.GetProperty(k)
		if p != nil {
			//保存历史
			if p.Historical {
				his[k] = v
			}
		}

		//聚合计算
		if a, ok := d.aggregators[k]; ok {
			_ = a.Push(v)
		}
	}

	//保存历史
	if len(his) > 0 {
		his["device_id"] = d.Id
		his["date"] = time.Now()
		historyStore.InsertOne(his)
	}

	//监听变化
	for w, _ := range d.watchers {
		w.OnDeviceValuesChange(d.values)
	}
}

func (d *Device) WriteHistory(history map[string]any, timestamp int64) {
	history["device_id"] = d.Id
	history["date"] = time.UnixMilli(timestamp)
	historyStore.InsertOne(history)
}

func (d *Device) WriteValues(values map[string]any) error {

	//检查数据
	for k, _ := range values {
		p := d.product.GetProperty(k)
		if p != nil {
			if !p.Writable {
				return errors.New(p.Label + " 不能写入")
			}
		} else {
			return errors.New("未知的属性：" + k)
		}
	}

	//向网关发送写指令
	if d.gatewayClient != nil {
		return publishDirectly(d.gatewayClient, "down/device/"+d.Id.Hex()+"/property", values)
	}

	return nil
}

func (d *Device) Action(action string, values map[string]any) (map[string]any, error) {

	//检查参数

	//向网关发送写指令
	if d.gatewayClient != nil {
		payload := PayloadAction{Action: action, Values: values}
		err := publishDirectly(d.gatewayClient, "down/device/"+d.Id.Hex()+"/action", &payload)
		if err != nil {
			return nil, err
		}
		d.pendingActions[action] = make(chan map[string]any)

		//等待结果
		select {
		case <-time.After(time.Minute):
			return nil, errors.New("超时")
		case results := <-d.pendingActions[action]:
			return results, nil
		}
	}

	return nil, errors.New("不可到达")
}

func (d *Device) Values() map[string]any {
	return d.values
}

func (d *Device) Watch(watcher base.DeviceValuesWatcher) {
	d.watchers[watcher] = 1
}

func (d *Device) UnWatch(watcher base.DeviceValuesWatcher) {
	delete(d.watchers, watcher)
}
