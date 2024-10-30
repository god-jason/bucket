package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("GET", "me", me)
	api.Register("GET", "logout", logout)
	api.Register("POST", "password", password)
}

func me(ctx *gin.Context) {
	id := ctx.GetString("user")

	if id == "" {
		api.Fail(ctx, "未登录")
		return
	}

	api.OK(ctx, gin.H{"id": id})
}
