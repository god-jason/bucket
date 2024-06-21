package project

import (
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "project/create", api.Create(&_table, Load))

	api.Register("POST", "project/update/:id", api.Update(&_table, Load))

	api.Register("GET", "project/delete/:id", api.Delete(&_table, Unload))

	api.Register("GET", "project/detail/:id", api.Detail(&_table, nil))

	api.Register("GET", "project/enable/:id", api.Enable(&_table, Load))

	api.Register("GET", "project/disable/:id", api.Disable(&_table, Unload))

	api.Register("POST", "project/count", api.Count(&_table))

	api.Register("POST", "project/search", api.Search(&_table, nil))

	api.Register("POST", "project/group", api.Group(&_table, nil))

	api.Register("POST", "project/import", api.Import(&_table, func(ids []primitive.ObjectID) error {
		for _, id := range ids {
			err := Load(id)
			if err != nil {
				log.Error(err)
			}
		}
		return nil
	}))

	api.Register("POST", "project/export", api.Export(&_table))

}
