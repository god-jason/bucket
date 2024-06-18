package condition

import "errors"

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
