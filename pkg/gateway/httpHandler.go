package gateway

import 	(
  "fmt"
  "net/http"
  "github.com/horvatic/freighter/pkg/proxy"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
  response := route(&Request{UriPath: req.URL.Path}, &proxy.RequestProxy{})
  fmt.Fprintf(w, response)
}

func Start() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
