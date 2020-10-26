package gateway

import (
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"net/http"
	"strings"
)

type Request struct {
	UriPath string
}

func route(req *Request, p proxy.Proxy, d datastore.DataStore) (string, int) {
	parts := strings.SplitAfter(strings.TrimLeft(req.UriPath, "/"), "/")
	s, serr := d.GetService(strings.TrimRight(parts[0], "/"))
	if serr != nil {
		return serr.Error(), http.StatusInternalServerError
	}
	body, err, statusCode := p.GetRequest("http://" + s.Host + ":" + s.Port + "/" + strings.Join(parts[1:], ""))
	if err != nil {
		return "Could not complete request", http.StatusInternalServerError
	}
	return string(body), statusCode
}
