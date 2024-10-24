package table

import (
	"github.com/god-jason/bucket/mongodb"
)

type Hook struct {
	BeforeInsert func(doc any) error
	AfterInsert  func(id string, doc any) error
	BeforeUpdate func(id string, update any) error
	AfterUpdate  func(id string, update any, base mongodb.Document) error
	BeforeDelete func(id string) error
	AfterDelete  func(id string, doc mongodb.Document) error
}
