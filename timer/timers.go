package timer

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/table"
)

var timers lib.Map[Timer]

func Get(id string) *Timer {
	return timers.Load(id)
}

func From(t *Timer) (err error) {
	tt := timers.LoadAndStore(t.Id, t)
	if tt != nil {
		_ = tt.Close()
	}
	return t.Open()
}

func Load(id string) error {
	var timer Timer
	has, err := _table.Get(id, &timer)
	if err != nil {
		return err
	}
	if !has {
		return exception.New("找不到记录")
	}
	return From(&timer)
}

func Unload(id string) error {
	t := timers.LoadAndDelete(id)
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Timer](&_table, base.FilterEnabled, 100, func(t *Timer) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}

func Execute(id string) error {
	t := timers.Load(id)
	if t != nil {
		return t.Execute()
	}
	return exception.New("找不到定时器")
}
