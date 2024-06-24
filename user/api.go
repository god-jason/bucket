package user

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {

	api.Register("GET", "logout", logout)
	api.Register("POST", "password", password)
	api.Register("GET", "user/me", userMe)
	api.Register("POST", "user/count", api.Count(&_table))
	api.Register("POST", "user/search", api.Search(&_table, nil))
	api.Register("POST", "user/create", api.Create(&_table, nil))
	api.Register("GET", "user/:id", api.Detail(&_table, nil))
	api.Register("POST", "user/:id", api.Update(&_table, nil))
	api.Register("GET", "user/:id/delete", api.Delete(nil, nil))
	api.Register("GET", "user/:id/password", userPassword)
	api.Register("GET", "user/:id/enable", api.Enable(&_table, nil))
	api.Register("GET", "user/:id/disable", api.Disable(&_table, nil))
}

func userMe(ctx *gin.Context) {
	id := ctx.GetString("user")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	var user User
	has, err := _table.Get(oid, &user)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if !has {
		//todo 删除session
		api.Fail(ctx, "找不到信息")
		return
	}
	api.OK(ctx, &user)
}

func userPassword(ctx *gin.Context) {
	id := ctx.GetString("user")
	oid, err := db.ParseObjectId(id)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	pwd := ctx.PostForm("password")
	//p.Password = md5hash(pwd)

	err = _passwordTable.Update(oid, bson.M{"password": pwd})
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
