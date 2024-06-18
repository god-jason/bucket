package function

import "github.com/dop251/goja"

type Function struct {
	Name   string
	Script string
	Method string //get post

	program *goja.Program
}

func (f *Function) Compile() (err error) {
	f.program, err = goja.Compile(f.Name, f.Script, false)
	return
}
