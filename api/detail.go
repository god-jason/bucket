package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Detail(tab *table.Table, after func(id primitive.ObjectID, doc db.Document) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := db.ParseObjectId(ctx.Param("id"))
		if err != nil {
			Error(ctx, err)
			return
		}

		var doc db.Document
		err = tab.Get(id, &doc)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			err = after(id, doc)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, doc)
	}
}
