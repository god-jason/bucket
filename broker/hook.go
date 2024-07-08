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
	//执行unsubscribe
	subs := cl.State.Subscriptions.GetAll()
	for _, sub := range subs {
		handleUnsubscribe(sub.Filter)
	}
}

func (h *Hook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
	//device/+/values
	//project/+/values
	//space/+/values
	for _, f := range pk.Filters {
		handleSubscribe(f.Filter)
	}
}

func (h *Hook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	for _, f := range pk.Filters {
		handleUnsubscribe(f.Filter)
	}
}

func handleSubscribe(filter string) {
	ss := strings.Split(filter, "/")
	if len(ss) == 3 && ss[2] == "values" {
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

func handleUnsubscribe(filter string) {
	ss := strings.Split(filter, "/")
	if len(ss) == 3 && ss[2] == "values" {
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
