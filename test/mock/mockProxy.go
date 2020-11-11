package mock

import (
	"io"
	"net/http"
)

type MockProxy struct {
	Body       io.ReadCloser
	Error      error
	StatusCode int
}

var ProxyUriRequest string
var ProxyHeaderRequest http.Header

func (p *MockProxy) GetRequest(uri string, headers http.Header) (io.ReadCloser, error, int) {
	ProxyUriRequest = uri
	ProxyHeaderRequest = headers
	return p.Body, p.Error, p.StatusCode
}
