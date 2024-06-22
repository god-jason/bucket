package scene

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var scenes lib.Map[Scene]

func Get(id string) *Scene {
	return scenes.Load(id)
}

func From(t *Scene) (err error) {
	tt := scenes.LoadAndStore(t.Id.Hex(), t)
	if tt != nil {
		_ = tt.Close()
	}
	return t.Open()
}

func Load(id primitive.ObjectID) error {
	var scene Scene
	err := _table.Get(id, &scene)
	if err != nil {
		return err
	}
	return From(&scene)
}

func Unload(id primitive.ObjectID) error {
	t := scenes.LoadAndDelete(id.Hex())
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Scene](&_table, base.FilterEnabled, 100, func(t *Scene) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}

func Execute(id primitive.ObjectID) error {
	t := scenes.Load(id.Hex())
	if t != nil {
		return t.Execute()
	}
	return errors.New("找不到场景")
}
