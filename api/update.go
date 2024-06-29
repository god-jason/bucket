package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
)

func Update(tab *table.Table, after func(id string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var update db.Document
		err := ctx.ShouldBind(&update)
		if err != nil {
			Error(ctx, err)
			return
		}

		err = tab.Update(id, bson.D{{"$set", update}})
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

		OK(ctx, nil)
	}
}
