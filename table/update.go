package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/db"
)

func init() {
	api.Register("POST", "table/:table/update/:id", apiUpdate)
}

func apiUpdate(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	id, err := db.ParseObjectId(ctx.Param("id"))
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

	err = table.Update(id, doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
