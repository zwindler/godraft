package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

const (
	serverAddrEnvVar  = "SERVER_ADDR"
	serverPortEnvVar  = "SERVER_PORT"
	configPathEnvVar  = "GODRAFT_CONFIG"
)

func parse(path string) (Config, error) {
	var config Config

	// check if file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Warnf("Unable to parse, configuration file not found: %w. Skipping", err)
		return (Config{}), nil
	} else {
		// try to parse file
		if _, err := toml.DecodeFile(path, &config); err != nil {
			return (Config{}), err
		}
	}
	log.Printf("Configuration extracted from %s file", path)
	return config, nil
}

func LoadConfig() (config Config, err error) {
	configPathFromEnv := os.Getenv(configPathEnvVar)
	if configPathFromEnv != "" {
		config, err = parse(configPathFromEnv)
	} else {
		config, err = parse("config.toml")
	}
	if err != nil {
		return (Config{}), err
	}
	return config, err
}

func (c *Config) SetDefaults() {

	if c.ServerAddr == "" {
		c.ServerAddr = "0.0.0.0"
	}

	if c.ServerPort == 0 {
		c.ServerPort = 3000
	}
}

func (c *Config) SetConfigFromEnv() (err error) {
	serverAddrFromEnv := os.Getenv(serverAddrEnvVar)
	if serverAddrFromEnv != "" {
		c.ServerAddr = serverAddrFromEnv
	}

	portFromEnv := os.Getenv(serverPortEnvVar)
	if portFromEnv != "" {
		serverPort, err := strconv.Atoi(portFromEnv)
		if err != nil {
			err = fmt.Errorf("%w %s", ErrInvalidVar, serverPortEnvVar)
			return err
		}
		c.ServerPort = serverPort
	}

	return nil
}
