package space

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

var _table = table.Table{
	Name: base.BucketSpace,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

var _hook = table.Hook{
	AfterInsert: func(id string, doc any) error {
		return Load(id)
	},
	AfterUpdate: func(id string, update any, base db.Document) error {
		return Load(id)
	},
	AfterDelete: func(id string, doc db.Document) error {
		go func() {
			//删除相关报警
			v, e := table.Get(base.BucketValidator)
			if e == nil {
				ids, err := v.DistinctId(bson.D{{"space_id", id}})
				if err == nil {
					for _, id := range ids {
						_ = v.Delete(id)
					}
				}
			}

			//删除相关场景
			s, e := table.Get(base.BucketScene)
			if e == nil {
				ids, err := s.DistinctId(bson.D{{"space_id", id}})
				if err == nil {
					for _, id := range ids {
						_ = s.Delete(id)
					}
				}
			}

			//删除相关定时
			t, e := table.Get(base.BucketTimer)
			if e == nil {
				ids, err := t.DistinctId(bson.D{{"space_id", id}})
				if err == nil {
					for _, id := range ids {
						_ = s.Delete(id)
					}
				}
			}
		}()
		return Unload(id)
	},
}

func init() {
	_table.Hook = &_hook
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
