package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
)

func ApiUpdate(ctx *gin.Context) {
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

	var update db.Document
	err = ctx.ShouldBindJSON(&update)
	if err != nil {
		Error(ctx, err)
		return
	}

	db.ConvertObjectId(update)

	err = table.Update(id, update)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
