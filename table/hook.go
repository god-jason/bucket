package table

import (
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hook struct {
	BeforeInsert func(doc any) error
	AfterInsert  func(id primitive.ObjectID, doc any) error
	BeforeUpdate func(id primitive.ObjectID, update any) error
	AfterUpdate  func(id primitive.ObjectID, update any) error
	BeforeDelete func(id primitive.ObjectID) error
	AfterDelete  func(id primitive.ObjectID, doc db.Document) error
}
