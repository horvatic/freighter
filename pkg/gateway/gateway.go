package gateway

import (
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getQuery(query url.Values) (string) {
	q := ""
	if query != nil {
		for key, elements := range query {
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
	return q
}

func route(req *Request, p proxy.Proxy, d datastore.DataStore) (io.ReadCloser, int) {
	parts := strings.SplitAfter(strings.TrimLeft(req.UriPath, "/"), "/")
	s, serr := d.GetService(strings.TrimRight(parts[0], "/"))
	if serr != nil {
		return ioutil.NopCloser(strings.NewReader(serr.Error())), http.StatusInternalServerError
	}

	body, err, statusCode := p.GetRequest(s.Host + ":" + s.Port + "/" + strings.Join(parts[1:], "") + getQuery(req.Query))
	if err != nil {
		if body != nil {
			body.Close()
		}
		return ioutil.NopCloser(strings.NewReader(err.Error())), http.StatusInternalServerError
	}
	return body, statusCode
}
