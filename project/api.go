package project

import (
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "project/create", api.Create(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("POST", "project/update/:id", api.Update(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("GET", "project/delete/:id", api.Delete(&_table, func(id primitive.ObjectID) error {
		projects.Delete(id.Hex())
		//todo 删除相关 空间，设备绑定，场景，定时，等
		return nil
	}))

	api.Register("GET", "project/detail/:id", api.Detail(&_table, nil))

	api.Register("POST", "project/count", api.Count(&_table))

	api.Register("POST", "project/search", api.Search(&_table, nil))

	api.Register("POST", "project/group", api.Group(&_table, nil))

}
