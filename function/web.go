package function

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/javascript"
)

func init() {
	api.Register("GET", "function/:function/*ext", functionGet)
	api.Register("POST", "function/:function/*ext", functionPost)
}

func functionGet(ctx *gin.Context) {
	f := ctx.Param("function")
	function := functions.Load(f)
	if function == nil || function.Method != "GET" {
		//404
		return
	}

	vm := javascript.Runtime()

	//参数 url, query
	_ = vm.Set("url", ctx.Request.URL.String())
	_ = vm.Set("uri", ctx.Request.URL.RequestURI())
	_ = vm.Set("query", ctx.Request.URL.Query())

	ret, err := vm.RunProgram(function.program)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, ret.Export())
}

func functionPost(ctx *gin.Context) {
	f := ctx.Param("function")
	function := functions.Load(f)
	if function == nil || function.Method != "POST" {
		//404
		return
	}

	//仅支持json数据格式
	var body any
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	vm := javascript.Runtime()

	//参数 url, query, body
	_ = vm.Set("url", ctx.Request.URL.String())
	_ = vm.Set("uri", ctx.Request.URL.RequestURI())
	_ = vm.Set("query", ctx.Request.URL.Query())
	_ = vm.Set("body", body)

	ret, err := vm.RunProgram(function.program)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, ret.Export())
}
