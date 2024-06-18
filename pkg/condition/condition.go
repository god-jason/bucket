package condition

import "errors"

type Condition struct {
	//外or，内and
	Children []*And `json:"children,omitempty"`
}

func (a *Condition) Init() error {
	for _, c := range a.Children {
		err := c.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Condition) Eval(ctx map[string]any) (bool, error) {
	if len(a.Children) == 0 {
		return false, errors.New("没有对比")
	}
	for _, c := range a.Children {
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
