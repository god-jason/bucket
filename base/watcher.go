package base

type DeviceValuesWatcher interface {
	OnDeviceValuesChange(map[string]any) //监听属性变化
}
