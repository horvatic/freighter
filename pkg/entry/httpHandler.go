package entry

import (
	"github.com/horvatic/freighter/pkg/config"
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/gateway"
	"github.com/horvatic/freighter/pkg/profile"
	"github.com/horvatic/freighter/pkg/proxy"
	"github.com/horvatic/freighter/pkg/request"
	"io"
	"net/http"
)

var store *datastore.MemoryDataStore
var gatewayConfig *config.GatewayConfig

func handleRequest(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	response, statusCode := gateway.ProcessRequest(request.NewRequest(req.Method, req.URL.Path, req.URL.Query(), req.URL.RawQuery, req.Header, req.Body), &proxy.RequestProxy{}, store, gatewayConfig)
	defer response.Close()
	w.WriteHeader(statusCode)
	io.Copy(w, response)
}

func Start() {
	gatewayConfig = config.BuildGatewayConfig()
	store = datastore.NewMemoryDataStore(gatewayConfig)
	profile := profile.NewProfile(store)
	http.HandleFunc("/", handleRequest)
	profile.Setup()
	http.ListenAndServe(":"+gatewayConfig.Port, nil)
}
