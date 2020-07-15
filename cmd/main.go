package main

import (
	"resource/pkg/router"

	"github.com/maxwell92/log"
)

var logger = log.Log

func main() {
	logger.Infoln("start...")

	r := router.RegisterRoutes()

	_ = r.Run("localhost:8080")
}
