package service

import (
	"github.com/god-jason/bucket/config"
	"github.com/god-jason/bucket/lib"
)

const MODULE = "service"

func init() {
	config.Register(MODULE, "name", lib.AppName())
	config.Register(MODULE, "display", "物联大师")
	config.Register(MODULE, "description", "物联网数据中台")
	config.Register(MODULE, "arguments", []string{})
}
