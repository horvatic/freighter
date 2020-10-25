package proxy

import (
  "io/ioutil"
  "net/http"
)

type Proxy interface {
  GetRequest(uri string) ([] byte, error, int)
}

type RequestProxy struct {
}

func (p *RequestProxy) GetRequest(uri string) ([] byte, error, int) {
  resp, err := http.Get(uri)
  if err != nil {
    return nil, err, http.StatusInternalServerError
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err, http.StatusInternalServerError
  }

  return body, nil, resp.StatusCode
}
