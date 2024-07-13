package main

import (
	_ "github.com/god-jason/bucket/aggregate"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/boot"
	_ "github.com/god-jason/bucket/function"
	"github.com/god-jason/bucket/log"
	_ "github.com/god-jason/bucket/table"
	"github.com/god-jason/bucket/web"
)

func main() {
	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
	}

	//注册接口
	api.RegisterRoutes(web.Engine.Group("api"))

	//注册静态文件
	web.Static.PutDir("", "www", "", "index.html")

	//监听
	err = web.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
