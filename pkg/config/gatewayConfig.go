package config

import (
	"os"
)

type GatewayConfig struct {
	AuthKey string
}

func BuildGatewayConfig() *GatewayConfig {
		return &GatewayConfig{
		AuthKey: os.Getenv("API_KEY"),
	}
}
