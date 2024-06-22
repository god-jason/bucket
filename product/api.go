package product

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "product/create", api.Create(&_table, nil))
	api.Register("POST", "product/update/:id", api.Update(&_table, nil))
	api.Register("GET", "product/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "product/detail/:id", api.Detail(&_table, nil))
	api.Register("POST", "product/count", api.Count(&_table))
	api.Register("POST", "product/search", api.Search(&_table, nil))
	api.Register("POST", "product/group", api.Group(&_table, nil))
	api.Register("POST", "product/import", api.Import(&_table, nil))
	api.Register("POST", "product/export", api.Export(&_table))
}
