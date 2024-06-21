package scene

import (
	"context"
	"github.com/PaesslerAG/gval"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/pkg/calc"
	"github.com/god-jason/bucket/pkg/errors"
)

//todo 条件 外or，内and，每个条件都要选具体设备

//todo 注册到具体的设备上，监听变化

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

func (a *Condition) Eval() (bool, error) {
	if len(a.Conditions) == 0 {
		return false, errors.New("没有对比")
	}

	//外部用or
	for _, c := range a.Conditions {
		if len(c) == 0 {
			continue
		}

		//内部用and
		and := true
		for _, v := range c {
			ret, err := v.Eval()
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
	DeviceId string `json:"device_id,omitempty" bson:"device_id,omitempty"`
	Variable string `json:"variable,omitempty"` //变量
	Operator string `json:"operator,omitempty"` //对比算子 > >= < <= !=
	Value    string `json:"value,omitempty"`    //值，支持表达式

	watchers device.Watcher
	expr     gval.Evaluable
}

func (c *Compare) Init() (err error) {
	c.expr, err = calc.New(c.Variable + c.Operator + "(" + c.Value + ")")
	return errors.Wrap(err)
}

func (c *Compare) Eval() (bool, error) {
	dev := device.Get(c.DeviceId)
	if dev == nil {
		return false, errors.New("设备未上线")
	}

	ret, err := c.expr.EvalBool(context.Background(), dev.Values())
	return ret, errors.Wrap(err)
}
