package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
)

func init() {
	api.Register("POST", "table/:table/update", apiUpdate)
}

func apiUpdate(ctx *gin.Context) {
	table, err := GetTable(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	id := ctx.Param("id")

	var doc map[string]interface{}
	err = ctx.ShouldBindJSON(&doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	err = table.Update(id, doc)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
