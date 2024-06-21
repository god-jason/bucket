package scene

import (
	"github.com/god-jason/bucket/boot"
)

func init() {
	boot.Register("scene", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "pool", "log", "database", "product", "project", "space", "device"},
	})
}

func Startup() error {

	return LoadAll()
}

func Shutdown() error {
	return nil
}
