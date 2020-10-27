package gateway

import (
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	UriPath string
}

func route(req *Request, p proxy.Proxy, d datastore.DataStore) (io.ReadCloser, int) {
	parts := strings.SplitAfter(strings.TrimLeft(req.UriPath, "/"), "/")
	s, serr := d.GetService(strings.TrimRight(parts[0], "/"))
	if serr != nil {
		return ioutil.NopCloser(strings.NewReader(serr.Error())), http.StatusInternalServerError
	}
	body, err, statusCode := p.GetRequest("http://" + s.Host + ":" + s.Port + "/" + strings.Join(parts[1:], ""))
	if err != nil {
		if body != nil {
			body.Close()
		}
		return ioutil.NopCloser(strings.NewReader("Could not complete request")), http.StatusInternalServerError
	}
	return body, statusCode
}
