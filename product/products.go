package product

import "github.com/god-jason/bucket/lib"

var products lib.Map[Product]

func Get(id string) *Product {
	return products.Load(id)
}
