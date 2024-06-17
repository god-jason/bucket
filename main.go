package main

import (
	_ "github.com/god-jason/bucket/aggregate"
	"github.com/god-jason/bucket/boot"
	_ "github.com/god-jason/bucket/device"
	_ "github.com/god-jason/bucket/history"
	"github.com/god-jason/bucket/log"
	_ "github.com/god-jason/bucket/product"
	_ "github.com/god-jason/bucket/table"
	"github.com/god-jason/bucket/web"
)

func main() {
	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
	}

	_ = web.Serve()
}
