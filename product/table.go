package product

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name: base.BucketProduct,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "icon", Label: "图标", Type: "string"},
		{Name: "type", Label: "类型", Type: "string"},
		{Name: "properties", Label: "属性", Type: "array"},
		{Name: "actions", Label: "操作", Type: "array"},
		{Name: "events", Label: "事件", Type: "array"},
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
		return Unload(id)
	},
}

func init() {
	table.Register(&_table)

	_table.Hook = &_hook
}

func Table() *table.Table {
	return &_table
}
