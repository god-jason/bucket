package history

import "github.com/god-jason/bucket/table"

const Bucket = "bucket.history"

var _table = table.Table{
	Name:   Bucket,
	Schema: nil,
	Fields: []*table.Field{
		{Name: "device_id", Label: "设备", Type: "string", Index: true, Required: true, Foreign: &table.Foreign{
			Table: "device",
			Field: "_id",
			As:    "device",
		}},
		{Name: "date", Label: "日期", Type: "date"},
		//{Name: "created", Label: "创建日期", Type: "date"},
	},
	TimeSeries: &table.TimeSeries{
		TimeField:   "date",
		MetaField:   "",
		Granularity: "seconds",
	},
}

func Table() *table.Table {
	return &_table
}
