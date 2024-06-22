package space

import (
	"github.com/god-jason/bucket/boot"
)

func init() {
	boot.Register("space", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "pool", "log", "database", "project"},
	})
}

func Startup() error {

	return LoadAll()
}

func Shutdown() error {
	return nil
}
