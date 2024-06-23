package user

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/table"
)

var _table = table.Table{
	Name: base.BucketTimer,
	Fields: []*table.Field{
		{Name: "name", Label: "名称", Type: "string", Required: true},
		{Name: "username", Label: "用户名", Type: "string", Required: true},
		{Name: "created", Label: "创建日期", Type: "date", Created: true},
	},
}

var _passwordTable = table.Table{
	Name: base.BucketTimer,
	Fields: []*table.Field{
		{Name: "password", Label: "密码", Type: "string", Required: true},
	},
}

var _roleTable = table.Table{
	Name: base.BucketTimer,
	Fields: []*table.Field{
		base.UserIdField,
		{Name: "privileges", Label: "权限", Type: "array"},
	},
}

func init() {
	table.Register(&_table)
	table.Register(&_passwordTable)
}

func Table() *table.Table {
	return &_table
}
