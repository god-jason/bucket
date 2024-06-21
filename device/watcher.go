package device

type Watcher interface {
	OnDeviceValuesChange(map[string]any) //监听属性变化
}
