package proxy

import (
	"io"
	"net/http"
)

type Proxy interface {
	GetRequest(uri string) (io.ReadCloser, error, int)
}

type RequestProxy struct {
}

func (p *RequestProxy) GetRequest(uri string) (io.ReadCloser, error, int) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return resp.Body, nil, resp.StatusCode
}
