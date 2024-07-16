package table

import (
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/log"
	"github.com/robfig/cron/v3"
)

func init() {
	boot.Register("table", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"database", "pool"},
	})
}

var _cron *cron.Cron

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

	//定时快照
	_cron = cron.New()
	tables.Range(func(name string, item *Table) bool {
		if item.Snapshot != nil && item.Snapshot.Crontab != "" {
			_, err := _cron.AddFunc("0 0 * * *", func() {
				err := item.snapshot()
				if err != nil {
					log.Error(err)
				}
			})
			if err != nil {
				log.Error(err)
			}
		}
		return true
	})
	_cron.Start()

	return nil
}

func Shutdown() error {
	return _cron.Stop().Err()
}
