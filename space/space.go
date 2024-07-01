package space

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/pool"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

type Space struct {
	Id        string `json:"_id" bson:"_id"`
	ProjectId string `json:"project_id" bson:"project_id"`
	Name      string `json:"name"`
	Disabled  bool   `json:"disabled"`

	running bool

	valuesWatchers map[base.SpaceValuesWatcher]any
}

func (s *Space) Open() error {
	s.valuesWatchers = make(map[base.SpaceValuesWatcher]any)
	s.running = true
	return nil
}

func (s *Space) Close() error {
	s.valuesWatchers = nil
	s.running = false
	return nil
}

func (s *Space) Devices(productId string) (ids []string, err error) {
	if !s.running {
		return nil, exception.New("空间已经关闭")
	}
	deviceTable, _ := table.Get(base.BucketDevice)
	return deviceTable.DistinctId(bson.D{
		{"space_id", s.Id},
		{"product_id", productId},
	})
}

func (s *Space) OnDeviceValuesChange(product, device string, values map[string]any) {
	for w, _ := range s.valuesWatchers {
		_ = pool.Insert(func() {
			w.OnSpaceValuesChange(s.Id, product, device, values)
		})
	}
}

func (s *Space) WatchValues(w base.SpaceValuesWatcher) {
	s.valuesWatchers[w] = 1
}

func (s *Space) UnWatchValues(w base.SpaceValuesWatcher) {
	delete(s.valuesWatchers, w)
}
