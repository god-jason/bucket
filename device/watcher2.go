package device

type Watcher interface {
	OnDeviceAdd(device *Device) //监听属性变化
	OnDeviceRemove(device *Device)
}

type ValuesWatcher interface {
	OnProjectValuesChange(device *Device, values map[string]any) //监听属性变化
}
