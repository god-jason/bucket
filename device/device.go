package device

import (
	"github.com/god-jason/bucket/aggregator"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
	"time"
)

const table = "bucket.device"

var devices lib.Map[Device]

var aggregateStore = db.Batch{
	Collection:   "bucket.aggregate",
	WriteTimeout: time.Second,
	BufferSize:   200,
}

type Device struct {
	id string

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
