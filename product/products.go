package product

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/table"
)

var products lib.Map[Product]

func Get(id string) *Product {
	return products.Load(id)
}

func From(v *Product) (err error) {
	products.Store(v.Id, v)
	return nil
}

func Load(id string) error {
	var product Product
	has, err := _table.Get(id, &product)
	if err != nil {
		return err
	}
	if !has {
		return exception.New("找不到记录")
	}
	return From(&product)
}

func Unload(id string) error {
	products.Delete(id)
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
