package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
)

func init() {
	api.Register("GET", "table/:table", apiDetail)
	api.Register("GET", "table/:table/detail", apiDetail)
}

func apiDetail(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	id, err := db.ParseObjectId(ctx.Query("id"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var doc Document
	err = table.Get(id, &doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, doc)
}
