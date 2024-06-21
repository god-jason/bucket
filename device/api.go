package device

import (
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "device/create", api.Create(&_table, Load))
	api.Register("POST", "device/update/:id", api.Update(&_table, Load))
	api.Register("GET", "device/delete/:id", api.Delete(&_table, Unload))
	api.Register("GET", "device/detail/:id", api.Detail(&_table, nil))
	api.Register("GET", "device/enable/:id", api.Enable(&_table, Load))
	api.Register("GET", "device/disable/:id", api.Disable(&_table, Unload))
	api.Register("POST", "device/count", api.Count(&_table))
	api.Register("POST", "device/search", api.Search(&_table, nil))
	api.Register("POST", "device/group", api.Group(&_table, nil))
	api.Register("POST", "device/import", api.Import(&_table, func(ids []primitive.ObjectID) error {
		for _, id := range ids {
			err := Load(id)
			if err != nil {
				log.Error(err)
			}
		}
		return nil
	}))
	api.Register("POST", "device/export", api.Export(&_table))

}
