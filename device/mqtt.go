package device

import (
	"encoding/json"
	"github.com/god-jason/bucket/log"
	"github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type PayloadHistory struct {
	Values    map[string]any `json:"values"`
	Timestamp int64          `json:"timestamp"`
}

func HandleMqtt(id string, typ string, cl *mqtt.Client, payload []byte) {
	dev := devices.Load(id)
	if dev == nil {
		return
	}

	switch typ {
	case "values":
		var values map[string]any
		err := json.Unmarshal(payload, &values)
		if err != nil {
			log.Error(err)
			return
		}
		dev.PatchValues(values)
	case "history":
		var histories []PayloadHistory
		err := json.Unmarshal(payload, &histories)
		if err != nil {
			log.Error(err)
			return
		}
		for _, history := range histories {
			dev.WriteHistory(history.Values, history.Timestamp)
		}
	case "action":
		//action调用web接口

	case "event":
		//解析事件，生成 alarm

	}
}

// 直接向mqtt客户端发布消息
func publishDirectly(cl *mqtt.Client, topic string, payload any) error {
	buf, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	pkg := packets.Packet{
		FixedHeader: packets.FixedHeader{
			Type: packets.Publish,
		},
		TopicName: topic,
		Payload:   buf,
	}
	return cl.WritePacket(pkg)
}
