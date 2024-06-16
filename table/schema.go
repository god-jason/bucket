package table

import "github.com/santhosh-tekuri/jsonschema/v6"

var compiler *jsonschema.Compiler

func init() {
	compiler = jsonschema.NewCompiler()
}

type Schema struct {
	string
	schema *jsonschema.Schema
}

func (s *Schema) Compile() (err error) {
	if s.string != "" {
		s.schema, err = compiler.Compile(s.string)
	}
	return
}

func (s *Schema) Validate(doc any) error {
	if s.schema != nil {
		return s.schema.Validate(doc)
	}
	return nil
}
