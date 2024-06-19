package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(tab *table.Table, after func(id primitive.ObjectID) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		oid, err := db.ParseObjectId(id)
		if err != nil {
			Error(ctx, err)
			return
		}

		err = tab.Delete(oid)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			err = after(oid)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, nil)
	}

}
