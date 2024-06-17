package device

import (
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "device/create", nil)
	api.Register("POST", "device/update/:id", nil)
	api.Register("GET", "device/delete/:id", nil)

	api.Register("GET", "device/history/:id", nil)
	api.Register("GET", "device/watch/:id", nil)

}
