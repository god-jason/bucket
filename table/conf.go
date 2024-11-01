package table

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
)

func ApiConf(ctx *gin.Context) {
	tab := ctx.Param("table")
	conf := ctx.Param("conf")

	fn := filepath.Join(viper.GetString("data"), Path, tab, conf)
	_, err := os.Stat(fn)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			Fail(ctx, "不存在")
			return
		}
		Error(ctx, err)
		return
	}

	ctx.File(fn)
}

func ApiConfUpdate(ctx *gin.Context) {
	tab := ctx.Param("table")
	conf := ctx.Param("conf")

	fn := filepath.Join(viper.GetString("data"), Path, tab, conf)
	_ = os.MkdirAll(filepath.Dir(fn), os.ModePerm)
	//file, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file, err := os.Create(fn) //覆盖
	if err != nil {
		Error(ctx, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, ctx.Request.Body)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
