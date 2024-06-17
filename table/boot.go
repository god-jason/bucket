package table

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("table", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"database", "pool"},
	})
}

func Startup() error {
	//TODO 加载表定义，编译schema，创建表

	return nil
}
