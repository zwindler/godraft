package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/zwindler/godraft/server"
	"github.com/zwindler/godraft/services"
)

func init() {
	// Logrus config
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	services.Register()
	server.Start()
}