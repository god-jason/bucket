package product

import "github.com/god-jason/bucket/table"

var _table = table.Table{
	Name:   Bucket,
	Schema: nil,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "icon", Label: "图标", Type: "string"},
		{Name: "properties", Label: "属性", Type: "array"},
		{Name: "actions", Label: "操作", Type: "array"},
		{Name: "events", Label: "事件", Type: "array"},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

func Table() *table.Table {
	return &_table
}
