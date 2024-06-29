package main

import (
	"fmt"
	"github.com/god-jason/bucket/db"
)

func main() {

	a := map[string]any{
		"_id": "123456789012345678901234",
		"abc": []map[string]any{{
			"_id": "123456789012345678901234",
		}},
	}

	fmt.Println(a)

	db.ParseDocumentObjectId(a)

	fmt.Println(a)

	db.StringifyDocumentObjectId(a)

	fmt.Println(a)

}
