package project

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/table"
)

var projects lib.Map[Project]

func Get(id string) *Project {
	return projects.Load(id)
}

func From(t *Project) (err error) {
	tt := projects.LoadAndStore(t.Id, t)
	if tt != nil {
		_ = tt.Close()
	}
	//禁用的不再打开
	if t.Disabled {
		return nil
	}
	return t.Open()
}

func Load(id string) error {
	var project Project
	has, err := _table.Get(id, &project)
	if err != nil {
		return err
	}
	if !has {
		return exception.New("找不到记录")
	}
	return From(&project)
}

func Unload(id string) error {
	t := projects.LoadAndDelete(id)
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Project](&_table, base.FilterEnabled, 100, func(t *Project) error {
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}
