package gateway

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "gateway/create", api.Create(&_table, nil))

	api.Register("POST", "gateway/update/:id", api.Update(&_table, nil))

	api.Register("GET", "gateway/delete/:id", api.Delete(&_table, nil))

	api.Register("GET", "gateway/detail/:id", api.Detail(&_table, nil))

	api.Register("GET", "gateway/enable/:id", api.Enable(&_table, nil))

	api.Register("GET", "gateway/disable/:id", api.Disable(&_table, nil))

	api.Register("POST", "gateway/count", api.Count(&_table))

	api.Register("POST", "gateway/search", api.Search(&_table, nil))

	api.Register("POST", "gateway/group", api.Group(&_table, nil))

	api.Register("POST", "gateway/import", api.Import(&_table, nil))

	api.Register("POST", "gateway/export", api.Export(&_table))

}
