package table

import (
	"github.com/god-jason/bucket/db"
)

type Hook struct {
	BeforeInsert func(doc any) error
	AfterInsert  func(id string, doc any) error
	BeforeUpdate func(id string, update any) error
	AfterUpdate  func(id string, update any, base db.Document) error
	BeforeDelete func(id string) error
	AfterDelete  func(id string, doc db.Document) error
}
