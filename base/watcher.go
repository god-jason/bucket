package base

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ValuesWatcher interface {
	OnValuesChange(product, device primitive.ObjectID, values map[string]any) //监听属性变化
}

//type ProductValuesWatcher interface {
//	OnProductValuesChange(product, device primitive.ObjectID, values map[string]any) //监听产品属性变化
//}

//type Watcher interface {
//	OnDeviceAdd(device *device.Device) //监听属性变化
//	OnDeviceRemove(device *device.Device)
//}

//type ValuesWatcher interface {
//	OnProjectValuesChange(device primitive.ObjectID, values map[string]any) //监听属性变化
//}
