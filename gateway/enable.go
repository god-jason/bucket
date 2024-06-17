package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	api.Register("GET", "gateway/enable/:id", gatewayEnable)
	api.Register("GET", "gateway/disable/:id", gatewayDisable)

}

func gatewayEnable(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	_, err = db.UpdateById(Bucket, oid, bson.D{{"$set", bson.M{"disabled": false}}}, false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

func gatewayDisable(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	_, err = db.UpdateById(Bucket, oid, bson.D{{"$set", bson.M{"disabled": true}}}, false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//todo 关闭连接？

	curd.OK(ctx, nil)
}
