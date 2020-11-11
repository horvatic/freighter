package gateway

import (
	"net/http"
	"net/url"
)

type Request struct {
	UriPath string
	Headers http.Header
	Query   url.Values
}

func NewRequest(uriPath string, query url.Values, headers http.Header) *Request {
	return &Request{
		UriPath: uriPath,
		Query:   query,
		Headers: headers,
	}
}
