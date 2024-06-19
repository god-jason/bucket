package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "table/:table/create", apiCreate)
}

func apiCreate(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		api.Error(ctx, err)
		return
	}

	var doc Document
	err = ctx.ShouldBindJSON(&doc)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	id, err := table.Insert(doc)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, id)
}
