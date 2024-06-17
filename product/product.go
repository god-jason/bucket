package product

import "github.com/god-jason/bucket/lib"

type Product struct {
	Name       string
	Type       string //泛类型，比如：电表，水表
	Properties []*Property
}

type Property struct {
	Name        string
	Label       string
	Type        string //bool string number array object
	Default     any
	Writable    bool
	Historical  bool
	Aggregators []*Aggregator

	//Children *Property
}

type Aggregator struct {
	//Table  string        //默认 bucket.aggregate
	//Period time.Duration //1h
	Type string //inc sum count avg last first max min
	As   string
}

var products lib.Map[Product]

func Get(id string) *Product {
	return products.Load(id)
}
