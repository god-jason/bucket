package project

import "github.com/god-jason/bucket/device"

type Watcher interface {
	OnDeviceAdd(device *device.Device) //监听属性变化
	OnDeviceRemove(device *device.Device)
}
