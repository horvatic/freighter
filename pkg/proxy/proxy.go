package proxy

import (
  "io/ioutil"
  "net/http"
)

type Proxy interface {
  GetRequest(uri string) ([] byte, error)
}

type RequestProxy struct {
}

func (p *RequestProxy) GetRequest(uri string) ([] byte, error) {
  resp, err := http.Get(uri)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return body, nil
}
