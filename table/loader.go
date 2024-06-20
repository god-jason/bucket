package table

import (
	"go.mongodb.org/mongo-driver/bson"
)

func BatchLoad[T any](t *Table, filter bson.D, page int64, f func(t T)) {
	var skip int64 = 0
	for {
		var ts []T
		err := t.Find(filter, bson.D{{}}, skip, page, ts)
		if err != nil {
			return
		}
		if len(ts) == 0 {
			break
		}
		for _, t := range ts {
			f(t)
		}
	}
}
