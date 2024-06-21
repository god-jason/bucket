package product

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("product", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"web", "database", "log"},
	})
}

func Startup() error {
	return LoadAll()
}
