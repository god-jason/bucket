package condition

import (
	"context"
	"errors"
	"github.com/PaesslerAG/gval"
	"github.com/god-jason/bucket/pkg/calc"
)

type Compare struct {
	Variable string `json:"variable,omitempty"` //变量
	Operator string `json:"operator,omitempty"` //对比算子 > >= < <= !=
	Value    string `json:"value,omitempty"`    //值，支持表达式
	_value   gval.Evaluable
}

func (c *Compare) Init() (err error) {
	c._value, err = calc.New(c.Variable + c.Operator + "(" + c.Value + ")")
	return
}

func (c *Compare) Eval(ctx map[string]any) (bool, error) {
	return c._value.EvalBool(context.Background(), ctx)
}

type And struct {
	Compares []*Compare `json:"compares,omitempty"`
}

func (a *And) Init() error {
	for _, c := range a.Compares {
		err := c.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *And) Eval(ctx map[string]any) (bool, error) {
	if len(a.Compares) == 0 {
		return false, errors.New("没有对比")
	}
	for _, c := range a.Compares {
		ret, err := c.Eval(ctx)
		if err != nil {
			return ret, err
		}
		if !ret {
			return false, nil
		}
	}
	return true, nil
}

type Or struct {
	Compares []*Compare `json:"compares,omitempty"`
}

func (a *Or) Init() error {
	for _, c := range a.Compares {
		err := c.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Or) Eval(ctx map[string]any) (bool, error) {
	if len(a.Compares) == 0 {
		return false, errors.New("没有对比")
	}
	for _, c := range a.Compares {
		ret, err := c.Eval(ctx)
		if err != nil {
			return ret, err
		}
		if ret {
			return true, nil
		}
	}
	return true, nil
}
