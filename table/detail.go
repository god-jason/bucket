package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/mongodb"
)

func ApiDetail(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var doc mongodb.Document
	has, err := table.Get(ctx.Param("id"), &doc)
	if err != nil {
		Error(ctx, err)
		return
	}
	if !has {
		Fail(ctx, "找不到记录")
		return
	}

	OK(ctx, doc)
}
