package space

import (
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "space/create", api.Create(&_table, Load))
	api.Register("POST", "space/update/:id", api.Update(&_table, Load))
	api.Register("GET", "space/delete/:id", api.Delete(&_table, Unload))
	api.Register("GET", "space/detail/:id", api.Detail(&_table, nil))
	api.Register("POST", "space/count", api.Count(&_table))
	api.Register("POST", "space/search", api.Search(&_table, nil))
	api.Register("POST", "space/group", api.Group(&_table, nil))
	api.Register("POST", "space/import", api.Import(&_table, func(ids []primitive.ObjectID) error {
		for _, id := range ids {
			err := Load(id)
			if err != nil {
				log.Error(err)
			}
		}
		return nil
	}))
	api.Register("POST", "space/export", api.Export(&_table))

}
