package aggregator

import (
	"errors"
	"fmt"
)

var ErrorBlank = errors.New("无数据")

type Aggregator interface {
	Push(any) error
	Pop() any
	Snap() //快照
}

// New 新建
func New(typ string) (agg Aggregator, err error) {
	switch typ {
	case "inc", "increase":
		agg = &inc{}
	case "dec", "decrease":
		agg = &dec{}
	case "sum", "acc":
		agg = &sum{}
	case "avg", "average":
		agg = &avg{}
	case "cnt", "count":
		agg = &count{}
	case "min", "minimum":
		agg = &minimum{}
	case "max", "maximum":
		agg = &maximum{}
	case "first":
		agg = &first{}
	case "last":
		agg = &last{}
	default:
		err = fmt.Errorf("Unknown aggregate type %s ", typ)
	}
	return
}
