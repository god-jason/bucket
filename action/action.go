package action

import "github.com/god-jason/bucket/table"

type Action struct {
	Batch      bool              `json:"batch,omitempty"` //批量操作
	ProductId  string            `json:"product_id,omitempty" bson:"product_id"`
	DeviceId   string            `json:"device_id,omitempty" bson:"device_id"`
	Action     string            `json:"action"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

var Fields = []*table.Field{
	{Name: "batch", Label: "批量", Type: "bool"},
	{Name: "product_id", Label: "产品", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
		Table: "bucket.product",
		Field: "_id",
		As:    "product",
	}},
	{Name: "device_id", Label: "设备", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
		Table: "bucket.device",
		Field: "_id",
		As:    "device",
	}},
	{Name: "action", Label: "操作", Type: "string", Required: true},
	{Name: "parameters", Label: "参数", Type: "object"},
}
