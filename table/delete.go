package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
)

func ApiDelete(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	//ids := ctx.QueryArray("id") //依次删除
	id, err := db.ParseObjectId(ctx.Param("id"))
	if err != nil {
		Error(ctx, err)
		return
	}

	err = table.Delete(id)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
