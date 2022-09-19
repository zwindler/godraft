package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/zwindler/godraft/services"
)

func Start() error {
	configRepository, err := services.GetConfig()
	if err != nil {
		return err
	}

	addr := configRepository.BuildServerAddr()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Start listening on", addr)
	return http.ListenAndServe(addr, nil)
}
