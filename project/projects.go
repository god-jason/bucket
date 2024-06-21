package project

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var projects lib.Map[Project]

func Get(id string) *Project {
	return projects.Load(id)
}

func From(t *Project) (err error) {
	tt := projects.LoadAndStore(t.Id.Hex(), t)
	if tt != nil {
		_ = tt.Close()
	}
	//禁用的不再打开
	if t.Disabled {
		return nil
	}
	return t.Open()
}

func Load(id primitive.ObjectID) error {
	var project Project
	err := _table.Get(id, &project)
	if err != nil {
		return err
	}
	return From(&project)
}

func Unload(id primitive.ObjectID) error {
	t := projects.LoadAndDelete(id.Hex())
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
