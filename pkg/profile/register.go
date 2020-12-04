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
		return err.Error(), http.StatusBadRequest
	}
	p.store.SetService(&s)
	result := fmt.Sprintf("Service: %+v", s)
	return result, http.StatusOK
}
