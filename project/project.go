package project

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/pool"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	Id       string `json:"_id" bson:"_id"`
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`

	running bool

	valuesWatchers map[base.ProjectValuesWatcher]any
}

func (p *Project) Open() error {
	//todo 启动所有空间

	p.valuesWatchers = make(map[base.ProjectValuesWatcher]any)
	p.running = true
	return nil
}

func (p *Project) Close() error {
	//todo 关闭所有空间

	p.valuesWatchers = nil
	p.running = false
	return nil
}

func (p *Project) Devices(productId string) (ids []string, err error) {
	if !p.running {
		return nil, exception.New("项目已经关闭")
	}
	deviceTable, _ := table.Get(base.BucketDevice)
	return deviceTable.DistinctId(bson.D{
		{"project_id", p.Id},
		{"product_id", productId},
	})
}

func (p *Project) OnDeviceValuesChange(product, device string, values map[string]any) {
	for w, _ := range p.valuesWatchers {
		_ = pool.Insert(func() {
			w.OnProjectValuesChange(p.Id, product, device, values)
		})
	}
}

func (p *Project) WatchValues(w base.ProjectValuesWatcher) {
	p.valuesWatchers[w] = 1
}

func (p *Project) UnWatchValues(w base.ProjectValuesWatcher) {
	delete(p.valuesWatchers, w)
}
