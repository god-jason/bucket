package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
)

func init() {
	api.Register("POST", "table/:table/create", Create)
}

func Create(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var doc Document
	err = ctx.ShouldBindJSON(&doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	id, err := table.Insert(doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, id)
}
