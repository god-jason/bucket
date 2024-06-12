package db

import (
	"context"
	"github.com/god-jason/bucket/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var bucket *gridfs.Bucket

func Open() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.GetString(MODULE, "url")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	db = client.Database(config.GetString(MODULE, "database"))

	//默认bucket files.fs files.chunks
	bucket, err = gridfs.NewBucket(db)
	if err != nil {
		_ = Close()
		return err
	}

	return err
}

func Ping() error {
	return db.Client().Ping(context.Background(), nil)
}

func Close() error {
	if db == nil {
		return ErrDisconnect
	}
	err := db.Client().Disconnect(context.TODO())
	db = nil
	bucket = nil
	return err
}
