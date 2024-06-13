package main

import (
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/log"
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
