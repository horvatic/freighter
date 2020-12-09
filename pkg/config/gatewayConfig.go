package config

import (
	"os"
)

type GatewayConfig struct {
	InitAuthKey string
	Port        string
}

func BuildGatewayConfig() *GatewayConfig {
	return &GatewayConfig{
		InitAuthKey: os.Getenv("INIT_API_KEY"),
		Port:        os.Getenv("LISTEN_PORT"),
	}
}
