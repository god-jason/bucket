package table

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/url"
)

type JavaScriptRuntime *goja.Runtime
type JavaScriptProgram *goja.Program

var req require.Registry

func CreateJavaScriptRuntime() *goja.Runtime {
	runtime := goja.New()

	//支持require
	req.Enable(runtime)

	//支持console.log
	console.Enable(runtime)

	//支持url解析
	url.Enable(runtime)

	return runtime
}

func CompileJavaScript(str string) (*goja.Program, error) {
	return goja.Compile("", str, false)
}
