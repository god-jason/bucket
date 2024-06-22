package product

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var products lib.Map[Product]

func Get(id string) *Product {
	return products.Load(id)
}

func From(v *Product) (err error) {
	products.Store(v.Id.Hex(), v)
	return nil
}

func Load(id primitive.ObjectID) error {
	var product Product
	err := _table.Get(id, &product)
	if err != nil {
		return err
	}
	return From(&product)
}

func Unload(id primitive.ObjectID) error {
	products.Delete(id.Hex())
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Product](&_table, base.FilterEnabled, 100, func(t *Product) error {
		//并行加载
		_ = From(t)
		//products.Store(t.Id.Hex(), t)
		return nil
	})
}
