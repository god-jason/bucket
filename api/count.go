package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/table"
)

func Count(tab *table.Table) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter any
		err := ctx.ShouldBindJSON(&filter)
		if err != nil {
			Error(ctx, err)
			return
		}

		ret, err := tab.Count(tab.Name)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, ret)
	}

}
