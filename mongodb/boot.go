package mongodb

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("mongodb", &boot.Task{
		Startup:  Open,
		Shutdown: Close,
		Depends:  []string{"config", "log"},
	})
}
