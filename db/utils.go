package db

import (
	"github.com/god-jason/bucket/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

func ParseObjectId(id any) (primitive.ObjectID, error) {
	switch val := id.(type) {
	case primitive.ObjectID:
		return val, nil
	case string:
		return primitive.ObjectIDFromHex(val)
	default:
		return primitive.NilObjectID, errors.New("invalid object id")
	}
}

func ConvertObjectId(doc any) {
	switch val := doc.(type) {
	case map[string]any:
		for k, v := range val {
			if strings.HasSuffix(k, "_id") {
				val[k], _ = ParseObjectId(v)
				continue
			}
			ConvertObjectId(v)
		}
	case []any:
		for _, v := range val {
			ConvertObjectId(v)
		}
	}
}
