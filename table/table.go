package table

import (
	"errors"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
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
	Name   string           `json:"name,omitempty"`
	Fields []*Field         `json:"fields,omitempty"`
	Schema *Schema          `json:"schema,omitempty"`
	Hooks  map[string]*Hook `json:"hooks,omitempty"`
}

func (t *Table) Aggregate(pipeline interface{}, results *[]Document) error {
	return db.Aggregate(t.Name, pipeline, results)
}

func (t *Table) Insert(doc Document) (id interface{}, err error) {

	//检查
	if t.Schema != nil {
		err = t.Schema.Validate(doc)
		if err != nil {
			return nil, err
		}
	}

	//before insert
	if t.Hooks != nil {
		if hook, ok := t.Hooks["before.insert"]; ok {
			err = hook.Run(map[string]any{"object": doc})
			if err != nil {
				return nil, err
			}
		}
	}

	ret, err := db.InsertOne(t.Name, doc)
	if err != nil {
		return nil, err
	}
	doc["_id"] = ret

	//after insert
	if t.Hooks != nil {
		if hook, ok := t.Hooks["after.insert"]; ok {
			err = hook.Run(map[string]any{"object": doc})
			if err != nil {
				return nil, err
			}
		}
	}

	return ret, nil
}

func (t *Table) Import(docs []Document) (ids []interface{}, err error) {

	//没有hook，则直接InsertMany
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.insert"]; !ok {
			if _, ok := t.Hooks["after.insert"]; !ok {
				ds := make([]interface{}, 0, len(docs))
				for _, doc := range docs {
					ds = append(ds, doc)
				}
				return db.InsertMany(t.Name, ds)
			}
		}
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
	//没有hook，则直接Delete
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.delete"]; !ok {
			if _, ok := t.Hooks["after.delete"]; !ok {
				_, err := db.DeleteByID(t.Name, id)
				return err
			}
		}
	}

	//before delete
	if t.Hooks != nil {
		if hook, ok := t.Hooks["before.delete"]; ok {
			err := hook.Run(map[string]any{"id": id})
			if err != nil {
				return err
			}
		}
	}

	var result Document
	err := db.FindOneAndDelete(t.Name, bson.D{{"_id", id}}, &result)
	if err != nil {
		return err
	}

	//after delete
	if t.Hooks != nil {
		if hook, ok := t.Hooks["after.delete"]; ok {
			err := hook.Run(map[string]any{"id": id, "object": result})
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (t *Table) Update(id interface{}, update Document) error {
	//没有hook，则直接Update
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.update"]; !ok {
			if _, ok := t.Hooks["after.update"]; !ok {
				_, err := db.UpdateByID(t.Name, id, update, false)
				return err
			}
		}
	}

	//before update
	if t.Hooks != nil {
		if hook, ok := t.Hooks["before.update"]; ok {
			err := hook.Run(map[string]any{"id": id, "update": update})
			if err != nil {
				return err
			}
		}
	}

	var result Document
	err := db.FindOneAndUpdate(t.Name,
		bson.D{{"_id", id}},
		bson.D{{"$set", update}},
		&result)
	if err != nil {
		return err
	}

	//after update
	if t.Hooks != nil {
		if hook, ok := t.Hooks["after.update"]; ok {
			err := hook.Run(map[string]any{"id": id, "update": update, "object": result})
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (t *Table) Get(id interface{}, result *Document) error {
	err := db.FindOne(t.Name, bson.D{{"_id", id}}, result)
	if err != nil {
		return err
	}

	//after get
	if t.Hooks != nil {
		if hook, ok := t.Hooks["after.get"]; ok {
			err := hook.Run(map[string]any{"id": id, "object": result})
			if err != nil {
				return err
			}
		}
	}

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
