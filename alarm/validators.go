package alarm

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/god-jason/bucket/table"
)

var validators lib.Map[Validator]

func Get(id string) *Validator {
	return validators.Load(id)
}

func From(v *Validator) (err error) {
	tt := validators.LoadAndStore(v.Id, v)
	if tt != nil {
		_ = tt.Close()
	}
	return v.Open()
}

func Load(id string) error {
	var validator Validator
	has, err := _validatorTable.Get(id, &validator)
	if err != nil {
		return err
	}
	if !has {
		return exception.New("找不到记录")
	}
	return From(&validator)
}

func Unload(id string) error {
	t := validators.LoadAndDelete(id)
	if t != nil {
		return t.Close()
	}
	return nil
}

func LoadAll() error {
	return table.BatchLoad[*Validator](&_validatorTable, base.FilterEnabled, 100, func(t *Validator) error {
		//并行加载
		err := From(t)
		if err != nil {
			log.Error(err)
			//return err
		}
		return nil
	})
}
