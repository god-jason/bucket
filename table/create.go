package table

import (
	"github.com/gin-gonic/gin"
)

func ApiCreate(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var doc Document
	err = ctx.ShouldBindJSON(&doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	id, err := table.Insert(doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, id)
}
