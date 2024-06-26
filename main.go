package main

import (
	_ "github.com/god-jason/bucket/action"
	_ "github.com/god-jason/bucket/aggregate"
	_ "github.com/god-jason/bucket/alarm"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/broker"
	_ "github.com/god-jason/bucket/device"
	_ "github.com/god-jason/bucket/function"
	_ "github.com/god-jason/bucket/gateway"
	_ "github.com/god-jason/bucket/history"
	"github.com/god-jason/bucket/log"
	_ "github.com/god-jason/bucket/product"
	_ "github.com/god-jason/bucket/scene"
	_ "github.com/god-jason/bucket/table"
	_ "github.com/god-jason/bucket/timer"
	_ "github.com/god-jason/bucket/user"
	"github.com/god-jason/bucket/web"
)

func main() {
	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
	}

	//注册接口
	api.RegisterRoutes(web.Engine.Group("api"))

	//mqtt 监听Websocket
	web.Engine.GET("/mqtt", broker.GinBridge)

	//注册静态文件
	web.Static.PutDir("", "www", "", "index.html")

	//监听
	err = web.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
