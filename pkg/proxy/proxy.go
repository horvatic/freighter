package proxy

import (
	"io"
	"net/http"
)

type Proxy interface {
	DoRequest(method string, uri string, headers http.Header, body io.ReadCloser) (io.ReadCloser, error, int)
}

type RequestProxy struct {
}

func (p *RequestProxy) DoRequest(method string, uri string, headers http.Header, body io.ReadCloser) (io.ReadCloser, error, int) {
	req, err := http.NewRequest(method, uri, body)
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
