package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _id primitive.ObjectID

func ParseObjectId(id any) (primitive.ObjectID, error) {
	switch val := id.(type) {
	case primitive.ObjectID:
		return val, nil
	case string:
		return primitive.ObjectIDFromHex(val)
	default:
		return _id, errors.New("invalid object id")
	}
}

func EmptyObjectId() primitive.ObjectID {
	return _id
}
