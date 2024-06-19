package device

import (
	"errors"
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/table"
)

var devices lib.Map[Device]

func Get(id string) *Device {
	return devices.Load(id)
}

func Load(id string) error {
	oid, err := db.ParseObjectId(id)
	if err != nil {
		return err
	}

	var doc table.Document
	err = db.FindById(base.BucketDevice, oid, &doc)
	if err != nil {
		return err
	}

	return From(doc)
}

func From(doc table.Document) (err error) {
	dev := new(Device)

	if id, ok := doc["_id"]; !ok {
		if dev.id, err = db.ParseObjectId(id); err != nil {
			return errors.New("_id 类型不正确")
		}
	} else {
		return errors.New("缺少 _id")
	}

	if id, ok := doc["product_id"]; !ok {
		if dev.productId, err = db.ParseObjectId(id); err != nil {
			return errors.New("product_id 类型不正确")
		}
	} else {
		return errors.New("缺少 product_id")
	}

	devices.Store(dev.ID(), dev)

	return dev.Open()
}
