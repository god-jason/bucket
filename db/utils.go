package db

import (
	"github.com/god-jason/bucket/pkg/exception"
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
		return primitive.NilObjectID, exception.New("invalid object id")
	}
}

func StringifyObjectId(id any) (string, error) {
	switch val := id.(type) {
	case primitive.ObjectID:
		return val.Hex(), nil
	case string:
		return val, nil
	default:
		return "", exception.New("invalid object id")
	}
}

func ParseDocumentObjectId(doc any) {
	switch val := doc.(type) {
	case map[string]any:
		for k, v := range val {
			if strings.HasSuffix(k, "_id") {
				val[k], _ = ParseObjectId(v)
				continue
			}
			ParseDocumentObjectId(v)
		}
	case []any:
		for _, v := range val {
			ParseDocumentObjectId(v)
		}
	}
}

func StringifyDocumentObjectId(doc any) {
	switch val := doc.(type) {
	case map[string]any:
		for k, v := range val {
			if strings.HasSuffix(k, "_id") {
				val[k], _ = StringifyObjectId(v)
				continue
			}
			StringifyDocumentObjectId(v)
		}
	case []any:
		for _, v := range val {
			StringifyDocumentObjectId(v)
		}
	}
}
