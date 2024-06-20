package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
)

func ApiCreate(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var doc db.Document
	err = ctx.ShouldBindJSON(&doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	db.ConvertObjectId(doc)

	id, err := table.Insert(doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, id)
}
