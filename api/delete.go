package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/table"
)

func Delete(tab *table.Table, after func(id string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := tab.Delete(id)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			err = after(id)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, nil)
	}

}
