package pool

import (
	"github.com/god-jason/bucket/config"
)

const MODULE = "pool"

func init() {
	config.Register(MODULE, "size", 10000)
}
