package function

import (
	"github.com/dop251/goja"
	"github.com/god-jason/bucket/pkg/errors"
)

type Function struct {
	Name   string
	Script string
	Method string //get post

	program *goja.Program
}

func (f *Function) Compile() (err error) {
	f.program, err = goja.Compile(f.Name, f.Script, false)
	return errors.Wrap(err)
}
