package table

import (
	"github.com/santhosh-tekuri/jsonschema/v6"
)

var compiler *jsonschema.Compiler

func init() {
	compiler = jsonschema.NewCompiler()
}
