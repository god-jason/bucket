package gateway

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type IncomingHook struct {
	mqtt.HookBase
}

func (h *IncomingHook) ID() string {
	return "incoming"
}

func (h *IncomingHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	//todo 检查用户名密码
	return true
}

func (h *IncomingHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	//todo 只允许发送属性事件
	return true
}

func (h *IncomingHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	return nil
}

func (h *IncomingHook) OnSessionEstablish(cl *mqtt.Client, pk packets.Packet) {
}

func (h *IncomingHook) OnSessionEstablished(cl *mqtt.Client, pk packets.Packet) {
}

func (h *IncomingHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
}

func (h *IncomingHook) OnAuthPacket(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	return pk, nil
}

func (h *IncomingHook) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	//todo 只允许订阅自己的消息
	return pk
}

func (h *IncomingHook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
}

func (h *IncomingHook) OnUnsubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	return pk
}

func (h *IncomingHook) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
}

func (h *IncomingHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	return pk, nil
}

func (h *IncomingHook) OnClientExpired(cl *mqtt.Client) {
}
