package mock

import (
	"io"
	"io/ioutil"
	"net/http"
)

type MockProxy struct {
	Body       io.ReadCloser
	Error      error
	StatusCode int
}

var ProxyUriRequest string
var ProxyHeaderRequest http.Header
var ProxyBodyRequest string
var ProxyMethod string

func (p *MockProxy) DoRequest(method string, uri string, headers http.Header, body io.ReadCloser) (io.ReadCloser, error, int) {
	ProxyUriRequest = uri
	ProxyHeaderRequest = headers
	if body != nil {
		bodyBytes, _ := ioutil.ReadAll(body)
		ProxyBodyRequest = string(bodyBytes)
	}
	ProxyMethod = method
	return p.Body, p.Error, p.StatusCode
}
