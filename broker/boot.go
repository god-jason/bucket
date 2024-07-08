package broker

import (
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/web"
	mqtt "github.com/mochi-mqtt/server/v2"
)

func init() {
	boot.Register("broker", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "log", "database"},
	})
}

var server *mqtt.Server

func Startup() error {
	opts := &mqtt.Options{
		InlineClient: true,
	}
	server = mqtt.New(opts)

	err := server.AddHook(new(Hook), nil)
	if err != nil {
		return err
	}

	//监听Websocket
	web.Engine.GET("/mqtt", GinBridge)

	return server.Serve()
}

func Shutdown() error {
	if server != nil {
		return server.Close()
	}
	return nil
}
