package timer

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "timer/create", api.Create(&_table, nil))
	api.Register("POST", "timer/update/:id", api.Update(&_table, nil))
	api.Register("GET", "timer/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "timer/detail/:id", api.Detail(&_table, nil))
	api.Register("GET", "timer/enable/:id", api.Update(&_table, Load))
	api.Register("GET", "timer/disable/:id", api.Delete(&_table, Unload))
	api.Register("GET", "timer/execute/:id", api.Operator(Execute))
	api.Register("POST", "timer/count", api.Count(&_table))
	api.Register("POST", "timer/search", api.Search(&_table, nil))
	api.Register("POST", "timer/group", api.Group(&_table, nil))
	api.Register("POST", "timer/import", api.Import(&_table, nil))
	api.Register("POST", "timer/export", api.Export(&_table))

}
