package product

import (
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "product/create", api.Create(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("POST", "product/update/:id", api.Update(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("GET", "product/delete/:id", api.Delete(&_table, func(id primitive.ObjectID) error {
		products.Delete(id.Hex())
		return nil
	}))

	api.Register("GET", "product/detail/:id", api.Detail(&_table, nil))

	api.Register("POST", "product/count", api.Count(&_table))

	api.Register("POST", "product/search", api.Search(&_table, nil))

	api.Register("POST", "product/group", api.Group(&_table, nil))

	api.Register("POST", "product/import", api.Import(&_table, func(id []primitive.ObjectID) error {
		//TODO 加载
		return nil
	}))

	api.Register("POST", "product/export", api.Export(&_table))

}
