package user

import (
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/web"
)

func init() {
	boot.Register("user", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"database", "web"},
	})
}

func Startup() error {

	//鉴权接口
	web.Engine.GET("api/auth", auth)

	web.Engine.POST("api/login", login)

	return nil
}
