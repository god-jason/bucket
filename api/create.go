package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/mongodb"
	"github.com/god-jason/bucket/table"
)

func Create(tab *table.Table, after func(id string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var doc mongodb.Document
		err := ctx.ShouldBind(&doc)
		if err != nil {
			Error(ctx, err)
			return
		}
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
