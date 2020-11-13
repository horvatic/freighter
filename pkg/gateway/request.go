package gateway

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	UriPath string
	Headers http.Header
	Query   url.Values
	Body    io.ReadCloser
	Method  string
}

func NewRequest(method string, uriPath string, query url.Values, headers http.Header, body io.ReadCloser) *Request {
	return &Request{
		UriPath: uriPath,
		Query:   query,
		Headers: headers,
		Body:    body,
		Method:  method,
	}
}
