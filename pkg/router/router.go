package router

import (
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"github.com/horvatic/freighter/pkg/request"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func getQuery(query string) string {
	if query == "" {
		return ""
	}
	return "?" + query
}

func Route(req *request.Request, p proxy.Proxy, d datastore.DataStore) (io.ReadCloser, int) {
	parts := strings.SplitAfter(strings.TrimLeft(req.UriPath, "/"), "/")
	s, serr := d.GetService(strings.TrimRight(parts[0], "/"))
	if serr != nil {
		return ioutil.NopCloser(strings.NewReader(serr.Error())), http.StatusInternalServerError
	}

	body, err, statusCode := p.DoRequest(req.Method, s.Host+":"+s.Port+"/"+strings.Join(parts[1:], "")+getQuery(req.RawQuery), req.Headers, req.Body)
	if err != nil {
		if body != nil {
			body.Close()
		}
		return ioutil.NopCloser(strings.NewReader(err.Error())), http.StatusInternalServerError
	}
	return body, statusCode
}
