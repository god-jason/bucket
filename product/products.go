package product

import (
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var products lib.Map[Product]

func Get(id string) *Product {
	return products.Load(id)
}

func From(v *Product) (err error) {
	tt := products.LoadAndStore(v.Id.Hex(), v)
	if tt != nil {
		_ = tt.Close()
	}
	return v.Open()
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
	t := products.LoadAndDelete(id.Hex())
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Product](&_table, nil, 100, func(t *Product) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}
