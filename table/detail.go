package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
)

func ApiDetail(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	id, err := db.ParseObjectId(ctx.Param("id"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var doc db.Document
	err = table.Get(id, &doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, doc)
}
