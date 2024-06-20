package alarm

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "validator/create", api.Create(&_validatorTable, nil))
	api.Register("POST", "validator/update/:id", api.Update(&_validatorTable, nil))
	api.Register("GET", "validator/delete/:id", api.Delete(&_validatorTable, nil))
	api.Register("GET", "validator/detail/:id", api.Detail(&_validatorTable, nil))
	api.Register("POST", "validator/count", api.Count(&_validatorTable))
	api.Register("POST", "validator/search", api.Search(&_validatorTable, nil))
	api.Register("POST", "validator/group", api.Group(&_validatorTable, nil))
}

func init() {
	api.Register("POST", "alarm/count", api.Count(&_validatorTable))
	api.Register("POST", "alarm/search", api.Search(&_validatorTable, nil))
	api.Register("POST", "alarm/group", api.Group(&_validatorTable, nil))
}
