package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
)

func init() {
	api.Register("GET", "device/start/:id", deviceStart)
	api.Register("GET", "device/stop/:id", deviceStop)
	api.Register("GET", "device/restart/:id", deviceRestart)

}

func deviceStart(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var doc table.Document

	err = db.FindByID(Bucket, oid, &doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	err = Load(doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	id := ctx.Param("id")
	dev := Get(id)
	if dev == nil {
		curd.Fail(ctx, "设备不存在")
		return
	}

	devices.Delete(id)

	curd.OK(ctx, nil)
}

func deviceRestart(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	dev := Get(id)
	if dev != nil {
		err = dev.Close()
		//报错
	}

	var doc table.Document
	err = db.FindByID(Bucket, oid, &doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	err = Load(doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
