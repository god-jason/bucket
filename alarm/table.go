package alarm

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _validatorTable = table.Table{
	Name: base.BucketValidator,
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
	},
}

var _alarmTable = table.Table{
	Name: base.BucketAlarm,
	Fields: []*table.Field{
		base.ProjectIdField,
		base.SpaceIdField,
		base.ProductIdField,
		base.DeviceIdField,
		{Name: "title", Label: "标题", Type: "string", Required: true},
		{Name: "type", Label: "类型", Type: "string", Required: true},
		{Name: "level", Label: "等级", Type: "number", Required: true},
		{Name: "message", Label: "消息", Type: "string", Required: true},
		{Name: "created", Label: "日期", Type: "date"},
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
	_validatorTable.Hook = &_hook
	table.Register(&_validatorTable)
	table.Register(&_alarmTable)
}
