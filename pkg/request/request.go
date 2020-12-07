package request

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	UriPath string
	Headers http.Header
	Query   url.Values
	RawQuery string
	Body    io.ReadCloser
	Method  string
}

func NewRequest(method string, uriPath string, query url.Values, rawQuery string, headers http.Header, body io.ReadCloser) *Request {
	return &Request{
		UriPath: uriPath,
		Query:   query,
		RawQuery: rawQuery,
		Headers: headers,
		Body:    body,
		Method:  method,
	}
}
