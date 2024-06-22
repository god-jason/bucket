package history

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _tags = []*table.Field{
	{Name: "device_id", Label: "设备", Type: "string", Foreign: &table.Foreign{
		Table: "device",
		Field: "_id",
		As:    "device",
	}},
	{Name: "product_id", Label: "产品", Type: "string", Foreign: &table.Foreign{
		Table: "product",
		Field: "_id",
		As:    "product",
	}},
}

var _table = table.Table{
	Name: base.BucketHistory,
	Fields: []*table.Field{
		{Name: "tags", Label: "标签", Type: "object", Children: _tags},
		{Name: "date", Label: "日期", Type: "date"},
	},
	TimeSeries: options.TimeSeries().
		SetTimeField("date").
		SetMetaField("tags").
		SetGranularity("minutes"), //默认按分钟存储
}

func init() {
	table.Register(&_table)
}

func Table() *table.Table {
	return &_table
}
