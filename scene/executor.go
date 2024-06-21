package scene

import "github.com/god-jason/bucket/base"

type Executor interface {
	Execute(action []*base.Action) error
}
