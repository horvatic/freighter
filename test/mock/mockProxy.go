package mock

import (
	"io"
)

type MockProxy struct {
	Body       io.ReadCloser
	Error      error
	StatusCode int
}

func (p *MockProxy) GetRequest(uri string) (io.ReadCloser, error, int) {
	return p.Body, p.Error, p.StatusCode
}
