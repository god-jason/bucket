package alarm

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "alarm/create", api.Create(&_table, nil))
	api.Register("POST", "alarm/update/:id", api.Update(&_table, nil))
	api.Register("GET", "alarm/delete/:id", api.Delete(&_table, nil))
	api.Register("GET", "alarm/detail/:id", api.Detail(&_table, nil))
	api.Register("POST", "alarm/count", api.Count(&_table))
	api.Register("POST", "alarm/search", api.Search(&_table, nil))
	api.Register("POST", "alarm/group", api.Group(&_table,nil))
}
