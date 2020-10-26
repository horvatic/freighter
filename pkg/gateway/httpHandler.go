package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"net/http"
)

var store *datastore.MemoryDataStore

func handleConfig(w http.ResponseWriter, req *http.Request) {
	var s datastore.Service
	err := json.NewDecoder(req.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.SetService(&s)
	fmt.Fprintf(w, "Service: %+v", s)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	response, statusCode := route(&Request{UriPath: req.URL.Path}, &proxy.RequestProxy{}, store)
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, response)
}

func Start() {
	store = datastore.NewMemoryDataStore()
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/config", handleConfig)
	http.ListenAndServe(":8080", nil)
}
