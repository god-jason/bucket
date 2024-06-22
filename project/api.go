package project

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "project/create", api.Create(&_table, nil))
	api.Register("POST", "project/update/:id", api.Update(&_table, nil))
	api.Register("GET", "project/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "project/detail/:id", api.Detail(&_table, nil))
	api.Register("GET", "project/enable/:id", api.Enable(&_table, Load))
	api.Register("GET", "project/disable/:id", api.Disable(&_table, Unload))
	api.Register("POST", "project/count", api.Count(&_table))
	api.Register("POST", "project/search", api.Search(&_table, nil))
	api.Register("POST", "project/group", api.Group(&_table, nil))
	api.Register("POST", "project/import", api.Import(&_table, nil))
	api.Register("POST", "project/export", api.Export(&_table))

}
