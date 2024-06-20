package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
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

func ConvertObjectId(doc any) {
	switch val := doc.(type) {
	case map[string]interface{}:
		for k, v := range val {
			if strings.HasSuffix(k, "_id") {
				val[k], _ = ParseObjectId(v)
				continue
			}
			ConvertObjectId(v)
		}
	case []interface{}:
		for _, v := range val {
			ConvertObjectId(v)
		}
	}
}
