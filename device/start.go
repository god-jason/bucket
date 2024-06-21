package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/db"
)

func init() {
	api.Register("GET", "device/start/:id", deviceStart)
	api.Register("GET", "device/stop/:id", deviceStop)
}

func deviceStart(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	err = Load(oid)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	err = Unload(oid)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
