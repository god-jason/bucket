package table

import (
	"encoding/json"
	"errors"
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/javascript"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var scripts = []string{
	"before.insert",
	"after.insert",
	"before.delete",
	"after.delete",
	"before.update",
	"after.update",
}

type InfoEx struct {
	Name          string `json:"name,omitempty"`
	Label         string `json:"label,omitempty"`
	Fields        int    `json:"fields,omitempty"`
	Schema        bool   `json:"schema,omitempty"`
	Scripts       int    `json:"scripts,omitempty"`
	Accumulations int    `json:"accumulations,omitempty"`
	TimeSeries    bool   `json:"time_series,omitempty"`
	Snapshot      bool   `json:"snapshot,omitempty"`
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

func Load(id string, strict bool) error {
	dir := filepath.Join(viper.GetString("data"), Path, id)

	var table Table

	//加载信息
	fn := filepath.Join(dir, "info.json")
	err := loadJson(fn, &table.Info)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			if strict {
				return err
			} else {
				log.Error(fn, err)
			}
		}
	}

	table.Id = id //避免被Info覆盖

	//加载字段
	fn = filepath.Join(dir, "fields.json")
	err = loadJson(fn, &table.Fields)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			if strict {
				return err
			} else {
				log.Error(fn, err)
			}
		}
	}

	//加载模式
	fn = filepath.Join(dir, "schema.json")
	table.Schema, err = loadSchema(fn)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			if strict {
				return err
			} else {
				log.Error(err)
			}
		}
	}

	//加载累加器
	fn = filepath.Join(dir, "accumulations.json")
	err = loadJson(fn, &table.Accumulations)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			if strict {
				return err
			} else {
				log.Error(fn, err)
			}
		}
	}
	err = table.init()
	if err != nil {
		if strict {
			return err
		} else {
			log.Error(err)
		}
	}

	//加载时序配置
	fn = filepath.Join(dir, "time-serials.json")
	err = loadJson(fn, &table.TimeSeries)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			if strict {
				return err
			} else {
				log.Error(fn, err)
			}
		}
	}

	//加载快照配置
	fn = filepath.Join(dir, "snapshot.json")
	err = loadJson(fn, &table.Snapshot)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			if strict {
				return err
			} else {
				log.Error(fn, err)
			}
		}
	}

	//加载脚本
	table.Scripts = make(map[string]*goja.Program)
	for _, k := range scripts {
		fn := filepath.Join(dir, k+".js")
		js, err := loadJavaScript(fn)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				if strict {
					return err
				} else {
					log.Error(fn, err)
				}
			}
			continue
		} else {
			table.Scripts[k] = js
		}
	}

	//直接注册表
	Register(&table)

	return nil
}

func ApiTableList(ctx *gin.Context) {

	var infos []Info
	//&InfoEx{
	//	Name:          tab.Name,
	//	Label:         tab.Name,
	//	Fields:        len(tab.Fields),
	//	Schema:        tab.Schema != nil,
	//	Scripts:       len(tab.Scripts),
	//	Accumulations: len(tab.Accumulations),
	//	TimeSeries:    tab.TimeSeries != nil,
	//	Snapshot:      tab.Snapshot != nil,
	//}
	tables.Range(func(name string, tab *Table) bool {
		infos = append(infos, tab.Info)
		return true
	})

	OK(ctx, infos)
}

func ApiTableReload(ctx *gin.Context) {
	tab := ctx.Param("table")
	err := Load(tab, true)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}

type rename struct {
	Name string `json:"name"`
}

func ApiTableRename(ctx *gin.Context) {
	var r rename
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		Error(ctx, err)
		return
	}

	tab := ctx.Param("table")
	old := filepath.Join(viper.GetString("data"), Path, tab)
	name := filepath.Join(viper.GetString("data"), Path, r.Name)
	err = os.Rename(old, name)
	if err != nil {
		Error(ctx, err)
		return
	}

	//直接修改map，不雅
	t := tables.LoadAndDelete(tab)
	t.Id = r.Name
	tables.Store(r.Name, t)

	OK(ctx, nil)
}

func ApiTableRemove(ctx *gin.Context) {
	tab := ctx.Param("table")
	dir := filepath.Join(viper.GetString("data"), Path, tab)
	err := os.RemoveAll(dir)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}

func ApiTableCreate(ctx *gin.Context) {
	tab := ctx.Param("table")
	dir := filepath.Join(viper.GetString("data"), Path, tab)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		Error(ctx, err)
		return
	}

	//TODO 空白的
	_ = Load(tab, false)

	OK(ctx, nil)

}
