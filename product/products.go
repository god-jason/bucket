package product

import (
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson"
)

var products lib.Map[Product]

func Get(id string) *Product {
	return products.Load(id)
}

func Load(id string) error {
	oid, err := db.ParseObjectId(id)
	if err != nil {
		return err
	}

	var product Product
	err = db.FindById(Bucket, oid, &product)
	if err != nil {
		return err
	}

	products.Store(id, &product)

	return product.Open()
}

func From(product *Product) error {
	products.Store(product.Id.Hex(), product)
	return product.Open()
}

func LoadAll() error {
	var ps []*Product
	err := db.Find(Bucket, bson.D{}, nil, 0, 0, &ps)
	if err != nil {
		return err
	}
	for _, p := range ps {
		err := From(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}
	return nil
}
