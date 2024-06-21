package timer

import (
	"github.com/god-jason/bucket/boot"
	"github.com/robfig/cron/v3"
)

var _cron *cron.Cron

func init() {
	boot.Register("scene", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "pool", "log", "database", "product", "project", "space", "device"},
	})
}

func Startup() error {
	_cron = cron.New()
	_cron.Start()

	return LoadAll()
}

func Shutdown() error {
	return _cron.Stop().Err()
}
