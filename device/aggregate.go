package device

import (
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pool"
	"time"
)

func hourAggregate() {
	log.Println("整点聚合")
	now := time.Now()

	//先创建快照
	devices.Range(func(_ string, dev *Device) bool {
		dev.snap()
		return true
	})

	//再慢慢写入历史数据库
	devices.Range(func(_ string, dev *Device) bool {
		_ = pool.Insert(func() {
			dev.aggregate(now)
		})
		return true
	})

}
