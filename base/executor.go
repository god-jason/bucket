package base

type Executor interface {
	Execute(actions []*Action)
}
