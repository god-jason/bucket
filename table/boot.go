package table

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("table", &boot.Task{
		Startup:  nil, //TODO 启动
		Shutdown: nil,
		Depends:  []string{"db"},
	})
}
