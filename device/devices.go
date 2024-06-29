package device

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/table"
)

var devices lib.Map[Device]

func Get(id string) *Device {
	return devices.Load(id)
}

func From(v *Device) (err error) {
	tt := devices.LoadAndStore(v.Id, v)
	if tt != nil {
		_ = tt.Close()
	}
	if v.Disabled {
		return nil
	}
	return v.Open()
}

func Load(id string) error {
	var device Device
	has, err := _table.Get(id, &device)
	if err != nil {
		return err
	}
	if !has {
		return exception.New("找不到记录")
	}
	return From(&device)
}

func Unload(id string) error {
	t := devices.LoadAndDelete(id)
	if t != nil {
		return t.Close()
	}
	return nil
}

func Open(id string) error {
	dev := devices.Load(id)
	if dev == nil {
		return exception.New("找不到设备")
	}
	return dev.Open()
}

func Close(id string) error {
	dev := devices.Load(id)
	if dev == nil {
		return exception.New("找不到设备")
	}
	return dev.Close()
}

func LoadAll() error {
	return table.BatchLoad[*Device](&_table, base.FilterEnabled, 100, func(t *Device) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}
