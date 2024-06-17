package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
)

func init() {
	api.Register("GET", "device/values/:id", deviceValues)
	api.Register("POST", "device/values/:id", nil)

}

func deviceValues(ctx *gin.Context) {
	dev := Get(ctx.Param("id"))
	if dev == nil {
		curd.Fail(ctx, "设备不存在")
		return
	}
	curd.OK(ctx, dev.values)
}

func deviceValuesUpdate(ctx *gin.Context) {
	dev := Get(ctx.Param("id"))
	if dev == nil {
		curd.Fail(ctx, "设备不存在")
		return
	}
	var values map[string]any
	err := ctx.ShouldBindJSON(&values)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	err = dev.WriteValues(values)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
