package mongodb

import (
	"github.com/god-jason/bucket/config"
)

const MODULE = "mongodb"

func init() {
	config.Register(MODULE, "url", "mongodb://localhost:27017")
	config.Register(MODULE, "database", "bucket")
	config.Register(MODULE, "auth", "admin")
	config.Register(MODULE, "username", "admin")
	config.Register(MODULE, "password", "123456")
}
