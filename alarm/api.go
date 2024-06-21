package alarm

import (
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "validator/create", api.Create(&_validatorTable, Load))
	api.Register("POST", "validator/update/:id", api.Update(&_validatorTable, Load))
	api.Register("GET", "validator/delete/:id", api.Delete(&_validatorTable, Unload))
	api.Register("GET", "validator/detail/:id", api.Detail(&_validatorTable, nil))
	api.Register("GET", "validator/enable/:id", api.Update(&_validatorTable, Load))
	api.Register("GET", "validator/disable/:id", api.Delete(&_validatorTable, Unload))
	api.Register("POST", "validator/count", api.Count(&_validatorTable))
	api.Register("POST", "validator/search", api.Search(&_validatorTable, nil))
	api.Register("POST", "validator/group", api.Group(&_validatorTable, nil))
	api.Register("POST", "validator/import", api.Import(&_validatorTable, func(ids []primitive.ObjectID) error {
		for _, id := range ids {
			err := Load(id)
			if err != nil {
				log.Error(err)
			}
		}
		return nil
	}))
	api.Register("POST", "validator/export", api.Export(&_validatorTable))
}

func init() {
	api.Register("POST", "alarm/search", api.Search(&_alarmTable, nil))
	api.Register("POST", "alarm/group", api.Group(&_alarmTable, nil))
	api.Register("POST", "alarm/export", api.Export(&_alarmTable))
	api.Register("POST", "alarm/count", api.Count(&_alarmTable))
}
