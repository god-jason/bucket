package timer

import "github.com/god-jason/bucket/base"

type Executor interface {
	Execute(actions []*base.Action)
}
