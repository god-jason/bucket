package table

import (
	"github.com/dop251/goja"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/pkg/javascript"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Table struct {
	Name   string   `json:"name,omitempty"`
	Fields []*Field `json:"fields,omitempty"`

	Schema string `json:"schema,omitempty"`
	schema *jsonschema.Schema

	Scripts map[string]string `json:"scripts,omitempty"`
	scripts map[string]*goja.Program

	TimeSeries *options.TimeSeriesOptions `json:"-"` //时间序列参数
	Hook       *Hook                      `json:"-"`
}

func (t *Table) init() (err error) {

	//JSONSchema
	if t.Schema != "" {
		t.schema, err = compiler.Compile(t.Schema)
		if err != nil {
			return err
		}
	}

	//编译脚本
	t.scripts = make(map[string]*goja.Program)
	for hook, str := range t.Scripts {
		t.scripts[hook], err = javascript.Compile(str)
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
	if t.schema != nil {
		err = t.schema.Validate(doc)
		if err != nil {
			return primitive.NilObjectID, errors.Wrap(err)
		}
	}

	//before insert
	if t.Hook != nil && t.Hook.BeforeInsert != nil {
		err := t.Hook.BeforeInsert(&doc)
		if err != nil {
			return primitive.NilObjectID, errors.Wrap(err)
		}
	}
	if hook, ok := t.scripts["before.insert"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("object", doc)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return primitive.NilObjectID, errors.Wrap(err)
		}
	}

	//补充创建时间
	if mp, ok := doc.(map[string]interface{}); ok {
		for _, f := range t.Fields {
			if f.Created || f.Updated {
				if _, ok := mp[f.Name]; !ok {
					mp[f.Name] = time.Now()
				}
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
	if t.Hook != nil && t.Hook.AfterInsert != nil {
		err := t.Hook.AfterInsert(ret, &doc)
		if err != nil {
			return primitive.NilObjectID, errors.Wrap(err)
		}
	}
	if hook, ok := t.scripts["after.insert"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", ret)
		_ = vm.Set("object", doc)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return primitive.NilObjectID, errors.Wrap(err)
		}
	}

	return ret, nil
}

func (t *Table) Import(docs []any) (ids []primitive.ObjectID, err error) {

	//没有hook，则直接InsertMany
	if t.Hook == nil || t.Hook.BeforeInsert == nil && t.Hook.AfterInsert == nil {
		if _, ok := t.scripts["before.insert"]; !ok {
			if _, ok := t.scripts["after.insert"]; !ok {
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
	if t.Hook == nil || t.Hook.BeforeInsert == nil && t.Hook.AfterInsert == nil {
		if _, ok := t.scripts["before.insert"]; !ok {
			if _, ok := t.scripts["after.insert"]; !ok {
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
	//before delete
	if t.Hook != nil && t.Hook.BeforeDelete != nil {
		err := t.Hook.BeforeDelete(id)
		if err != nil {
			return errors.Wrap(err)
		}
	}
	if hook, ok := t.scripts["before.delete"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_, err := vm.RunProgram(hook)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	var result db.Document
	err := db.FindOneAndDelete(t.Name, bson.D{{"_id", id}}, &result)
	if err != nil {
		return err
	}

	//把删除保存到修改历史表
	_, _ = db.InsertOne(t.Name+".deleted", result)

	//after delete
	if t.Hook != nil && t.Hook.AfterDelete != nil {
		err := t.Hook.AfterDelete(id, result)
		if err != nil {
			return errors.Wrap(err)
		}
	}
	if hook, ok := t.scripts["after.delete"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_ = vm.Set("object", result)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return err
}

func (t *Table) Update(id primitive.ObjectID, update any) error {

	//before update
	if t.Hook != nil && t.Hook.BeforeUpdate != nil {
		err := t.Hook.BeforeUpdate(id, update)
		if err != nil {
			return errors.Wrap(err)
		}
	}
	if hook, ok := t.scripts["before.update"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_ = vm.Set("change", update)
		_, err := vm.RunProgram(hook)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	//补充更新时间
	if mp, ok := update.(map[string]interface{}); ok {
		for _, f := range t.Fields {
			if f.Updated {
				if _, ok := mp[f.Name]; !ok {
					mp[f.Name] = time.Now()
				}
			}
		}
	}

	var result db.Document
	err := db.FindOneAndUpdate(t.Name, bson.D{{"_id", id}}, bson.D{{"$set", update}}, &result)
	if err != nil {
		return err
	}

	//把差异保存到修改历史表
	_, _ = db.InsertOne(t.Name+".change", bson.M{"object_id": id, "base": result, "change": update})

	//after update
	if t.Hook != nil && t.Hook.AfterUpdate != nil {
		err := t.Hook.AfterUpdate(id, update, result)
		if err != nil {
			return errors.Wrap(err)
		}
	}
	if hook, ok := t.scripts["after.update"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_ = vm.Set("change", update)
		_ = vm.Set("base", result)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return err
}

func (t *Table) Get(id primitive.ObjectID, result any) error {
	return db.FindOne(t.Name, bson.D{{"_id", id}}, result)
}

func (t *Table) GetDocument(id primitive.ObjectID, result *db.Document) error {
	return db.FindOne(t.Name, bson.D{{"_id", id}}, result)
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
