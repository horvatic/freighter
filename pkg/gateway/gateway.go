package gateway

import (
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func route(req *Request, p proxy.Proxy, d datastore.DataStore) (io.ReadCloser, int) {
	parts := strings.SplitAfter(strings.TrimLeft(req.UriPath, "/"), "/")
	s, serr := d.GetService(strings.TrimRight(parts[0], "/"))
	if serr != nil {
		return ioutil.NopCloser(strings.NewReader(serr.Error())), http.StatusInternalServerError
	}
	q := ""
	if req.Query != nil {
		for key, elements := range req.Query {
			for _, element := range elements {
				if q != "" {
					q = q + "&"
				}
				q = q + key + "=" + element
			}
		}
	}
	if q != "" {
		q = "?" + q
	}
	body, err, statusCode := p.GetRequest(s.Host + ":" + s.Port + "/" + strings.Join(parts[1:], "") + q)
	if err != nil {
		if body != nil {
			body.Close()
		}
		return ioutil.NopCloser(strings.NewReader(err.Error())), http.StatusInternalServerError
	}
	return body, statusCode
}
