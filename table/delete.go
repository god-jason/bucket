package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/db"
)

func init() {
	api.Register("GET", "table/:table/delete/:id", apiDelete)
	api.Register("DELETE", "table/:table/delete/:id", apiDelete)
	//api.Register("POST", "table/:table/delete", DeleteMany)
}

func apiDelete(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		api.Error(ctx, err)
		return
	}

	//ids := ctx.QueryArray("id") //依次删除
	id, err := db.ParseObjectId(ctx.Param("id"))
	if err != nil {
		api.Error(ctx, err)
		return
	}

	err = table.Delete(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
