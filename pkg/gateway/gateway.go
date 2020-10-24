package gateway

import (
	"github.com/horvatic/freighter/pkg/proxy"
)

type Request struct {
	UriPath string
}

func route(req *Request, p proxy.Proxy) string {
	body, err := p.GetRequest("http://127.0.0.1:8000/" + req.UriPath)
	if err != nil {
		//Add error handling
		return "Could not complete request"
	}
	return string(body)
}
