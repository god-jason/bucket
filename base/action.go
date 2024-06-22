package base

import "go.mongodb.org/mongo-driver/bson/primitive"

type Action struct {
	ProductId  primitive.ObjectID `json:"product_id,omitempty" bson:"product_id"`
	DeviceId   primitive.ObjectID `json:"device_id,omitempty" bson:"device_id"`
	Name       string             `json:"action"`
	Parameters map[string]any     `json:"parameters,omitempty"`
}
