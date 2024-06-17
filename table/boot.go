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
	//加载表定义，编译schema，创建表

	err := LoadAll()
	if err != nil {
		return err
	}

	err = Sync()
	if err != nil {
		return err
	}

	return nil
}
