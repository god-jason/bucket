package table

import (
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/mongodb"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"path/filepath"
)

const Path = "tables"

var tables lib.Map[Table]

func Get(name string) (*Table, error) {
	table := tables.Load(name)
	if table == nil {
		return nil, exception.New("没有表定义 " + name)
	}
	return table, nil
}

func Register(table *Table) {
	tables.Store(table.Name, table)
}

func LoadAll() error {
	d := filepath.Join(viper.GetString("data"), Path)
	_ = os.MkdirAll(d, os.ModePerm)

	es, err := os.ReadDir(d)
	if err != nil {
		return err
	}

	for _, e := range es {
		if e.IsDir() {
			err = Load(e.Name(), false)
			if err != nil {
				log.Error(err)
				//return err
			}
			continue
		}
	}

	return nil
}

func Sync() error {
	tabs, err := mongodb.Tables()
	if err != nil {
		return err
	}

	//这里锁了，合适不???
	tables.Range(func(name string, table *Table) bool {
		//log.Println("table sync", name)
		for _, t := range tabs {
			if t == name {
				//todo 检查索引
				//log.Println("table sync", name, "skip")
				return true
			}
		}
		log.Println("table sync", name)

		//创建表
		opts := options.CreateCollection()
		if table.TimeSeries != nil {
			//时序参数
			opts.SetTimeSeriesOptions(table.TimeSeries)
		}
		err = mongodb.CreateTable(name, opts)
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
			err = mongodb.CreateIndex(name, keys)
			if err != nil {
				return false
			}
		}

		log.Println("table sync", name, "finished")

		return true
	})

	return err
}
