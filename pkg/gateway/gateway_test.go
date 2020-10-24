package gateway

import (
  "testing"
)

type FakeProxy struct {
  Body []byte
  Error error
}

func (p *FakeProxy) GetRequest(uri string) ([]byte, error) {
  return p.Body, p.Error
}

func TestRoute(t *testing.T) {
  testBody := "test123"

  result := route(&Request{UriPath: "test"}, &FakeProxy{Body: []byte(testBody), Error: nil})

  if result != testBody {
      t.Errorf("got %q want %q", result, testBody)
  }
}
