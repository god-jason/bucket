package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	api.Register("POST", "device/create", deviceCreate)
	api.Register("POST", "device/update/:id", deviceUpdate)
	api.Register("GET", "device/delete/:id", deviceDelete)
	api.Register("GET", "device/detail/:id", deviceDetail)
}

func deviceCreate(ctx *gin.Context) {
	var doc table.Document
	err := ctx.ShouldBind(&doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	id, err := db.InsertOne(Bucket, doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	err = Load(id.Hex())
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, id)
}

func deviceUpdate(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var update table.Document
	err = ctx.ShouldBind(&update)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	_, err = db.UpdateById(Bucket, oid, bson.D{{"$set", update}}, false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	dev := Get(id)
	if dev != nil {
		err = dev.Close()
		//报错
	}

	err = Load(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

func deviceDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	_, err = db.DeleteById(Bucket, oid)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	dev := Get(id)
	if dev != nil {
		err = dev.Close()
		//报错
	}

	curd.OK(ctx, nil)
}

func deviceDetail(ctx *gin.Context) {
	id, err := db.ParseObjectId(ctx.Param("id"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var doc table.Document
	err = db.FindById(Bucket, id, &doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, doc)
}
