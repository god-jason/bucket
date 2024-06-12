package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
)

func Upload(filename string, metadata interface{}) (*gridfs.UploadStream, error) {
	if bucket == nil {
		return nil, ErrDisconnect
	}
	opts := options.GridFSUpload().SetMetadata(metadata)
	return bucket.OpenUploadStream(filename, opts)
}

func UploadFrom(filename string, metadata interface{}, reader io.Reader) (id interface{}, err error) {
	if bucket == nil {
		return nil, ErrDisconnect
	}
	opts := options.GridFSUpload().SetMetadata(metadata)
	return bucket.UploadFromStream(filename, reader, opts)
}

func Download(filename string) (*gridfs.DownloadStream, error) {
	if bucket == nil {
		return nil, ErrDisconnect
	}
	return bucket.OpenDownloadStreamByName(filename)
}

func DownloadTo(filename string, writer io.Writer) (int64, error) {
	if bucket == nil {
		return 0, ErrDisconnect
	}
	return bucket.DownloadToStreamByName(filename, writer)
}

func DownloadByID(id interface{}) (*gridfs.DownloadStream, error) {
	if bucket == nil {
		return nil, ErrDisconnect
	}
	return bucket.OpenDownloadStream(id)
}

func DownloadToByID(id interface{}, writer io.Writer) (int64, error) {
	if bucket == nil {
		return 0, ErrDisconnect
	}
	return bucket.DownloadToStream(id, writer)
}

func Rename(id interface{}, filename string) error {
	if bucket == nil {
		return ErrDisconnect
	}
	return bucket.Rename(id, filename)
}

func Remove(id interface{}) error {
	if bucket == nil {
		return ErrDisconnect
	}
	return bucket.Delete(id)
}

func FindFile(filter interface{}, sort interface{}, skip int32, limit int32, results interface{}) error {
	if bucket == nil {
		return ErrDisconnect
	}

	opts := options.GridFSFind()
	if sort != nil {
		opts.SetSort(sort)
	}
	if skip > 0 {
		opts.SetSkip(skip)
	}
	if limit > 0 {
		opts.SetLimit(limit)
	}

	ret, err := bucket.Find(filter, opts)
	if err != nil {
		return err
	}
	return ret.All(context.Background(), results)
}
