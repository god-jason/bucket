package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrDisconnect = errors.New("数据库未连接")

func Aggregate(col string, pipeline interface{}, results interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	cursor, err := db.Collection(col).Aggregate(context.Background(), pipeline)
	if err != nil {
		return err
	}
	return cursor.All(context.Background(), results)
}

func BulkWrite(col string, models []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	return db.Collection(col).BulkWrite(context.Background(), models)
}

func InsertOne(col string, doc interface{}) (id primitive.ObjectID, err error) {
	if db == nil {
		return _id, ErrDisconnect
	}
	ret, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return _id, err
	}
	return ParseObjectId(ret.InsertedID)
}

func InsertMany(col string, docs []interface{}) (ids []interface{}, err error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	ret, err := db.Collection(col).InsertMany(context.Background(), docs)
	if err != nil {
		return nil, err
	}
	return ret.InsertedIDs, nil
}

func DeleteOne(col string, filter interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(col).DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func DeleteMany(col string, filter interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(col).DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func DeleteById(col string, id primitive.ObjectID) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(col).DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func ReplaceOne(col string, filter interface{}, result interface{}, upsert bool) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	opts := options.Replace().SetUpsert(upsert)
	ret, err := db.Collection(col).ReplaceOne(context.Background(), filter, result, opts)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func UpdateOne(col string, filter interface{}, update interface{}, upsert bool) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	opts := options.Update().SetUpsert(upsert)
	ret, err := db.Collection(col).UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func UpdateMany(col string, filter interface{}, update interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(col).UpdateMany(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func UpdateById(col string, id primitive.ObjectID, update interface{}, upsert bool) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	opts := options.Update().SetUpsert(upsert)
	ret, err := db.Collection(col).UpdateByID(context.Background(), id, update, opts)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func Find(col string, filter interface{}, sort interface{}, skip int64, limit int64, results interface{}) error {
	if db == nil {
		return ErrDisconnect
	}

	opts := options.Find()
	if sort != nil {
		opts.SetSort(sort)
	}
	if skip > 0 {
		opts.SetSkip(skip)
	}
	if limit > 0 {
		opts.SetLimit(limit)
	}

	ret, err := db.Collection(col).Find(context.Background(), filter, opts)
	if err != nil {
		return err
	}
	return ret.All(context.Background(), results)
}

func FindOne(col string, filter interface{}, result interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(col).FindOne(context.Background(), filter)
	return ret.Decode(result)
}

func FindOneAndDelete(col string, filter interface{}, raw interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(col).FindOneAndDelete(context.Background(), filter)
	return ret.Decode(raw)
}

func FindOneAndUpdate(col string, filter interface{}, update interface{}, raw interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(col).FindOneAndUpdate(context.Background(), filter, update)
	return ret.Decode(raw)
}

func FindOneAndReplace(col string, filter interface{}, replace interface{}, raw interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(col).FindOneAndUpdate(context.Background(), filter, replace)
	return ret.Decode(raw)
}

func FindById(col string, id primitive.ObjectID, result interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(col).FindOne(context.Background(), bson.D{{"_id", id}})
	return ret.Decode(result)
}

func Count(col string, filter interface{}) (count int64, err error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	return db.Collection(col).CountDocuments(context.Background(), filter)
}

func Drop(col string) error {
	if db == nil {
		return ErrDisconnect
	}
	return db.Collection(col).Drop(context.Background())
}

func Distinct(col string, filter interface{}, field string) (values []interface{}, err error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	return db.Collection(col).Distinct(context.Background(), field, filter)
}
