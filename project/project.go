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

	devices []*device.Device

	watchers map[Watcher]any
}

func (p *Project) Open() error {
	p.watchers = make(map[Watcher]any)

	//todo 加载场景，报警，定时

	return nil
}

func (p *Project) Close() error {

	return nil
}

func (p *Project) Execute(actions []*base.Action) {

}

func (p *Project) Devices() []*device.Device {
	return p.devices
}

func (p *Project) Watch(w Watcher) {
	p.watchers[w] = 1
}
