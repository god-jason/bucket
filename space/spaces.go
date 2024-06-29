package space

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/table"
)

var spaces lib.Map[Space]

func Get(id string) *Space {
	return spaces.Load(id)
}

func From(t *Space) (err error) {
	tt := spaces.LoadAndStore(t.Id, t)
	if tt != nil {
		_ = tt.Close()
	}
	return t.Open()
}

func Load(id string) error {
	var space Space
	has, err := _table.Get(id, &space)
	if err != nil {
		return err
	}
	if !has {
		return exception.New("找不到记录")
	}
	return From(&space)
}

func Unload(id string) error {
	t := spaces.LoadAndDelete(id)
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Space](&_table, base.FilterEnabled, 100, func(t *Space) error {
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}
