package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
)

func init() {
	//api.Register("GET", "table/:table/count", apiCount)
	api.Register("POST", "table/:table/count", apiCount)
}

func apiCount(ctx *gin.Context) {
	table, err := GetTable(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var filter interface{}
	err = ctx.ShouldBindJSON(&filter)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	ret, err := table.Count(filter)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, ret)
}
