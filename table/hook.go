package table

import (
	"github.com/dop251/goja"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/pkg/javascript"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JavaScriptHook struct {
	string
	program *goja.Program
}

func (h *JavaScriptHook) Compile() (err error) {
	if h.string != "" {
		h.program, err = javascript.Compile(h.string)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *JavaScriptHook) Run(context map[string]any) error {
	if h.program != nil {
		runtime := javascript.Runtime()
		err := runtime.Set("context", context)
		if err != nil {
			return err
		}

		_, err = runtime.RunProgram(h.program)
		//打印返回值？？？
		if err != nil {
			return errors.Wrap(err)
		}
	}
	return nil
}

type NativeHook struct {
	BeforeInsert func(doc any) error
	AfterInsert  func(id primitive.ObjectID, doc any) error
	BeforeUpdate func(id primitive.ObjectID, update any) error
	AfterUpdate  func(id primitive.ObjectID, update any) error
	BeforeDelete func(id primitive.ObjectID) error
	AfterDelete  func(id primitive.ObjectID, doc db.Document) error
	AfterGet     func(id primitive.ObjectID, doc db.Document) error
}
