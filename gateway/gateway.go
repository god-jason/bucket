package gateway

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gateway struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
}
