package base

import "go.mongodb.org/mongo-driver/bson"

var FilterEnabled = bson.D{{"disabled", bson.D{{"$ne", true}}}}
