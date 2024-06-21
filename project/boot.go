package project

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("project", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "pool", "log", "database", "device"},
	})
}

func Startup() error {

	return LoadAll()
}

func Shutdown() error {
	return nil
}
