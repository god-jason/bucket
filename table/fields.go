package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/types"
)

func init() {
	api.Register("GET", "table/:table/fields", apiFields)
	api.Register("POST", "table/:table/fields", apiFields)
}

func apiFields(ctx *gin.Context) {
	//table, err := Get(ctx.Param("table"))
	//if err != nil {
	//	curd.Error(ctx, err)
	//	return
	//}

	var fields []types.SmartField

	api.OK(ctx, fields)
}
