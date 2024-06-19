package device

import (
	"github.com/god-jason/bucket/aggregate"
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/history"
	"github.com/robfig/cron/v3"
	"time"
)

var _cron *cron.Cron

func init() {
	boot.Register("device", &boot.Task{
		Startup:  Startup, //启动
		Shutdown: Shutdown,
		Depends:  []string{"web", "pool", "log", "database", "product"},
	})
}

func Startup() error {
	aggregateStore = &db.Batch{
		Collection:   aggregate.Bucket,
		WriteTimeout: time.Second,
		BufferSize:   200,
	}

	historyStore = &db.Batch{
		Collection:   history.Bucket,
		WriteTimeout: time.Second,
		BufferSize:   200,
	}

	_cron = cron.New()
	_cron.Start()

	//整点聚合
	_, err := _cron.AddFunc("0 * * * *", hourAggregate)
	if err != nil {
		return err
	}

	return nil
}

func Shutdown() error {

	_, err := aggregateStore.Flush()
	if err != nil {
		//return err
	}

	_, err = historyStore.Flush()
	if err != nil {
		//return err
	}

	return _cron.Stop().Err()
}
