package db

import (
	"github.com/god-jason/bucket/config"
)

const MODULE = "database"

func init() {
	config.Register(MODULE, "url", "mongodb://localhost:27017")
	config.Register(MODULE, "database", "bucket")
}
