package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	api.Register("POST", "gateway/create", gatewayCreate)
	api.Register("POST", "gateway/update/:id", gatewayUpdate)
	api.Register("GET", "gateway/delete/:id", gatewayDelete)
	api.Register("GET", "gateway/detail/:id", gatewayDetail)
}

func gatewayCreate(ctx *gin.Context) {
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

	curd.OK(ctx, id)
}

func gatewayUpdate(ctx *gin.Context) {
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

	curd.OK(ctx, nil)
}

func gatewayDelete(ctx *gin.Context) {
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

	//TODO 断开连接？

	curd.OK(ctx, nil)
}

func gatewayDetail(ctx *gin.Context) {
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
