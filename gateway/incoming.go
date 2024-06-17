package gateway

import (
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/pool"
	"github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"strings"
)

type IncomingHook struct {
	mqtt.HookBase
}

func (h *IncomingHook) ID() string {
	return "incoming"
}

func (h *IncomingHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	id, err := db.ParseObjectId(pk.Connect.ClientIdentifier)
	if err != nil {
		return false
	}

	//检查用户名密码
	var gw Gateway
	err = db.FindById(Bucket, id, &gw)
	if err != nil {
		return false
	}

	//检查用户名密码
	if gw.Username != "" {
		if gw.Username != string(pk.Connect.Username) || gw.Password != string(pk.Connect.Password) {
			return false
		}
	}

	return true
}

func (h *IncomingHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	//todo 只允许发送属性事件
	//cl.WritePacket(nil)

	return true
}

func (h *IncomingHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	//todo 网关离线，相关设备置为离线状态

}

func (h *IncomingHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	//直接处理数据
	topics := strings.Split(pk.TopicName, "/")
	if len(topics) != 3 {
		return pk, nil
	}

	//up/device/+/values 数据上传
	//up/device/+/action 接口响应
	//up/device/+/event 事件上报
	if topics[0] == "up" {
		//池化处理，避免拥堵
		_ = pool.Insert(func() {
			//解析数据，仅支持json格式（虽然效率低了点，但是没办法，大家都在用）
			//var payload map[string]interface{}
			//if len(pk.Payload) > 0 {
			//	err := json.Unmarshal(pk.Payload, &payload)
			//	if err != nil {
			//		return
			//	}
			//}

			//执行消息
			switch topics[1] {
			case "device":
				dev := device.Get(topics[2])
				if dev != nil {
					dev.HandleMqtt(topics[3], cl, pk.Payload)
				}
			case "tunnel":

			}
		})
	}

	return pk, nil
}
