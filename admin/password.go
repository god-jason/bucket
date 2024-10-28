package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/config"
)

type passwordObj struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func password(ctx *gin.Context) {
	if config.GetBool(MODULE, "lock") {
		Fail(ctx, "禁止修改密码")
		return
	}

	var obj passwordObj
	if err := ctx.ShouldBind(&obj); err != nil {
		Error(ctx, err)
		return
	}

	if obj.Old != config.GetString(MODULE, "password") {
		Fail(ctx, "密码错误")
		return
	}

	//更新密码
	config.Set(MODULE, "password", obj.New)

	err := config.Store()
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, nil)
}
