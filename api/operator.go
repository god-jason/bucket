package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Operator(hook func(id primitive.ObjectID) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		oid, err := db.ParseObjectId(id)
		if err != nil {
			Error(ctx, err)
			return
		}

		err = hook(oid)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, nil)
	}
}
