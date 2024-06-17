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

func Aggregate(tab string, pipeline interface{}, results interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	cursor, err := db.Collection(tab).Aggregate(context.Background(), pipeline)
	if err != nil {
		return err
	}
	return cursor.All(context.Background(), results)
}

func BulkWrite(tab string, models []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	return db.Collection(tab).BulkWrite(context.Background(), models)
}

func InsertOne(tab string, doc interface{}) (id primitive.ObjectID, err error) {
	if db == nil {
		return _id, ErrDisconnect
	}
	ret, err := db.Collection(tab).InsertOne(context.Background(), doc)
	if err != nil {
		return _id, err
	}
	return ParseObjectId(ret.InsertedID)
}

func InsertMany(tab string, docs []interface{}) (ids []interface{}, err error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	ret, err := db.Collection(tab).InsertMany(context.Background(), docs)
	if err != nil {
		return nil, err
	}
	return ret.InsertedIDs, nil
}

func DeleteOne(tab string, filter interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(tab).DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func DeleteMany(tab string, filter interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(tab).DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func DeleteById(tab string, id primitive.ObjectID) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(tab).DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func ReplaceOne(tab string, filter interface{}, result interface{}, upsert bool) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	opts := options.Replace().SetUpsert(upsert)
	ret, err := db.Collection(tab).ReplaceOne(context.Background(), filter, result, opts)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func UpdateOne(tab string, filter interface{}, update interface{}, upsert bool) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	opts := options.Update().SetUpsert(upsert)
	ret, err := db.Collection(tab).UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func UpdateMany(tab string, filter interface{}, update interface{}) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	ret, err := db.Collection(tab).UpdateMany(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func UpdateById(tab string, id primitive.ObjectID, update interface{}, upsert bool) (int64, error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	opts := options.Update().SetUpsert(upsert)
	ret, err := db.Collection(tab).UpdateByID(context.Background(), id, update, opts)
	if err != nil {
		return 0, err
	}
	return ret.ModifiedCount, nil
}

func Find(tab string, filter interface{}, sort interface{}, skip int64, limit int64, results interface{}) error {
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

	ret, err := db.Collection(tab).Find(context.Background(), filter, opts)
	if err != nil {
		return err
	}
	return ret.All(context.Background(), results)
}

func FindOne(tab string, filter interface{}, result interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(tab).FindOne(context.Background(), filter)
	return ret.Decode(result)
}

func FindOneAndDelete(tab string, filter interface{}, raw interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(tab).FindOneAndDelete(context.Background(), filter)
	return ret.Decode(raw)
}

func FindOneAndUpdate(tab string, filter interface{}, update interface{}, raw interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(tab).FindOneAndUpdate(context.Background(), filter, update)
	return ret.Decode(raw)
}

func FindOneAndReplace(tab string, filter interface{}, replace interface{}, raw interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(tab).FindOneAndUpdate(context.Background(), filter, replace)
	return ret.Decode(raw)
}

func FindById(tab string, id primitive.ObjectID, result interface{}) error {
	if db == nil {
		return ErrDisconnect
	}
	ret := db.Collection(tab).FindOne(context.Background(), bson.D{{"_id", id}})
	return ret.Decode(result)
}

func Count(tab string, filter interface{}) (count int64, err error) {
	if db == nil {
		return 0, ErrDisconnect
	}
	return db.Collection(tab).CountDocuments(context.Background(), filter)
}

func Drop(tab string) error {
	if db == nil {
		return ErrDisconnect
	}
	return db.Collection(tab).Drop(context.Background())
}

func Distinct(tab string, filter interface{}, field string) (values []interface{}, err error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	return db.Collection(tab).Distinct(context.Background(), field, filter)
}

func CreateIndex(tab string, keys []string) error {
	if db == nil {
		return ErrDisconnect
	}
	var ks bson.D
	for _, k := range keys {
		ks = append(ks, bson.E{Key: k, Value: 1}) //未支持降序索引
	}
	_, err := db.Collection(tab).Indexes().CreateOne(context.Background(), mongo.IndexModel{Keys: ks})
	return err
}
