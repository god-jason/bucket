package table

import (
	"github.com/gin-gonic/gin"
)

func ApiCount(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var filter interface{}
	err = ctx.ShouldBindJSON(&filter)
	if err != nil {
		Error(ctx, err)
		return
	}

	ret, err := table.Count(filter)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, ret)
}
