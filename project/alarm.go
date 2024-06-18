package project

import "github.com/god-jason/bucket/pkg/condition"

//外or，内and

// todo 支持项目下单个设备的报警，
type Alarm struct {
	Condition condition.Condition
}
