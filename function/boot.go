package function

import "github.com/god-jason/bucket/boot"

const Path = "functions"

func init() {
	boot.Register("function", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: nil,
		Depends:  []string{"log"},
	})
}

func Startup() error {

	err := LoadAll()
	if err != nil {
		return err
	}

	return nil
}
