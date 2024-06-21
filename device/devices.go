package device

import (
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var devices lib.Map[Device]

func Get(id string) *Device {
	return devices.Load(id)
}

func From(v *Device) (err error) {
	tt := devices.LoadAndStore(v.Id.Hex(), v)
	if tt != nil {
		_ = tt.Close()
	}
	return v.Open()
}

func Load(id primitive.ObjectID) error {
	var device Device
	err := _table.Get(id, &device)
	if err != nil {
		return err
	}
	return From(&device)
}

func Unload(id primitive.ObjectID) error {
	t := devices.LoadAndDelete(id.Hex())
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Device](&_table, nil, 100, func(t *Device) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}
