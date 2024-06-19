package alarm

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name:   base.BucketAlarm,
	Schema: nil,
	Fields: []*table.Field{
		base.ProjectIdField,
		base.SpaceIdField,
		base.ProductIdField,
		base.DeviceIdField,
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "title", Label: "标题", Type: "string", Required: true},
		{Name: "type", Label: "类型", Type: "string", Required: true},
		{Name: "level", Label: "等级", Type: "number", Required: true},
		{Name: "message", Label: "消息", Type: "string", Required: true},
		{Name: "created", Label: "日期", Type: "date"},
		//{Name: "created", Label: "创建日期", Type: "date"},
	},
}

func init() {
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
