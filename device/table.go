package device

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name:   base.BucketDevice,
	Schema: nil,
	Fields: []*table.Field{
		base.ProductIdField,
		base.ProjectIdField,
		base.SpaceIdField,
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "disabled", Label: "禁用", Type: "boolean"},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

func init() {
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
