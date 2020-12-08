package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func setApiKey(body io.Reader, p *profile) (string, int) {
	var s ApiKeyRequest
	err := json.NewDecoder(body).Decode(&s)
	if err != nil {
		return "{ \"error\": \"" + err.Error() + "\"}", http.StatusBadRequest
	}
	p.store.SetApiKey(s.ApiKey)
	result, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("{ \"error\": \"" + err.Error() + "\"}"), http.StatusInternalServerError
	} else {
		return string(result), http.StatusOK
	}
	return string(result), http.StatusOK
}

func getApiKey(p *profile) (string, int) {
	apiKey, _ := p.store.GetApiKey()
	s := ApiKeyRequest{ApiKey: apiKey}
	result, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("{ \"error\": \"" + err.Error() + "\"}"), http.StatusInternalServerError
	} else {
		return string(result), http.StatusOK
	}
}
