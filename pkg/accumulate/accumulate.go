package accumulate

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/god-jason/bucket/pkg/calc"
	"strings"
)

type Field struct {
	Key   string `json:"key"`
	Value any    `json:"value"`

	_key   gval.Evaluable
	_value gval.Evaluable
}

func (f *Field) Compile() (err error) {
	if expr, has := strings.CutPrefix(f.Key, "="); has {
		f._key, err = calc.New(expr)
		if err != nil {
			return err
		}
	}
	if value, ok := f.Value.(string); ok {
		if expr, has := strings.CutPrefix(value, "="); has {
			f._value, err = calc.New(expr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Result struct {
	Target string         `json:"target"`
	Filter map[string]any `json:"filter"`
	Values map[string]any `json:"values"`
}

// Accumulation 累积器，空间换时间，主要用于统计
type Accumulation struct {
	Target string   `json:"target"`
	Filter []*Field `json:"filter"`
	Values []*Field `json:"values"`

	_target gval.Evaluable
}

func (a *Accumulation) Init() (err error) {
	if expr, has := strings.CutPrefix(a.Target, "="); has {
		a._target, err = calc.New(expr)
		if err != nil {
			return err
		}
	}

	for _, f := range a.Filter {
		err = f.Compile()
		if err != nil {
			return err
		}
	}

	for _, f := range a.Values {
		err = f.Compile()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Accumulation) Evaluate(args any) (result *Result, err error) {
	var ret Result

	//目标库
	if a._target != nil {
		ret.Target, err = a._target.EvalString(context.Background(), args)
		if err != nil {
			return
		}
	} else {
		ret.Target = a.Target
	}

	//过滤器
	ret.Filter = make(map[string]any)
	for _, f := range a.Filter {

		key := f.Key
		if f._key != nil {
			key, err = f._key.EvalString(context.Background(), args)
			if err != nil {
				return
			}
		}

		value := f.Value
		if f._value != nil {
			value, err = f._key(context.Background(), args)
			if err != nil {
				return
			}
		}
		ret.Filter[key] = value
	}

	//值
	ret.Values = make(map[string]any)
	for _, f := range a.Values {

		key := f.Key
		if f._key != nil {
			key, err = f._key.EvalString(context.Background(), args)
			if err != nil {
				return
			}
		}

		value := f.Value
		if f._value != nil {
			value, err = f._key(context.Background(), args)
			if err != nil {
				return
			}
		}
		ret.Filter[key] = value
	}

	return &ret, nil
}
