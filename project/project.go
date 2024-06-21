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
}

func (p *Project) Open() error {
	return nil
}

func (p *Project) Close() error {
	return nil
}

func (p *Project) Action() error {
	return nil
}

func (p *Project) Execute(actions []*base.Action) {

}

func (p *Project) Devices() []*device.Device {
	return p.devices
}
