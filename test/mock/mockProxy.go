package mock

import (
	"io"
)

type MockProxy struct {
	Body       io.ReadCloser
	Error      error
	StatusCode int
}

var ProxyUriRequest string

func (p *MockProxy) GetRequest(uri string) (io.ReadCloser, error, int) {
	ProxyUriRequest = uri
	return p.Body, p.Error, p.StatusCode
}
