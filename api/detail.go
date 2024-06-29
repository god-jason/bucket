package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
)

func Detail(tab *table.Table, after func(id string, doc db.Document) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var doc db.Document
		has, err := tab.Get(id, &doc)
		if err != nil {
			Error(ctx, err)
			return
		}
		if !has {
			Fail(ctx, "找不到记录")
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
