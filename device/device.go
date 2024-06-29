package device

import (
	"github.com/god-jason/bucket/action"
	"github.com/god-jason/bucket/aggregate/aggregator"
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/product"
	"github.com/god-jason/bucket/project"
	"github.com/god-jason/bucket/space"
	"github.com/mochi-mqtt/server/v2"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Aggregator struct {
	aggregator.Aggregator
	As string
}

type Device struct {
	Id        string `json:"_id" bson:"_id"`
	ProductId string `json:"product_id" bson:"product_id"`
	ProjectId string `json:"project_id,omitempty" bson:"project_id"`
	SpaceId   string `json:"space_id,omitempty" bson:"space_id"`
	Name      string `json:"name"`
	Disabled  bool   `json:"disabled"`

	running bool

	//产品
	product *product.Product

	//变量
	values map[string]any

	//聚合器
	aggregators map[string]*Aggregator

	//等待的操作响应 todo 加锁
	pendingActions map[string]chan *PayloadActionUp

	//网关连接
	gatewayClient *mqtt.Client

	//监听
	valuesWatchers map[base.ValuesWatcher]any
}

func (d *Device) Open() error {

	d.product = product.Get(d.ProductId)
	if d.product == nil {
		return exception.New("找不到产品" + d.ProductId)
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

	d.pendingActions = make(map[string]chan *PayloadActionUp)

	d.valuesWatchers = make(map[base.ValuesWatcher]any)

	d.running = true

	//找到项目，空间，主动汇报数据
	if d.ProjectId != "" {
		prj := project.Get(d.ProjectId)
		if prj != nil {
			d.WatchValues(prj)
		} else {
			//return errors.New("找不到项目")
		}
	}
	if d.SpaceId != "" {
		spc := space.Get(d.SpaceId)
		if spc != nil {
			d.WatchValues(spc)
		} else {
			//return errors.New("找不到空间")
		}
	}

	return nil
}

func (d *Device) Close() error {
	d.running = false
	d.pendingActions = nil
	d.valuesWatchers = nil
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
			values["product_id"] = d.ProductId
			values["project_id"] = d.ProjectId
			values["space_id"] = d.SpaceId
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
	pid, id := d.ProductId, d.Id
	for w, _ := range d.valuesWatchers {
		w.OnValuesChange(pid, id, d.values)
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
				return exception.New(p.Label + " 不能写入")
			}
		} else {
			return exception.New("未知的属性：" + k)
		}
	}

	//向网关发送写指令
	if d.gatewayClient != nil {
		return publishDirectly(d.gatewayClient, "down/device/"+d.Id+"/property", values)
	}

	return nil
}

func (d *Device) Action(name string, values map[string]any) (map[string]any, error) {

	act := map[string]any{
		"product_id": d.ProductId,
		"device_id":  d.Id,
		"project_id": d.ProjectId,
		"space_id":   d.SpaceId,
		"name":       name,
		"parameters": values,
	}

	id, err := action.Table().Insert(&act)
	if err != nil {
		return nil, err
	}

	//检查参数

	//向网关发送写指令
	if d.gatewayClient != nil && !d.gatewayClient.Closed() {
		payload := PayloadActionDown{Id: id, Name: name, Parameters: values}
		err := publishDirectly(d.gatewayClient, "down/device/"+d.Id+"/name", &payload)
		if err != nil {
			return nil, err
		}
		d.pendingActions[id] = make(chan *PayloadActionUp)

		//等待结果
		select {
		case <-time.After(time.Minute):
			_ = action.Table().Update(id, bson.M{"result": "timeout"})
			return nil, exception.New("超时")
		case ret := <-d.pendingActions[name]:
			_ = action.Table().Update(id, bson.M{"return": ret.Return, "result": ret.Result})
			return ret.Return, nil
		}
	}

	_ = action.Table().Update(id, bson.M{"result": "offline"})
	return nil, exception.New("不可到达")
}

func (d *Device) Values() map[string]any {
	return d.values
}

func (d *Device) WatchValues(watcher base.ValuesWatcher) {
	d.valuesWatchers[watcher] = 1
}

func (d *Device) UnWatchValues(watcher base.ValuesWatcher) {
	delete(d.valuesWatchers, watcher)
}
