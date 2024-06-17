package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
)

func init() {
	api.Register("GET", "table/:table/delete", Delete)
	api.Register("DELETE", "table/:table/delete", Delete)
	//api.Register("POST", "table/:table/delete", DeleteMany)
}

func Delete(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//ids := ctx.QueryArray("id") //依次删除
	id, err := db.ParseObjectId(ctx.Query("id"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	err = table.Delete(id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}
