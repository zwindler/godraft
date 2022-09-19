package services

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/zwindler/godraft/config"
)

var currentConfig config.Config

func GetConfig() (config.Config, error) {
	var err error
	// if there is no config yet, create an empty one
	if currentConfig == (config.Config{}) {
		// Load configuration file
		currentConfig, err = config.LoadConfig()
		if err != nil {
			err = fmt.Errorf("a variable from config file contains invalid data: %s", err)
			log.Fatalf("Config error: %s", err)
		}
		// Override values when Environment variables that are not empty
		err = currentConfig.SetConfigFromEnv()
		if err != nil {
			if errors.Is(err, config.ErrInvalidVar) {
				return config.Config{}, fmt.Errorf("a variable from env vars contains invalid data: %s", err)
			}
			log.Fatalf("Config error: %s", err)
		}
		currentConfig.SetDefaults()
	}

	return currentConfig, nil
}
