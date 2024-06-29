package table

import (
	"github.com/gin-gonic/gin"
)

func ApiDelete(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	//ids := ctx.QueryArray("id") //依次删除

	err = table.Delete(ctx.Param("id"))
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
