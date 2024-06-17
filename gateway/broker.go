package gateway

import (
	"github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

var server *mqtt.Server

func Startup() error {
	opts := &mqtt.Options{
		InlineClient: true,
	}
	server = mqtt.New(opts)
	var cfs []listeners.Config
	_ = server.AddListenersFromConfig(cfs)
	return server.Serve()
}

func Shutdown() error {
	if server != nil {
		return server.Close()
	}
	return nil
}
