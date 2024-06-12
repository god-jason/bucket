package main

import (
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/web"
)

func main() {
	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := db.Ping()
		if err != nil {
			log.Error(err)
		}
	}()
	
	_ = web.Serve()
}
