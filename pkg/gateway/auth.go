package gateway

import (
	"github.com/horvatic/freighter/pkg/config"
	"net/http"
)

func hasValidApiKey(config *config.GatewayConfig, key string) bool {
	return config.AuthKey == key
}

func removeApiKeyHeader(headers http.Header) http.Header {
	if len(headers["Apikey"]) > 0 {
		delete(headers, "Apikey")
	}
	return headers
}

func getApiKeyFromHeader(headers http.Header) string {
	if len(headers["Apikey"]) > 0 {
		return headers["Apikey"][0]
	}
	return ""
}
