package device

import (
	"errors"
	"github.com/god-jason/bucket/aggregate"
	"github.com/god-jason/bucket/aggregate/aggregator"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/history"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/product"
	"github.com/god-jason/bucket/table"
	mqtt "github.com/mochi-mqtt/server/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var devices lib.Map[Device]

var aggregateStore = db.Batch{
	Collection:   aggregate.Bucket,
	WriteTimeout: time.Second,
	BufferSize:   200,
}

var historyStore = db.Batch{
	Collection:   history.Bucket,
	WriteTimeout: time.Second,
	BufferSize:   200,
}

type Aggregator struct {
	aggregator.Aggregator
	As string
}

func Get(id string) *Device {
	return devices.Load(id)
}

func Load(doc table.Document) (err error) {
	dev := new(Device)
	if id, ok := doc["_id"]; !ok {
		if dev.id, err = db.ParseObjectId(id); err != nil {
			return errors.New("_id 类型不正确")
		}
	} else {
		return errors.New("缺少 _id")
	}

	if id, ok := doc["product_id"]; !ok {
		if dev.productId, err = db.ParseObjectId(id); err != nil {
			return errors.New("product_id 类型不正确")
		}
	} else {
		return errors.New("缺少 product_id")
	}

	devices.Store(dev.ID(), dev)

	return dev.Open()
}

type Device struct {
	id primitive.ObjectID

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
}

func (d *Device) ID() string {
	return d.id.Hex()
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

	return nil
}

func (d *Device) Close() error {
	devices.Delete(d.id.Hex())
	return nil
}

func (d *Device) snap() {
	for _, agg := range d.aggregators {
		agg.Snap()
	}
}

func (d *Device) aggregate() {
	var values map[string]any
	for _, a := range d.aggregators {
		val := a.Pop()
		if val == nil {
			values[a.As] = val
		}
	}

	if len(values) > 0 {
		values["device_id"] = d.id
		//写入数据库，batch
		aggregateStore.InsertOne(values)
	}
}

func (d *Device) PatchValues(values map[string]any) {
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
		his["device_id"] = d.id
		his["date"] = time.Now()
		historyStore.InsertOne(his)
	}
}

func (d *Device) WriteHistory(history map[string]any, timestamp int64) {
	history["device_id"] = d.id
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
		return publishDirectly(d.gatewayClient, "down/device/"+d.id.Hex()+"/property", values)
	}

	return nil
}

func (d *Device) Action(action string, values map[string]any) (map[string]any, error) {

	//检查参数

	//向网关发送写指令
	if d.gatewayClient != nil {
		payload := PayloadAction{Action: action, Values: values}
		err := publishDirectly(d.gatewayClient, "down/device/"+d.id.Hex()+"/action", &payload)
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
