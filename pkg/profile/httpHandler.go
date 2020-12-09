package profile

import (
	"fmt"
	"io"
	"net/http"
)

type action func(p *profile) (string, int)
type actionWithBody func(body io.Reader, p *profile) (string, int)

func (p *profile) buildRouteActionWithBody(fn actionWithBody) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		s, code := fn(req.Body, p)
		w.WriteHeader(code)
		fmt.Fprint(w, s)
	}
}

func (p *profile) buildRouteAction(fn action) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		s, code := fn(p)
		w.WriteHeader(code)
		fmt.Fprint(w, s)
	}
}

func (p *profile) Setup() {
	http.HandleFunc("/config/register", p.buildRouteActionWithBody(register))
	http.HandleFunc("/config/currentservices", p.buildRouteAction(getAllServices))
	http.HandleFunc("/config/setapikey", p.buildRouteActionWithBody(setApiKey))
	http.HandleFunc("/config/getapikey", p.buildRouteAction(getApiKey))
}
