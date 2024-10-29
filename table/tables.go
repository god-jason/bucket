package table

import (
	"encoding/json"
	"errors"
	"github.com/dop251/goja"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/mongodb"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/pkg/javascript"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"path/filepath"
)

const Path = "tables"

var scripts = []string{
	"before.insert",
	"after.insert",
	"before.delete",
	"after.delete",
	"before.update",
	"after.update",
}

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

func loadText(filename string) (string, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func loadJson(filename string, value any) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, value)
}

func loadJavaScript(filename string) (*goja.Program, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return javascript.Compile(string(buf))
}

func loadSchema(filename string) (*jsonschema.Schema, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return compiler.Compile(string(buf))
}

func Load(name string) error {
	dir := filepath.Join(viper.GetString("data"), Path, name)

	var table Table
	table.Name = name

	//直接注册表
	Register(&table)

	//加载字段
	err := loadJson(filepath.Join(dir, "fields.json"), &table.Fields)
	if err != nil {
		return err
	}

	//加载模式
	table.Schema, err = loadSchema(filepath.Join(dir, "schema.json"))
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	//加载累加器
	err = loadJson(filepath.Join(dir, "accumulations.json"), &table.Accumulations)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	err = table.init()
	if err != nil {
		return err
	}

	//加载时序配置
	err = loadJson(filepath.Join(dir, "time-serials.json"), &table.TimeSeries)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	//加载快照配置
	err = loadJson(filepath.Join(dir, "snapshot.json"), &table.Snapshot)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	//加载脚本
	for _, k := range scripts {
		js, err := loadJavaScript(filepath.Join(dir, k+".js"))
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}
		table.Scripts[k] = js
	}

	return nil
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
			err = Load(e.Name())
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
