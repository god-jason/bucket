package db

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("database", &boot.Task{
		Startup:  Open,
		Shutdown: Close,
		Depends:  []string{"config"},
	})
}
