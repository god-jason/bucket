package table

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("table", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"database"},
	})
}

func Startup() error {
	//TODO 启动

	return nil
}
