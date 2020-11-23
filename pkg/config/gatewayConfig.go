package config

import (
	"os"
)

type GatewayConfig struct {
	AuthKey string
	Port    string
}

func BuildGatewayConfig() *GatewayConfig {
	return &GatewayConfig{
		AuthKey: os.Getenv("API_KEY"),
		Port:    os.Getenv("LISTEN_PORT"),
	}
}
