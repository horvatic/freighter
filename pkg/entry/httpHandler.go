package entry

import (
	"encoding/json"
	"fmt"
	"github.com/horvatic/freighter/pkg/config"
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/gateway"
	"github.com/horvatic/freighter/pkg/proxy"
	"github.com/horvatic/freighter/pkg/request"
	"io"
	"net/http"
)

var store *datastore.MemoryDataStore
var gatewayConfig *config.GatewayConfig

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
	defer req.Body.Close()
	response, statusCode := gateway.ProcessRequest(request.NewRequest(req.Method, req.URL.Path, req.URL.Query(), req.Header, req.Body), &proxy.RequestProxy{}, store, gatewayConfig)
	defer response.Close()
	w.WriteHeader(statusCode)
	io.Copy(w, response)
}

func Start() {
	gatewayConfig = config.BuildGatewayConfig()
	store = datastore.NewMemoryDataStore()
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/config", handleConfig)
	http.ListenAndServe(":"+gatewayConfig.Port, nil)
}
