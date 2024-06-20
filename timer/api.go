package timer

import (
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "timer/create", api.Create(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("POST", "timer/update/:id", api.Update(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("GET", "timer/delete/:id", api.Delete(&_table, func(id primitive.ObjectID) error {
		timers.Delete(id.Hex())
		return nil
	}))

	api.Register("GET", "timer/detail/:id", api.Detail(&_table, nil))

	api.Register("POST", "timer/count", api.Count(&_table))

	api.Register("POST", "timer/search", api.Search(&_table, nil))

	api.Register("POST", "timer/group", api.Group(&_table, nil))

	api.Register("POST", "timer/import", api.Import(&_table, func(id []primitive.ObjectID) error {
		//TODO 加载
		return nil
	}))

	api.Register("POST", "timer/export", api.Export(&_table))

}
