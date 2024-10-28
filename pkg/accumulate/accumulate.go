package accumulate

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/god-jason/bucket/pkg/calc"
	"strings"
)

type Field struct {
	Key   string
	Value any

	_key   gval.Evaluable
	_value gval.Evaluable
}

type Result struct {
	Target       string
	MetaFields   map[string]any
	TargetFields map[string]any
}

// Accumulation 累积器，空间换时间，主要用于统计
type Accumulation struct {
	Target       string         `json:"target"`
	MetaFields   map[string]any `json:"meta_fields"`
	TargetFields map[string]any `json:"target_fields"`

	_target gval.Evaluable
	_filter map[string]gval.Evaluable
	_fields []*Field
}

func (a *Accumulation) Init() (err error) {
	if expr, has := strings.CutPrefix(a.Target, "="); has {
		a._target, err = calc.New(expr)
		if err != nil {
			return err
		}
	}

	a._filter = make(map[string]gval.Evaluable)
	for key, value := range a.MetaFields {
		if val, ok := value.(string); ok {
			if expr, has := strings.CutPrefix(val, "="); has {
				a._filter[key], err = calc.New(expr)
				if err != nil {
					return err
				}
			}
		}
	}

	for key, value := range a.TargetFields {

		f := &Field{Key: key, Value: value}

		//键
		if expr, has := strings.CutPrefix(key, "="); has {
			f._key, err = calc.New(expr)
			if err != nil {
				return err
			}
		}

		//值
		if val, ok := value.(string); ok {
			if expr, has := strings.CutPrefix(val, "="); has {
				f._value, err = calc.New(expr)
				if err != nil {
					return err
				}
			}
		}

		a._fields = append(a._fields, f)
	}

	return nil
}

func (a *Accumulation) Evaluate(args any) (result *Result, err error) {
	var ret Result

	//目标
	if a._target != nil {
		ret.Target, err = a._target.EvalString(context.Background(), args)
		if err != nil {
			return
		}
	} else {
		ret.Target = a.Target
	}

	//过滤器
	ret.MetaFields = make(map[string]any)
	for key, value := range a._filter {
		if value != nil {
			ret.MetaFields[key], err = a._target(context.Background(), args)
			if err != nil {
				return
			}
		} else {
			ret.MetaFields[key] = a.MetaFields[key]
		}
	}

	ret.TargetFields = make(map[string]any)

	for _, f := range a._fields {
		key := f.Key
		if f._key != nil {
			key, err = f._key.EvalString(context.Background(), args)
			if err != nil {
				return
			}
		}

		val := f.Value
		if f._value != nil {
			val, err = f._value.EvalFloat64(context.Background(), args)
			if err != nil {
				return
			}

			//路过0值
			if val.(float64) == 0.0 {
				continue
			}
		}

		ret.TargetFields[key] = val
	}

	return &ret, nil
}
