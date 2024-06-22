package table

import (
	"go.mongodb.org/mongo-driver/bson"
)

func BatchLoad[T any](t *Table, filter bson.D, page int, f func(t T) error) error {
	var skip int = 0
	for {
		var ts []T
		err := t.Find(filter, nil, int64(skip), int64(page), &ts)
		if err != nil {
			return err
		}
		ln := len(ts)
		if ln == 0 {
			break
		}
		for _, t := range ts {
			err := f(t)
			if err != nil {
				return err
			}
		}
		if ln < page {
			break
		}
		skip += ln
	}
	return nil
}
