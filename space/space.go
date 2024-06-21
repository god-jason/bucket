package space

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/device"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Space struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	Name      string             `json:"name"`
	Disabled  bool               `json:"disabled"`

	devices []*device.Device

	watchers map[Watcher]any
}

func (s *Space) Open() error {
	s.watchers = make(map[Watcher]any)

	//todo 加载设备 加载场景
	return nil
}

func (s *Space) Close() error {
	return nil
}

func (s *Space) Execute(actions []*base.Action) {

}

func (s *Space) Devices() []*device.Device {
	return s.devices
}

func (s *Space) Watch(w Watcher) {
	s.watchers[w] = 1
}
