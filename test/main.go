package main

import (
	"fmt"
	"github.com/god-jason/bucket/mongodb"
)

func main() {

	a := map[string]any{
		"_id": "123456789012345678901234",
		"abc": []map[string]any{{
			"_id": "123456789012345678901234",
		}},
	}

	fmt.Println(a)

	mongodb.ParseDocumentObjectId(a)

	fmt.Println(a)

	mongodb.StringifyDocumentObjectId(a)

	fmt.Println(a)

}
