package space

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Space struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	Name      string             `json:"name"`
	Disabled  bool               `json:"disabled"`
}
