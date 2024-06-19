package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("GET", "device/values/:id", deviceValues)
	api.Register("POST", "device/values/:id", deviceValuesUpdate)

}

func deviceValues(ctx *gin.Context) {
	dev := Get(ctx.Param("id"))
	if dev == nil {
		api.Fail(ctx, "设备不存在")
		return
	}
	api.OK(ctx, dev.values)
}

func deviceValuesUpdate(ctx *gin.Context) {
	dev := Get(ctx.Param("id"))
	if dev == nil {
		api.Fail(ctx, "设备不存在")
		return
	}
	var values map[string]any
	err := ctx.ShouldBindJSON(&values)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	err = dev.WriteValues(values)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
