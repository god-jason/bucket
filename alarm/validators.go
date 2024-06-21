package alarm

import (
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validators lib.Map[Validator]

func Get(id string) *Validator {
	return validators.Load(id)
}

func From(v *Validator) (err error) {
	validators.Store(v.Id.Hex(), v)
	return v.Init()
}

func Load(id primitive.ObjectID) error {
	var validator Validator
	err := _validatorTable.Get(id, &validator)
	if err != nil {
		return err
	}
	return From(&validator)
}

func Unload(id primitive.ObjectID) error {
	t := validators.LoadAndDelete(id.Hex())
	if t != nil {
		//return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Validator](&_validatorTable, nil, 100, func(t *Validator) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
		}
		return nil
	})
}
