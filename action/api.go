package action

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "action/create", api.Create(&_table, nil))

	api.Register("GET", "action/delete/:id", api.Delete(&_table, nil))

	api.Register("GET", "action/detail/:id", api.Detail(&_table, nil))

	api.Register("POST", "action/count", api.Count(&_table))

	api.Register("POST", "action/search", api.Search(&_table, nil))

	api.Register("POST", "action/export", api.Export(&_table))

}
