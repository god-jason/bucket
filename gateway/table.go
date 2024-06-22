package gateway

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name: base.BucketGateway,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string"},
		{Name: "username", Label: "用户名", Type: "string"},
		{Name: "password", Label: "密码", Type: "string"},
		{Name: "clientId", Label: "客户端ID", Type: "string", Index: true},
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
