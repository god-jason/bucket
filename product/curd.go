package product

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	api.Register("POST", "product/create", productCreate)
	api.Register("POST", "product/update/:id", productUpdate)
	api.Register("GET", "product/delete/:id", productDelete)
	api.Register("GET", "product/detail/:id", productDetail)
}

func productCreate(ctx *gin.Context) {
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

func productUpdate(ctx *gin.Context) {
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

	_, err = db.UpdateByID(Bucket, oid, bson.D{{"$set", update}}, false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

func productDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	_, err = db.DeleteByID(Bucket, oid)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//TODO 断开连接？

	curd.OK(ctx, nil)
}

func productDetail(ctx *gin.Context) {
	id, err := db.ParseObjectId(ctx.Param("id"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var doc table.Document
	err = db.FindByID(Bucket, id, &doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, doc)
}
