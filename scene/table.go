package scene

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name: base.BucketScene,
	Fields: []*table.Field{
		{Name: "project_id", Label: "项目", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
			Table: "bucket.project",
			Field: "_id",
			As:    "project",
		}},
		{Name: "space_id", Label: "空间", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
			Table: "bucket.space",
			Field: "_id",
			As:    "space",
		}},
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "times", Label: "时间限制", Type: "array", Children: []*table.Field{
			{Name: "start", Label: "开始", Type: "number", Required: true},
			{Name: "end", Label: "结束", Type: "number", Required: true},
			{Name: "weekday", Label: "星期", Type: "array", Required: true},
		}},
		//condition
		{Name: "conditions", Label: "时间限制", Type: "array"},
		{Name: "disabled", Label: "禁用", Type: "boolean"},
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
	_table.Hook = &_hook
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
