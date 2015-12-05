package main

import (
	log "github.com/Sirupsen/logrus"

	"os"
)

func initializeLogger() {
	if os.Getenv("LOG_FORMAT") == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}
