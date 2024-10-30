package api

import "github.com/god-jason/bucket/table"

// 表相关接口只能放这里了，否则会import circle: api->table->api
func init() {

	Register("GET", "table/list", table.ApiList)
	Register("GET", "table/:table/file/:file", table.ApiManifest)
	Register("POST", "table/:table/file/:file", table.ApiManifestUpdate)

	Register("POST", "table/:table/count", table.ApiCount)
	Register("POST", "table/:table/create", table.ApiCreate)
	Register("POST", "table/:table/update/:id", table.ApiUpdate)
	Register("GET", "table/:table/delete/:id", table.ApiDelete)
	Register("GET", "table/:table/detail/:id", table.ApiDetail)
	Register("POST", "table/:table/group", table.ApiGroup)
	Register("POST", "table/:table/search", table.ApiSearch)
	Register("POST", "table/:table/import", table.ApiImport)
	Register("POST", "table/:table/export", table.ApiExport)
}
