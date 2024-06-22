package base

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Executor interface {
	Execute(actions []*Action)
}

type DeviceContainer interface {
	Devices(productId primitive.ObjectID) (ids []primitive.ObjectID, err error)
}
