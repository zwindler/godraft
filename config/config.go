package config

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidVar = errors.New("invalid var")
)

type Config struct {
	ServerAddr  string  `toml:"server_addr"`
	ServerPort  int     `toml:"server_port"`
}

func NewConfig(config Config) *Config {
	return &config
}

func (c *Config) BuildServerAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}
