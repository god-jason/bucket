package table

import "github.com/gin-gonic/gin"

func ApiReload(ctx *gin.Context) {
	tab := ctx.Param("table")
	err := Load(tab)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
