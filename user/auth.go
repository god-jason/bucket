package user

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/config"
	"github.com/god-jason/bucket/web"
	"go.mongodb.org/mongo-driver/bson"
)

func auth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	var users []*User
	err := _table.Find(bson.D{{"username", username}}, nil, 0, 1, &users)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if len(users) == 0 {
		api.Fail(ctx, "找不到用户")
		return
	}

	user := users[0]

	if user.Disabled {
		api.Fail(ctx, "用户已禁用")
		return
	}

	var obj Password
	has, err := _passwordTable.Get(user.Id, &obj)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	//初始化密码
	if !has {
		dp := config.GetString(MODULE, "default_password")
		obj.Password = passwordHash(dp)
	}

	if obj.Password != password {
		api.Fail(ctx, "密码错误")
		return
	}

	//生成Token
	token, err := web.JwtGenerate(user.Id)
	if err != nil {
		return
	}

	api.OK(ctx, gin.H{
		token: token,
	})
}
