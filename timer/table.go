package timer

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _table = table.Table{
	Name: base.BucketTimer,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

var _hook = table.Hook{
	AfterInsert: func(id primitive.ObjectID, doc any) error {
		return Load(id)
	},
	AfterUpdate: func(id primitive.ObjectID, update any, base db.Document) error {
		return Load(id)
	},
	AfterDelete: func(id primitive.ObjectID, doc db.Document) error {
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
