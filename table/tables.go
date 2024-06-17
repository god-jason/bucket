package table

import (
	"errors"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrTableNotFound = errors.New("没有表定义")

var tables lib.Map[Table]

func Get(name string) (*Table, error) {
	table := tables.Load(name)
	if table == nil {
		return nil, ErrTableNotFound
	}
	return table, nil
}

func Register(table *Table) {
	tables.Store(table.Name, table)
}

func Sync() error {
	tabs, err := db.Tables()
	if err != nil {
		return err
	}

	//todo 这里锁了，合适不
	tables.Range(func(name string, table *Table) bool {
		for _, t := range tabs {
			if t == name {
				//todo 检查索引
				return true
			}
		}

		//创建表
		opts := options.CreateCollection()
		if table.TimeSeries != nil {
			//时序参数
			opts.SetTimeSeriesOptions(table.TimeSeries)
		}
		err = db.CreateTable(name, opts)
		if err != nil {
			return false
		}

		//创建索引
		var keys []string
		for _, f := range table.Fields {
			if f.Index {
				keys = append(keys, f.Name)
			}
		}
		if len(keys) > 0 {
			err = db.CreateIndex(name, keys)
			if err != nil {
				return false
			}
		}

		return true
	})

	return err
}
