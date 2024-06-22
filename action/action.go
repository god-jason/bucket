package action

import "go.mongodb.org/mongo-driver/bson/primitive"

type Action struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ProductId  primitive.ObjectID `json:"product_id,omitempty" bson:"product_id"`
	DeviceId   primitive.ObjectID `json:"device_id,omitempty" bson:"device_id"`
	ProjectId  primitive.ObjectID `json:"project_id,omitempty" bson:"project_id"`
	SpaceId    primitive.ObjectID `json:"space_id,omitempty" bson:"space_id"`
	Name       string             `json:"action"`
	Parameters map[string]any     `json:"parameters,omitempty"`
	Result     string             `json:"result,omitempty"`
	Return     map[string]any     `json:"return,omitempty"`
}
