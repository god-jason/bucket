package scene

import (
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "scene/create", api.Create(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("POST", "scene/update/:id", api.Update(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("GET", "scene/delete/:id", api.Delete(&_table, func(id primitive.ObjectID) error {
		scenes.Delete(id.Hex())
		//todo 删除相关 空间，设备绑定，场景，定时，等
		return nil
	}))

	api.Register("GET", "scene/detail/:id", api.Detail(&_table, nil))

	api.Register("POST", "scene/count", api.Count(&_table))

	api.Register("POST", "scene/search", api.Search(&_table, nil))

	api.Register("POST", "scene/group", api.Group(&_table, nil))

	api.Register("POST", "scene/import", api.Import(&_table, func(id []primitive.ObjectID) error {
		//TODO 加载
		return nil
	}))

	api.Register("POST", "scene/export", api.Export(&_table))

}
