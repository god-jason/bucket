package gateway

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("gateway", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "log", "database"},
	})
}
