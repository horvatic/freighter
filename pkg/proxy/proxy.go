package proxy

import (
	"io"
	"net/http"
)

type Proxy interface {
	GetRequest(uri string, headers http.Header) (io.ReadCloser, error, int)
}

type RequestProxy struct {
}

func (p *RequestProxy) GetRequest(uri string, headers http.Header) (io.ReadCloser, error, int) {
	req, err := http.NewRequest("GET", uri, nil)
	setHeaders(headers, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return resp.Body, nil, resp.StatusCode
}

func setHeaders(headers http.Header, req *http.Request) {
	if headers != nil {
		for key, elements := range headers {
			for _, element := range elements {
				req.Header[key] = append(req.Header[key], element)
			}
		}
	}
}
