package table

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
)

type Info struct {
	Name          string `json:"name,omitempty"`
	Label         string `json:"label,omitempty"`
	Fields        int    `json:"fields,omitempty"`
	Schema        bool   `json:"schema,omitempty"`
	Scripts       int    `json:"scripts,omitempty"`
	Accumulations int    `json:"accumulations,omitempty"`
	TimeSeries    bool   `json:"time_series,omitempty"`
	Snapshot      bool   `json:"snapshot,omitempty"`
}

func ApiList(ctx *gin.Context) {

	var infos []*Info

	tables.Range(func(name string, tab *Table) bool {
		infos = append(infos, &Info{
			Name:          tab.Name,
			Label:         tab.Name,
			Fields:        len(tab.Fields),
			Schema:        tab.Schema != nil,
			Scripts:       len(tab.Scripts),
			Accumulations: len(tab.Accumulations),
			TimeSeries:    tab.TimeSeries != nil,
			Snapshot:      tab.Snapshot != nil,
		})
		return true
	})

	//TODO 加入空间统计等

	OK(ctx, infos)
}

func ApiConf(ctx *gin.Context) {
	tab := ctx.Param("table")
	conf := ctx.Param("conf")

	fn := filepath.Join(viper.GetString("data"), Path, tab, conf)
	_, err := os.Stat(fn)
	if err != nil {
		Error(ctx, err)
		return
	}

	ctx.File(fn)
}

func ApiConfUpdate(ctx *gin.Context) {
	tab := ctx.Param("table")
	conf := ctx.Param("conf")

	fn := filepath.Join(viper.GetString("data"), Path, tab, conf)
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		Error(ctx, err)
		return
	}

	_, err = io.Copy(file, ctx.Request.Body)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
