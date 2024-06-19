package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/log"
)

func init() {
	api.Register("GET", "device/start/:id", deviceStart)
	api.Register("GET", "device/stop/:id", deviceStop)
	api.Register("GET", "device/restart/:id", deviceRestart)

}

func deviceStart(ctx *gin.Context) {
	id := ctx.Param("id")

	err := Load(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	id := ctx.Param("id")

	dev := Get(id)
	if dev == nil {
		api.Fail(ctx, "设备不存在")
		return
	}

	err := dev.Close()
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}

func deviceRestart(ctx *gin.Context) {
	id := ctx.Param("id")

	dev := Get(id)
	if dev != nil {
		err := dev.Close()
		if err != nil {
			log.Error(err)
		}
		//报错
	}

	err := Load(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
