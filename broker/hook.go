package broker

import (
	"bytes"
	"github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
	"strings"
)

type Hook struct {
	mqtt.HookBase
}

func (h *Hook) ID() string {
	return "broker"
}
func (h *Hook) Provides(b byte) bool {
	//高效吗？
	return bytes.Contains([]byte{
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnUnsubscribed,
	}, []byte{b})
}

func (h *Hook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	//cl.Net.Listener todo websocket 直接鉴权通过

	return true
}

func (h *Hook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	//只允许发送属性事件

	return true
}

func (h *Hook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {

}

func (h *Hook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
	//device/+/values
	//project/+/values
	for _, f := range pk.Filters {
		ss := strings.Split(f.Filter, "/")
		if len(ss) == 3 {
			switch ss[0] {
			case "device":
				watchDeviceValues(ss[1])
			case "project":
				watchProjectValues(ss[1])
			case "space":
				watchSpaceValues(ss[1])
			}
		}
	}
}

func (h *Hook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	for _, f := range pk.Filters {
		ss := strings.Split(f.Filter, "/")
		if len(ss) == 3 {
			switch ss[0] {
			case "device":
				unWatchDeviceValues(ss[1])
			case "project":
				unWatchProjectValues(ss[1])
			case "space":
				unWatchSpaceValues(ss[1])
			}
		}
	}
}
