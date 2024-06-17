package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	api.Register("GET", "device/enable/:id", deviceEnable)
	api.Register("GET", "device/disable/:id", deviceDisable)

}

func deviceEnable(ctx *gin.Context) {
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

	err = Load(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

func deviceDisable(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	dev := Get(id)
	if dev != nil {
		_ = dev.Close()
	}

	_, err = db.UpdateById(Bucket, oid, bson.D{{"$set", bson.M{"disabled": true}}}, false)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
