package project

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("project", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"web", "database", "log", "product", "device"},
	})
}

func Startup() error {

	//todo 启动项目

	//todo 场景

	return nil
}
