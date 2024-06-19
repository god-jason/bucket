package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "table/:table/import", Import)
}

func Import(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		api.Error(ctx, err)
		return
	}

	var doc []Document
	err = ctx.ShouldBindJSON(&doc)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	id, err := table.Import(doc)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, id)
}
