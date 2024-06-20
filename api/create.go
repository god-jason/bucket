package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(tab *table.Table, after func(id primitive.ObjectID) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var doc table.Document
		err := ctx.ShouldBind(&doc)
		if err != nil {
			Error(ctx, err)
			return
		}

		db.ConvertObjectId(doc)

		id, err := tab.Insert(doc)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			err = after(id)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, id)
	}
}
