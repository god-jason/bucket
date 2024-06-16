package product

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("device", &boot.Task{
		Startup:  nil, //TODO 启动
		Shutdown: nil,
		Depends:  []string{"web"},
	})
}
