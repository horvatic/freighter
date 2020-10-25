package gateway

import (
	"fmt"
	"github.com/horvatic/freighter/pkg/proxy"
	"net/http"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	response, statusCode := route(&Request{UriPath: req.URL.Path}, &proxy.RequestProxy{})
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, response)
}

func Start() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
