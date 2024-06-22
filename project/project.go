package project

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/device"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name"`
	Disabled bool               `json:"disabled"`

	running bool

	devices []*device.Device

	watchers map[device.Watcher]any
}

func (p *Project) Open() error {
	p.watchers = make(map[device.Watcher]any)

	//todo 订阅设备 find ids watch

	p.running = true
	return nil
}

func (p *Project) Close() error {
	p.running = false
	return nil
}

func (p *Project) Execute(actions []*base.Action) {
	if !p.running {
		return
	}

}

func (p *Project) Devices() []*device.Device {
	return p.devices
}

func (p *Project) Watch(w device.Watcher) {
	p.watchers[w] = 1
}
