package user

import (
	"github.com/god-jason/bucket/api"
)

func init() {

	api.Register("POST", "role/count", api.Count(&_roleTable))
	api.Register("POST", "role/search", api.Search(&_roleTable, nil))
	api.Register("POST", "role/create", api.Create(&_roleTable, nil))
	api.Register("GET", "role/:id", api.Detail(&_roleTable, nil))
	api.Register("POST", "role/:id", api.Update(&_roleTable, nil))
	api.Register("GET", "role/:id/delete", api.Delete(&_roleTable, nil))
	api.Register("GET", "role/:id/disable", api.Disable(&_roleTable, nil))
	api.Register("GET", "role/:id/enable", api.Disable(&_roleTable, nil))
}
