package space

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "space/create", api.Create(&_table, nil))
	api.Register("POST", "space/update/:id", api.Update(&_table, nil))
	api.Register("GET", "space/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "space/detail/:id", api.Detail(&_table, nil))
	api.Register("POST", "space/count", api.Count(&_table))
	api.Register("POST", "space/search", api.Search(&_table, nil))
	api.Register("POST", "space/group", api.Group(&_table, nil))
	api.Register("POST", "space/import", api.Import(&_table, nil))
	api.Register("POST", "space/export", api.Export(&_table))

}
