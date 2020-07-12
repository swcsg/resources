package main

import (
	"github.com/maxwell92/log"
	"resource/pkg/router"
)

var logger = log.Log

func main() {
	logger.Infoln("start...")

	r := router.RegisterRoutes()

	_ = r.Run("localhost:8090")
}
