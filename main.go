package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/zwindler/godraft/server"
	"github.com/zwindler/godraft/services"
)

var Version string

func init() {
	// Logrus config
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	services.Version = Version
	log.Printf("godraft version v%s", services.Version)
	services.DefaultTemplateStruct.SetDefaults()
	services.Register()
	server.Start()
}
