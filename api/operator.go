package api

import (
	"github.com/gin-gonic/gin"
)

func Operator(hook func(id string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := hook(id)
		if err != nil {
			Error(ctx, err)
			return
		}

		OK(ctx, nil)
	}
}
