package device

import (
	"github.com/god-jason/bucket/aggregator"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/product"
	"time"
)

const table = "bucket.device"

var devices lib.Map[Device]

var aggregateStore = db.Batch{
	Collection:   "bucket.aggregate",
	WriteTimeout: time.Second,
	BufferSize:   200,
}

var historyStore = db.Batch{
	Collection:   "bucket.history",
	WriteTimeout: time.Second,
	BufferSize:   200,
}

func Get(id string) *Device {
	return devices.Load(id)
}

type Device struct {
	id string

	properties map[string]*product.Property

	historical  map[string]bool
	values      map[string]any
	aggregators map[string]aggregator.Aggregator
}

func (d *Device) ID() string {
	return d.id
}

func (d *Device) Snap() {
	for _, a := range d.aggregators {
		a.Snap()
	}
}

func (d *Device) Aggregate() {
	var values map[string]any
	for f, a := range d.aggregators {
		val := a.Pop()
		if val == nil {
			values[f] = val
		}
	}

	if len(values) > 0 {
		values["device_id"] = d.id
		//写入数据库，batch
		aggregateStore.InsertOne(values)
	}
}

func (d *Device) PatchValues(values map[string]any) {
	history := make(map[string]any)

	for k, v := range values {
		d.values[k] = v

		if p, ok := d.properties[k]; ok {
			//保存历史
			if p.Historical {
				history[k] = v
			}
		}

		//聚合计算
		if a, ok := d.aggregators[k]; ok {
			_ = a.Push(v)
		}
	}

	//保存历史
	if len(history) > 0 {
		history["device_id"] = d.id
		history["date"] = time.Now()
		historyStore.InsertOne(history)
	}
}

func (d *Device) WriteHistory(history map[string]any, timestamp int64) {
	history["device_id"] = d.id
	history["date"] = time.UnixMilli(timestamp)
	historyStore.InsertOne(history)
}
