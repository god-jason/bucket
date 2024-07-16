package table

import (
	"errors"
	"github.com/dop251/goja"
	"github.com/god-jason/bucket/accumulate"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/pkg/javascript"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var NotFound = errors.New("找不到记录")

type Table struct {
	Name   string   `json:"name,omitempty"`
	Fields []*Field `json:"fields,omitempty"`

	//Json Schema
	Schema string `json:"schema,omitempty"`
	schema *jsonschema.Schema

	//脚本
	Scripts map[string]string `json:"scripts,omitempty"`
	scripts map[string]*goja.Program

	//累加器
	Accumulations []*accumulate.Accumulation `json:"accumulations,omitempty"`

	TimeSeries *options.TimeSeriesOptions `json:"-"` //时间序列参数
	Hook       *Hook                      `json:"-"`

	//快照
	SnapshotOptions *SnapshotOptions `json:"snapshot,omitempty"`
	//备份
	DeleteOptions *DeleteOptions `json:"delete,omitempty"`
	//历史
	HistoryOptions *HistoryOptions `json:"history,omitempty"`
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

	//初始化累加器
	for _, a := range t.Accumulations {
		err = a.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Table) Aggregate(pipeline any, results any) error {
	db.ParseDocumentObjectId(pipeline)
	return db.Aggregate(t.Name, pipeline, results)
}

func (t *Table) AggregateDocument(pipeline any, results *[]db.Document) error {
	db.ParseDocumentObjectId(pipeline)
	return db.Aggregate(t.Name, pipeline, results)
}

func (t *Table) Insert(doc any) (id string, err error) {

	//检查
	if t.schema != nil {
		err = t.schema.Validate(doc)
		if err != nil {
			return "", exception.Wrap(err)
		}
	}

	//before insert
	if t.Hook != nil && t.Hook.BeforeInsert != nil {
		err := t.Hook.BeforeInsert(&doc)
		if err != nil {
			return "", exception.Wrap(err)
		}
	}
	if hook, ok := t.scripts["before.insert"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("object", doc)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return "", exception.Wrap(err)
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
	//todo 用反射检查 struct

	ret, err := db.InsertOne(t.Name, doc)
	if err != nil {
		return "", err
	}
	if d, ok := doc.(map[string]any); ok {
		d["_id"] = ret.Hex()
	}
	//struct 类型用 反射

	//after insert
	if t.Hook != nil && t.Hook.AfterInsert != nil {
		err := t.Hook.AfterInsert(ret.Hex(), &doc)
		if err != nil {
			return "", exception.Wrap(err)
		}
	}
	if hook, ok := t.scripts["after.insert"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", ret)
		_ = vm.Set("object", doc)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return "", exception.Wrap(err)
		}
	}

	//累加器
	for _, a := range t.Accumulations {
		ret, err := a.Evaluate(doc)
		if err != nil {
			log.Error(err)
			continue
		}

		if len(ret.Document) > 0 {
			_, err = db.UpdateMany(ret.Target, ret.Filter, bson.M{"$inc": ret.Document}, true)
			if err != nil {
				log.Error(err)
			}
		}
	}

	return ret.Hex(), nil
}

func (t *Table) Import(docs []any) (ids []string, err error) {
	//没有hook，则直接InsertMany
	//if t.Hook == nil || t.Hook.BeforeInsert == nil && t.Hook.AfterInsert == nil {
	//	if _, ok := t.scripts["before.insert"]; !ok {
	//		if _, ok := t.scripts["after.insert"]; !ok {
	//			oids, err := db.InsertMany(t.Name, docs)
	//			if err != nil {
	//				return nil, err
	//			}
	//			for _, id := range oids {
	//				ids = append(ids, id.Hex())
	//			}
	//			return ids, nil
	//		}
	//	}
	//}

	//依次插入
	for _, doc := range docs {
		id, err := t.Insert(doc)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (t *Table) ImportDocument(docs []db.Document) (ids []string, err error) {
	//没有hook，则直接InsertMany
	//if t.Hook == nil || t.Hook.BeforeInsert == nil && t.Hook.AfterInsert == nil {
	//	if _, ok := t.scripts["before.insert"]; !ok {
	//		if _, ok := t.scripts["after.insert"]; !ok {
	//			ds := make([]any, 0, len(docs))
	//			for _, doc := range docs {
	//				ds = append(ds, doc)
	//			}
	//			oids, err := db.InsertMany(t.Name, ds)
	//			if err != nil {
	//				return nil, err
	//			}
	//			for _, id := range oids {
	//				ids = append(ids, id.Hex())
	//			}
	//			return ids, nil
	//		}
	//	}
	//}

	//依次插入
	for _, doc := range docs {
		id, err := t.Insert(doc)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (t *Table) Delete(id string) error {
	//before delete
	if t.Hook != nil && t.Hook.BeforeDelete != nil {
		err := t.Hook.BeforeDelete(id)
		if err != nil {
			return exception.Wrap(err)
		}
	}
	if hook, ok := t.scripts["before.delete"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_, err := vm.RunProgram(hook)
		if err != nil {
			return exception.Wrap(err)
		}
	}

	oid, err := db.ParseObjectId(id)
	if err != nil {
		return err
	}

	var result db.Document
	has, err := db.FindOneAndDelete(t.Name, bson.D{{"_id", oid}}, &result)
	if err != nil {
		return err
	}
	if !has {
		return NotFound
	}

	//备份
	if t.DeleteOptions != nil && t.DeleteOptions.Backup {
		tab := t.DeleteOptions.BackupTable
		if tab == "" {
			tab = t.Name + ".history"
		}
		//把删除保存到修改历史表
		result["__id"] = oid
		result["deleted"] = time.Now()
		delete(result, "_id")
		_, _ = db.InsertOne(tab, result)
	}

	//转换_id
	db.StringifyDocumentObjectId(result)

	//after delete
	if t.Hook != nil && t.Hook.AfterDelete != nil {
		err := t.Hook.AfterDelete(id, result)
		if err != nil {
			return exception.Wrap(err)
		}
	}
	if hook, ok := t.scripts["after.delete"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_ = vm.Set("object", result)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return exception.Wrap(err)
		}
	}

	//累加器
	for _, a := range t.Accumulations {
		ret, err := a.Evaluate(result)
		if err != nil {
			log.Error(err)
			continue
		}
		if len(ret.Document) > 0 {
			_, err = db.UpdateMany(ret.Target, ret.Filter, bson.M{"$dec": ret.Document}, false)
			if err != nil {
				log.Error(err)
			}
		}
	}

	return err
}

func (t *Table) Update(id string, update any) error {

	//before update
	if t.Hook != nil && t.Hook.BeforeUpdate != nil {
		err := t.Hook.BeforeUpdate(id, update)
		if err != nil {
			return exception.Wrap(err)
		}
	}
	if hook, ok := t.scripts["before.update"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_ = vm.Set("change", update)
		_, err := vm.RunProgram(hook)
		if err != nil {
			return exception.Wrap(err)
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

	//转换_id
	oid, err := db.ParseObjectId(id)
	if err != nil {
		return err
	}
	db.ParseDocumentObjectId(update)

	var base db.Document
	has, err := db.FindOneAndUpdate(t.Name, bson.D{{"_id", oid}}, bson.D{{"$set", update}}, &base)
	if err != nil {
		return err
	}
	if !has {
		return NotFound
	}

	//把差异保存到修改历史表
	_, _ = db.InsertOne(t.Name+".change", bson.M{"object_id": oid, "base": base, "change": update})

	//备份
	if t.HistoryOptions != nil && t.HistoryOptions.Backup {
		tab := t.HistoryOptions.BackupTable
		if tab == "" {
			tab = t.Name + ".history"
		}
		//把删除保存到修改历史表
		base["__id"] = oid
		base["updated"] = time.Now()
		delete(base, "_id")
		_, _ = db.InsertOne(tab, base)
	}

	//转换—_id
	db.StringifyDocumentObjectId(update)
	db.StringifyDocumentObjectId(base)

	//after update
	if t.Hook != nil && t.Hook.AfterUpdate != nil {
		err := t.Hook.AfterUpdate(id, update, base)
		if err != nil {
			return exception.Wrap(err)
		}
	}
	if hook, ok := t.scripts["after.update"]; ok {
		vm := javascript.Runtime()
		_ = vm.Set("_id", id)
		_ = vm.Set("change", update)
		_ = vm.Set("base", base)
		_, err = vm.RunProgram(hook)
		if err != nil {
			return exception.Wrap(err)
		}
	}

	//累加器，先减，再加
	for _, a := range t.Accumulations {
		ret, err := a.Evaluate(base)
		if err != nil {
			log.Error(err)
			continue
		}
		if len(ret.Document) > 0 {
			_, err = db.UpdateMany(ret.Target, ret.Filter, bson.M{"$dec": ret.Document}, true)
			if err != nil {
				log.Error(err)
			}
		}
	}
	//补充字段，base已经被污染
	if u, ok := update.(map[string]any); ok {
		for k, v := range u {
			base[k] = v
		}
	}
	for _, a := range t.Accumulations {
		ret, err := a.Evaluate(update)
		if err != nil {
			log.Error(err)
			continue
		}
		if len(ret.Document) > 0 {
			_, err = db.UpdateMany(ret.Target, ret.Filter, bson.M{"$inc": ret.Document}, true)
			if err != nil {
				log.Error(err)
			}
		}
	}

	return err
}

func (t *Table) Get(id string, result any) (has bool, err error) {
	oid, err := db.ParseObjectId(id)
	if err != nil {
		return false, err
	}
	return db.FindOne(t.Name, bson.D{{"_id", oid}}, result)
}

func (t *Table) GetDocument(id string, result *db.Document) (has bool, err error) {
	oid, err := db.ParseObjectId(id)
	if err != nil {
		return false, err
	}
	db.StringifyDocumentObjectId(result)
	return db.FindOne(t.Name, bson.D{{"_id", oid}}, result)
}

func (t *Table) Find(filter any, sort any, skip int64, limit int64, results any) error {
	db.ParseDocumentObjectId(filter)
	return db.Find(t.Name, filter, sort, skip, limit, results)
}

func (t *Table) FindDocument(filter any, sort any, skip int64, limit int64, results *[]db.Document) error {
	db.ParseDocumentObjectId(filter)
	return db.Find(t.Name, filter, sort, skip, limit, results)
}

func (t *Table) Count(filter any) (count int64, err error) {
	db.ParseDocumentObjectId(filter)
	return db.Count(t.Name, filter)
}

func (t *Table) Drop() error {
	return db.Drop(t.Name)
}

func (t *Table) Distinct(filter any, field string) (values []any, err error) {
	db.ParseDocumentObjectId(filter)
	return db.Distinct(t.Name, filter, field)
}

func (t *Table) DistinctId(filter any) (values []string, err error) {
	db.ParseDocumentObjectId(filter)
	ids, err := db.DistinctId(t.Name, filter)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		values = append(values, id.Hex())
	}
	return values, nil
}

func (t *Table) Snapshot(into string) (err error) {
	var docs []db.Document

	//默认表名
	if into == "" {
		into = t.Name + ".snapshot"
	}

	now := time.Now()
	pipeline := mongo.Pipeline{
		bson.D{{"$set", bson.M{"object_id": "$_id", "date": now}}},
		bson.D{{"$unset", "_id"}},
		bson.D{{"$merge", bson.M{"into": into}}},
	}

	return t.Aggregate(pipeline, &docs)
}
