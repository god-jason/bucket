package table

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func ApiManifest(ctx *gin.Context) {
	tab := ctx.Param("table")
	fn := filepath.Join(viper.GetString("data"), Path, tab+".json")
	buf, err := os.ReadFile(fn)
	if err != nil {
		Error(ctx, err)
		return
	}

	var data any
	err = json.Unmarshal(buf, &data)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, data)
}

func ApiManifestUpdate(ctx *gin.Context) {
	tab := ctx.Param("table")
	var data any
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		Error(ctx, err)
		return
	}

	buf, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		Error(ctx, err)
		return
	}

	_ = os.MkdirAll(filepath.Join(viper.GetString("data"), Path), os.ModePerm)
	fn := filepath.Join(viper.GetString("data"), Path, tab+".json")
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		Error(ctx, err)
		return
	}
	defer file.Close()
	_, err = file.Write(buf)
	if err != nil {
		Error(ctx, err)
		return
	}

	//加载
	err = Load(tab)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
