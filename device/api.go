package device

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
)

func init() {
	api.Register("POST", "device/create", nil)
	api.Register("POST", "device/update/:id", nil)
	api.Register("GET", "device/delete/:id", nil)
	api.Register("GET", "device/start/:id", nil)
	api.Register("GET", "device/stop/:id", nil)
	api.Register("GET", "device/restart/:id", nil)
	api.Register("GET", "device/enable/:id", nil)
	api.Register("GET", "device/disable/:id", nil)
	api.Register("GET", "device/values/:id", deviceValues)
	api.Register("POST", "device/values/:id", nil)
	api.Register("GET", "device/history/:id", nil)
	api.Register("GET", "device/watch/:id", nil)

}

func deviceValues(ctx *gin.Context) {

}
