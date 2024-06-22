package main

import (
	"encoding/json"
	"github.com/god-jason/bucket/log"
)

type A struct {
	string
	B int
}

type C struct {
	A A
	C int
}

func main() {
	v := &C{
		A: A{string: "213", B: 1},
		C: 2,
	}

	vv, _ := json.Marshal(v)

	log.Println(v, string(vv))
}
