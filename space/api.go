package space

import (
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "space/create", api.Create(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("POST", "space/update/:id", api.Update(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("GET", "space/delete/:id", api.Delete(&_table, func(id primitive.ObjectID) error {
		spaces.Delete(id.Hex())
		//todo 删除场景，报警等
		return nil
	}))

	api.Register("GET", "space/detail/:id", api.Detail(&_table, nil))

	api.Register("POST", "space/count", api.Count(&_table))

	api.Register("POST", "space/search", api.Search(&_table, nil))

	api.Register("POST", "space/group", api.Group(&_table, nil))

	api.Register("POST", "space/import", api.Import(&_table, func(id []primitive.ObjectID) error {
		//TODO 加载
		return nil
	}))

	api.Register("POST", "space/export", api.Export(&_table))

}
