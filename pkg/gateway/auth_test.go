package gateway

import (
	"github.com/horvatic/freighter/pkg/config"
	"github.com/horvatic/freighter/test/mock"
	"testing"
)

func TestAuthSuccessful(t *testing.T) {
	database := &mock.MockDataStore{ApiKey: "letmein123"}
	config := &config.GatewayConfig{}

	if !hasValidApiKey(config, "letmein123", database) {
		t.Errorf("got %t want %t", false, true)
	}
}

func TestAuthFails(t *testing.T) {
	database := &mock.MockDataStore{ApiKey: "123"}
	config := &config.GatewayConfig{}

	if hasValidApiKey(config, "fefwefwefwe", database) {
		t.Errorf("got %t want %t", true, false)
	}
}

func TestAuthRemoveApiKeyHeader(t *testing.T) {
	headers := make(map[string][]string)
	headers["head"] = []string{"headerinfo"}
	headers["Apikey"] = []string{"fwefwefwef"}

	headers = removeApiKeyHeader(headers)

	if len(headers) != 1 || headers["head"][0] != "headerinfo" {
		want := make(map[string][]string)
		want["head"] = []string{"headerinfo"}
		t.Errorf("got %q want %q", headers, want)
	}
}

func TestAuthGetApiKeyHeaderProperCase(t *testing.T) {
	headers := make(map[string][]string)
	headers["head"] = []string{"headerinfo"}
	headers["Apikey"] = []string{"fwefwefwef"}

	key := getApiKeyFromHeader(headers)

	if key != "fwefwefwef" {
		t.Errorf("got %q want %q", key, "fwefwefwef")
	}
}

func TestAuthGetApiKeyHeaderEmpty(t *testing.T) {
	headers := make(map[string][]string)
	headers["head"] = []string{"headerinfo"}

	key := getApiKeyFromHeader(headers)

	if key != "" {
		t.Errorf("got %q want %q", key, "")
	}
}
