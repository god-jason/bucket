package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

type Batch struct {
	Collection   string
	WriteTimeout time.Duration
	BufferSize   int

	models []mongo.WriteModel
	locker sync.Locker
	timer  *time.Timer
}

func (b *Batch) InsertOne(doc interface{}) {
	model := mongo.NewInsertOneModel().SetDocument(doc)
	b.Write(model)
}

func (b *Batch) UpdateOne(filter interface{}, update interface{}, upsert bool) {
	model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(upsert)
	b.Write(model)
}

func (b *Batch) UpdateMany(filter interface{}, update interface{}, upsert bool) {
	model := mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(update).SetUpsert(upsert)
	b.Write(model)
}

func (b *Batch) DeleteOne(filter interface{}) {
	model := mongo.NewDeleteOneModel().SetFilter(filter)
	b.Write(model)
}

func (b *Batch) DeleteMany(filter interface{}) {
	model := mongo.NewDeleteManyModel().SetFilter(filter)
	b.Write(model)
}

func (b *Batch) Write(model mongo.WriteModel) {
	defer b.locker.Unlock()
	b.locker.Lock()
	b.models = append(b.models, model)

	//操作定时器
	if b.timer == nil {
		//启动定时器
		b.timer = time.AfterFunc(b.WriteTimeout, func() {
			b.timer = nil
			_, _ = b.Flush()
			//TODO log
		})
	} else {
		//满了就立即执行
		if b.BufferSize > 0 && len(b.models) >= b.BufferSize {
			b.timer.Reset(time.Millisecond)
		}
	}
}

func (b *Batch) Flush() (*mongo.BulkWriteResult, error) {
	if len(b.models) == 0 {
		return nil, nil
	}

	//取出models并置空
	b.locker.Lock()
	ms := b.models
	b.models = nil //此处需不需要锁？？?
	b.locker.Unlock()

	//错误就不管了
	return db.Collection(b.Collection).BulkWrite(context.Background(), ms)
}
