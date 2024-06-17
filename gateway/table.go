package gateway

import "github.com/god-jason/bucket/table"

var _table = table.Table{
	Name:   Bucket,
	Schema: nil,
	Fields: []*table.Field{
		{Name: "product_id", Label: "产品", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
			Table: "product",
			Field: "_id",
			As:    "product",
		}},
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "disabled", Label: "禁用", Type: "boolean"},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

func Table() *table.Table {
	return &_table
}
