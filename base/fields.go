package base

import (
	"github.com/god-jason/bucket/table"
)

var ProductIdField = &table.Field{Name: "product_id", Label: "产品", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
	Table: BucketProduct,
	Field: "_id",
	Let:   "name",
	As:    "product",
}}

var DeviceIdField = &table.Field{Name: "device_id", Label: "设备", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
	Table: BucketDevice,
	Field: "_id",
	Let:   "name",
	As:    "device",
}}

var ProjectIdField = &table.Field{Name: "project_id", Label: "项目", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
	Table: BucketProject,
	Field: "_id",
	Let:   "name",
	As:    "project",
}}

var SpaceIdField = &table.Field{Name: "space_id", Label: "空间", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
	Table: BucketSpace,
	Field: "_id",
	Let:   "name",
	As:    "space",
}}

var UserIdField = &table.Field{Name: "user_id", Label: "用户", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
	Table: BucketSpace,
	Field: "_id",
	Let:   "name",
	As:    "user",
}}

var ActionFields = []*table.Field{
	ProductIdField,
	DeviceIdField,
	{Name: "name", Label: "名称", Type: "string", Required: true},
	{Name: "parameters", Label: "参数", Type: "object"},
}
