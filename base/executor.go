package base

type Executor interface {
	Execute(actions []*Action)
}

type DeviceContainer interface {
	Devices(productId string) (ids []string, err error)
}
