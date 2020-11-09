package gateway

import (
	"net/url"
)

type Request struct {
	UriPath string
	Query   url.Values
}

func NewRequest(uriPath string, query url.Values) *Request {
	return &Request{
		UriPath: uriPath,
		Query:   query,
	}
}
