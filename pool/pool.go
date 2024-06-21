package pool

import (
	"github.com/god-jason/bucket/config"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/errors"
	ants "github.com/panjf2000/ants/v2"
)

var Pool *ants.Pool

func Startup() (err error) {
	Pool, err = ants.NewPool(config.GetInt(MODULE, "size"), ants.WithPanicHandler(func(err any) {
		log.Error(err)
	}))
	return errors.Wrap(err)
}

func Shutdown() error {
	if Pool != nil {
		Pool.Release()
		Pool = nil
	}
	return nil
}

func Insert(task func()) error {
	if Pool == nil {
		go task()
	}
	err := Pool.Submit(task)
	return errors.Wrap(err)
}
