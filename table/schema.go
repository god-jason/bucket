package table

import (
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

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
		err = errors.Wrap(err)
	}
	return
}

func (s *Schema) Validate(doc any) error {
	if s.schema != nil {
		err := s.schema.Validate(doc)
		return errors.Wrap(err)
	}
	return nil
}
