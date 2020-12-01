package profile

import (
	"fmt"
	"net/http"
)

func (p *profile) register(w http.ResponseWriter, req *http.Request) {
	s, code := register(req.Body, p)
	w.WriteHeader(code)
	fmt.Fprint(w, s)
}

func (p *profile) getAllServices(w http.ResponseWriter, req *http.Request) {
	s, code := getAllServices(p)
	w.WriteHeader(code)
	fmt.Fprint(w, s)
}

func (p *profile) Setup() {
	http.HandleFunc("/config/register", p.register)
	http.HandleFunc("/config/currentservices", p.getAllServices)
}
