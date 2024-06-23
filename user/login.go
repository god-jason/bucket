package user

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson"
)

type loginObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func md5hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	var obj loginObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		api.Error(ctx, err)
		return
	}

	var users []User
	var user User

	err := _table.Find(bson.D{{"username", obj.Username}}, nil, 0, 1, &users)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if len(users) == 0 {
		//管理员自动创建
		if obj.Username == "admin" {
			user.Name = "管理员"
			user.Admin = true

			user.Id, err = _table.Insert(&user)
			if err != nil {
				api.Error(ctx, err)
				return
			}
		} else {
			api.Fail(ctx, "找不到用户")
			return
		}
	} else {
		user = users[0]
	}

	if user.Disabled {
		api.Fail(ctx, "用户已禁用")
		return
	}

	var password Password
	err = _passwordTable.Get(user.Id, &password)
	if err != nil {
		//todo 找不到
		api.Error(ctx, err)
		return
	}

	//初始化密码
	if !has {
		dp := "123456"

		password.Id = user.Id
		password.Password = md5hash(dp)
		_, err = _passwordTable.Insert(&password)
		if err != nil {
			api.Error(ctx, err)
			return
		}
	}

	if password.Password != obj.Password {
		api.Fail(ctx, "密码错误")
		return
	}

	//_, _ = db.Engine.InsertOne(&types.UserEvent{UserId: user.id, ModEvent: types.ModEvent{Type: "登录"}})

	//存入session
	session.Set("user", user.Id)
	_ = session.Save()

	api.OK(ctx, user)
}

func logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	u := session.Get("user")
	if u == nil {
		api.Fail(ctx, "未登录")
		return
	}

	//user := u.(int64)
	//_, _ = db.Engine.InsertOne(&types.UserEvent{UserId: user, ModEvent: types.ModEvent{Type: "退出"}})

	session.Clear()
	_ = session.Save()
	api.OK(ctx, nil)
}

type passwordObj struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func password(ctx *gin.Context) {

	var obj passwordObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		api.Error(ctx, err)
		return
	}

	user := ctx.GetString("user")
	oid, err := db.ParseObjectId(user)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	var pwd Password
	err = _passwordTable.Get(oid, &pwd)
	if err != nil {
		//todo 找不到
		api.Error(ctx, err)
		return
	}
	if obj.Old != pwd.Password {
		api.Fail(ctx, "密码错误")
		return
	}

	//pwd.Password = obj.New //前端已经加密过
	err = _passwordTable.Update(oid, bson.M{"password": pwd})
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
