package user

import (
	"github.com/god-jason/bucket/config"
)

const MODULE = "user"

func init() {
	config.Register(MODULE, "default_password", "123456")
	config.Register(MODULE, "password_hash", "md5")
	config.Register(MODULE, "password_suffix", "")

}
