package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
)

func init() {
	api.Register("GET", "table/:table/delete", Delete)
	api.Register("DELETE", "table/:table/delete", Delete)
	api.Register("POST", "table/:table/delete", DeleteMany)
}

func Delete(ctx *gin.Context) {
	table, err := GetTable(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//ids := ctx.QueryArray("id") //依次删除
	id := ctx.Param("id")

	ret, err := table.DeleteByID(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, ret)
}

func DeleteMany(ctx *gin.Context) {
	table, err := GetTable(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var filter map[string]interface{}
	err = ctx.ShouldBindJSON(&filter)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	ret, err := table.DeleteMany(filter)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, ret)
}
