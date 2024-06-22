package table

import (
	"github.com/god-jason/bucket/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Table struct {
	Name       string                     `json:"name,omitempty"`
	Fields     []*Field                   `json:"fields,omitempty"`
	Schema     *Schema                    `json:"schema,omitempty"`
	Hooks      map[string]*Hook           `json:"hooks,omitempty"`
	TimeSeries *options.TimeSeriesOptions `json:"-"` //时间序列参数
}

func (t *Table) init() error {
	if t.Schema.string != "" {
		err := t.Schema.Compile()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) Aggregate(pipeline any, results any) error {
	return db.Aggregate(t.Name, pipeline, results)
}

func (t *Table) AggregateDocument(pipeline any, results *[]db.Document) error {
	return db.Aggregate(t.Name, pipeline, results)
}

func (t *Table) Insert(doc any) (id primitive.ObjectID, err error) {

	//检查
	if t.Schema != nil {
		err = t.Schema.Validate(doc)
		if err != nil {
			return primitive.NilObjectID, err
		}
	}

	//before insert
	if t.Hooks != nil {
		if hook, ok := t.Hooks["before.insert"]; ok {
			err = hook.Run(map[string]any{"object": doc})
			if err != nil {
				return primitive.NilObjectID, err
			}
		}
	}

	ret, err := db.InsertOne(t.Name, doc)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if d, ok := doc.(map[string]any); ok {
		d["_id"] = ret
	}
	//struct 类型用 反射

	//after insert
	if t.Hooks != nil {
		if hook, ok := t.Hooks["after.insert"]; ok {
			err = hook.Run(map[string]any{"object": doc})
			if err != nil {
				return primitive.NilObjectID, err
			}
		}
	}

	return ret, nil
}

func (t *Table) Import(docs []any) (ids []primitive.ObjectID, err error) {

	//没有hook，则直接InsertMany
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.insert"]; !ok {
			if _, ok := t.Hooks["after.insert"]; !ok {
				return db.InsertMany(t.Name, docs)
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

func (t *Table) ImportDocument(docs []db.Document) (ids []primitive.ObjectID, err error) {

	//没有hook，则直接InsertMany
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.insert"]; !ok {
			if _, ok := t.Hooks["after.insert"]; !ok {
				ds := make([]any, 0, len(docs))
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

func (t *Table) Delete(id primitive.ObjectID) error {
	//没有hook，则直接Delete
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.delete"]; !ok {
			if _, ok := t.Hooks["after.delete"]; !ok {
				_, err := db.DeleteById(t.Name, id)
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

	var result db.Document
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

func (t *Table) Update(id primitive.ObjectID, update any) error {
	//没有hook，则直接Update
	if t.Hooks != nil {
		if _, ok := t.Hooks["before.update"]; !ok {
			if _, ok := t.Hooks["after.update"]; !ok {
				_, err := db.UpdateById(t.Name, id, update, false)
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

	var result db.Document
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

func (t *Table) Get(id primitive.ObjectID, result any) error {
	err := db.FindOne(t.Name, bson.D{{"_id", id}}, result)
	if err != nil {
		return err
	}

	//after get todo 没有太大必要，删掉
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

func (t *Table) GetDocument(id primitive.ObjectID, result *db.Document) error {
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

func (t *Table) Find(filter any, sort any, skip int64, limit int64, results any) error {
	return db.Find(t.Name, filter, sort, skip, limit, results)
}

func (t *Table) FindDocument(filter any, sort any, skip int64, limit int64, results *[]db.Document) error {
	return db.Find(t.Name, filter, sort, skip, limit, results)
}

func (t *Table) Count(filter any) (count int64, err error) {
	return db.Count(t.Name, filter)
}

func (t *Table) Drop() error {
	return db.Drop(t.Name)
}

func (t *Table) Distinct(filter any, field string) (values []any, err error) {
	return db.Distinct(t.Name, filter, field)
}

func (t *Table) DistinctId(filter any) (values []primitive.ObjectID, err error) {
	return db.DistinctId(t.Name, filter)
}
