package device

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _table = table.Table{
	Name: base.BucketDevice,
	Fields: []*table.Field{
		base.ProductIdField,
		base.ProjectIdField,
		base.SpaceIdField,
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "disabled", Label: "禁用", Type: "boolean"},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

var _hook = table.Hook{
	AfterInsert: func(id primitive.ObjectID, doc any) error {
		return Load(id)
	},
	AfterUpdate: func(id primitive.ObjectID, doc any) error {
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
