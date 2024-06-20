package device

import (
	"github.com/god-jason/bucket/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	api.Register("POST", "device/create", api.Create(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("POST", "device/update/:id", api.Update(&_table, func(id primitive.ObjectID) error {
		dev := Get(id.Hex())
		if dev != nil {
			_ = dev.Close() //报错
		}
		return Load(id.Hex())
	}))

	api.Register("GET", "device/delete/:id", api.Delete(&_table, func(id primitive.ObjectID) error {
		dev := Get(id.Hex())
		if dev != nil {
			return dev.Close() //报错
		}
		//todo 删除报警，场景，等
		return nil
	}))

	api.Register("GET", "device/detail/:id", api.Detail(&_table, nil))

	api.Register("GET", "device/enable/:id", api.Enable(&_table, func(id primitive.ObjectID) error {
		return Load(id.Hex())
	}))

	api.Register("GET", "device/disable/:id", api.Disable(&_table, func(id primitive.ObjectID) error {
		dev := Get(id.Hex())
		if dev != nil {
			return dev.Close() //报错
		}
		return nil
	}))

	api.Register("POST", "device/count", api.Count(&_table))

	api.Register("POST", "device/search", api.Search(&_table, nil))

	api.Register("POST", "device/group", api.Group(&_table, nil))

	api.Register("POST", "device/import", api.Import(&_table, func(id []primitive.ObjectID) error {
		//todo 依次加载
		return nil
	}))

	api.Register("POST", "device/export", api.Export(&_table))

}
