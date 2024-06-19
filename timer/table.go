package timer

import (
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name:   db.BucketProject,
	Schema: nil,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "created", Label: "创建日期", Type: "date"},
	},
}

func init() {
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
