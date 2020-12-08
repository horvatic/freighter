package profile

import (
	"encoding/json"
	"fmt"
	"github.com/horvatic/freighter/pkg/datastore"
	"io"
	"net/http"
)

func register(body io.Reader, p *profile) (string, int) {
	var s datastore.Service
	err := json.NewDecoder(body).Decode(&s)
	if err != nil {
		return fmt.Sprintf("{ \"error\": \"" + err.Error() + "\"}"), http.StatusBadRequest
	}
	p.store.SetService(&s)
	result, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("{ \"error\": \"" + err.Error() + "\"}"), http.StatusInternalServerError
	} else {
		return string(result), http.StatusOK
	}
	return string(result), http.StatusOK
}
