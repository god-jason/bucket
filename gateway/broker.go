package gateway

import (
	mochi "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

var server *mochi.Server

func startup() error {
	opts := &mochi.Options{
		InlineClient: true,
	}
	server = mochi.New(opts)
	var cfs []listeners.Config
	_ = server.AddListenersFromConfig(cfs)
	return server.Serve()
}
