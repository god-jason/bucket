package aggregate

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "aggregate/create", api.Create(&_table, nil))
	api.Register("POST", "aggregate/update/:id", api.Update(&_table, nil))
	api.Register("GET", "aggregate/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "aggregate/detail/:id", api.Detail(&_table, nil))
	api.Register("POST", "aggregate/search", api.Search(&_table, nil))
	api.Register("POST", "aggregate/import", api.Import(&_table, nil))
	api.Register("POST", "aggregate/export", api.Export(&_table))
	api.Register("POST", "aggregate/count", api.Count(&_table))
}
