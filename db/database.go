package db

import (
	"context"
	"github.com/god-jason/bucket/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var bucket *gridfs.Bucket

func Open() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(config.GetString(MODULE, "url")).
		SetServerAPIOptions(serverAPI)

	//鉴权
	auth := config.GetString(MODULE, "auth")
	username := config.GetString(MODULE, "username")
	password := config.GetString(MODULE, "password")
	if auth != "" && username != "" && password != "" {
		opts.SetAuth(options.Credential{
			AuthSource: auth,
			Username:   username,
			Password:   password,
		})
	}

	//连接
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

func Database() (*mongo.Database, error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	return db, nil
}

func CreateTable(name string, opts *options.CreateCollectionOptions) error {
	if db == nil {
		return ErrDisconnect
	}
	return db.CreateCollection(context.Background(), name, opts)
}

func Tables() ([]string, error) {
	if db == nil {
		return nil, ErrDisconnect
	}
	return db.ListCollectionNames(context.Background(), bson.D{{}})
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
