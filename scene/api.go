package scene

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "scene/create", api.Create(&_table, nil))
	api.Register("POST", "scene/update/:id", api.Update(&_table, nil))
	api.Register("GET", "scene/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "scene/detail/:id", api.Detail(&_table, nil))
	api.Register("GET", "scene/enable/:id", api.Update(&_table, Load))
	api.Register("GET", "scene/disable/:id", api.Delete(&_table, Unload))
	api.Register("GET", "scene/execute/:id", api.Operator(Execute))
	api.Register("POST", "scene/count", api.Count(&_table))
	api.Register("POST", "scene/search", api.Search(&_table, nil))
	api.Register("POST", "scene/group", api.Group(&_table, nil))
	api.Register("POST", "scene/import", api.Import(&_table, nil))
	api.Register("POST", "scene/export", api.Export(&_table))

}
