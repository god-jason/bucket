package product

import (
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "product/create", api.Create(&_table, Load))
	api.Register("POST", "product/update/:id", api.Update(&_table, Load))
	api.Register("GET", "product/delete/:id", api.Delete(&_table, Unload))
	api.Register("GET", "product/detail/:id", api.Detail(&_table, nil))
	api.Register("POST", "product/count", api.Count(&_table))
	api.Register("POST", "product/search", api.Search(&_table, nil))
	api.Register("POST", "product/group", api.Group(&_table, nil))
	api.Register("POST", "product/import", api.Import(&_table, func(ids []primitive.ObjectID) error {
		for _, id := range ids {
			err := Load(id)
			if err != nil {
				log.Error(err)
			}
		}
		return nil
	}))
	api.Register("POST", "product/export", api.Export(&_table))
}
