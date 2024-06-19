package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/db"
)

func init() {
	api.Register("GET", "table/:table", apiDetail)
	api.Register("GET", "table/:table/detail/:id", apiDetail)
}

func apiDetail(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		api.Error(ctx, err)
		return
	}

	id, err := db.ParseObjectId(ctx.Param("id"))
	if err != nil {
		api.Error(ctx, err)
		return
	}

	var doc Document
	err = table.Get(id, &doc)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, doc)
}
