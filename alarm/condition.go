package alarm

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/god-jason/bucket/pkg/calc"
	"github.com/god-jason/bucket/pkg/exception"
)

type Condition struct {
	//外or，内and
	Conditions [][]*Compare `json:"conditions,omitempty"`
}

func (a *Condition) Init() error {
	for _, c := range a.Conditions {
		for _, v := range c {
			err := v.Init()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Condition) Eval(ctx map[string]any) (bool, error) {
	if len(a.Conditions) == 0 {
		return false, exception.New("没有对比")
	}

	//外部用or
	for _, c := range a.Conditions {
		if len(c) == 0 {
			continue
		}

		//内部用and
		and := true
		for _, v := range c {
			ret, err := v.Eval(ctx)
			if err != nil {
				return ret, err
			}
			//只要一个false，就退出
			if !ret {
				and = false
				break
			}
		}

		//只要有一个true，就返回
		if and {
			return and, nil
		}
	}

	return false, nil
}

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
