package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/mongodb"
)

func ApiUpdate(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var update mongodb.Document
	err = ctx.ShouldBindJSON(&update)
	if err != nil {
		Error(ctx, err)
		return
	}

	err = table.Update(ctx.Param("id"), update)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
