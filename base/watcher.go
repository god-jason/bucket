package base

type DeviceValuesWatcher interface {
	OnDeviceValuesChange(product, device string, values map[string]any) //监听属性变化
}

type ProjectValuesWatcher interface {
	OnProjectValuesChange(project, product, device string, values map[string]any) //监听属性变化
}

type SpaceValuesWatcher interface {
	OnSpaceValuesChange(space, product, device string, values map[string]any) //监听属性变化
}

//type ProductValuesWatcher interface {
//	OnProductValuesChange(product, device string, values map[string]any) //监听产品属性变化
//}

//type Watcher interface {
//	OnDeviceAdd(device *device.Device) //监听属性变化
//	OnDeviceRemove(device *device.Device)
//}

//type DeviceValuesWatcher interface {
//	OnProjectValuesChange(device string, values map[string]any) //监听属性变化
//}
