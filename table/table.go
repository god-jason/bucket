package table

import (
	"errors"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"go.mongodb.org/mongo-driver/bson"
)

var ErrTableNotFound = errors.New("没有表定义")

var tables lib.Map[Table]

func GetTable(name string) (*Table, error) {
	table := tables.Load(name)
	if table == nil {
		return nil, ErrTableNotFound
	}
	return table, nil
}

type Table struct {
	Name string

	//schema
	schema *jsonschema.Schema

	//钩子，CURD
	hooks map[string]interface{}
}

func (t *Table) Aggregate(pipeline interface{}, results *[]Document) error {
	return db.Aggregate(t.Name, pipeline, results)
}

func (t *Table) Insert(doc Document) (id interface{}, err error) {

	//检查
	if t.schema != nil {
		err := t.schema.Validate(doc)
		if err != nil {
			return nil, err
		}
	}

	ret, err := db.InsertOne(t.Name, doc)
	if err != nil {
		return nil, err
	}
	doc["_id"] = ret

	//TODO hook

	return ret, nil
}

func (t *Table) Import(docs []Document) (ids []interface{}, err error) {

	//TODO 没有hook，则直接InsertMany
	if false {
		ds := make([]interface{}, 0, len(docs))
		for _, doc := range docs {
			ds = append(ds, doc)
		}
		return db.InsertMany(t.Name, ds)
	}

	for _, doc := range docs {
		id, err := t.Insert(doc)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (t *Table) Delete(id interface{}) error {
	//TODO 没有hook，则直接Delete
	if false {
		_, err := db.DeleteByID(t.Name, id)
		return err
	}

	var result Document
	err := db.FindOneAndDelete(t.Name, bson.D{{"_id", id}}, &result)
	if err != nil {
		return err
	}
	//TODO hook

	return err
}

func (t *Table) Update(id interface{}, update Document) error {
	//TODO 没有hook，则直接Update
	if false {
		_, err := db.UpdateByID(t.Name, id, update, false)
		return err
	}

	var result Document
	err := db.FindOneAndUpdate(t.Name,
		bson.D{{"_id", id}},
		bson.D{{"$set", update}},
		&result)
	if err != nil {
		return err
	}
	//TODO hook

	return err
}

func (t *Table) Get(id interface{}, result *Document) error {
	err := db.FindOne(t.Name, bson.D{{"_id", id}}, result)
	if err != nil {
		return err
	}
	//TODO hook

	return err
}

func (t *Table) Find(filter interface{}, sort interface{}, skip int64, limit int64, results *[]Document) error {
	return db.Find(t.Name, filter, sort, skip, limit, results)
}

func (t *Table) Count(filter interface{}) (count int64, err error) {
	return db.Count(t.Name, filter)
}

func (t *Table) Drop() error {
	return db.Drop(t.Name)
}

func (t *Table) Distinct(filter interface{}, field string) (values []interface{}, err error) {
	return db.Distinct(t.Name, filter, field)
}
