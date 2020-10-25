package gateway

import (
  "net/http"
  "testing"
)

type FakeProxy struct {
  Body []byte
  Error error
  StatusCode int
}

func (p *FakeProxy) GetRequest(uri string) ([]byte, error, int) {
  return p.Body, p.Error, p.StatusCode
}

func TestRoute(t *testing.T) {
  testBody := "test123"

  result, StatusCode := route(&Request{UriPath: "test"}, &FakeProxy{Body: []byte(testBody), Error: nil, StatusCode: http.StatusOK})

  if result != testBody {
      t.Errorf("got %q want %q", result, testBody)
  }

  if StatusCode != http.StatusOK {
    t.Errorf("got %q want %q", http.StatusOK, StatusCode)
  }
}

func TestRoute404(t *testing.T) {
  testBody := "test123"

  result, StatusCode := route(&Request{UriPath: "test"}, &FakeProxy{Body: []byte(testBody), Error: nil, StatusCode: http.StatusNotFound})

  if result != testBody {
      t.Errorf("got %q want %q", result, testBody)
  }

  if StatusCode != http.StatusNotFound {
    t.Errorf("got %q want %q", http.StatusNotFound, StatusCode)
  }
}

func TestRoute500(t *testing.T) {
  testBody := "test123"

  result, StatusCode := route(&Request{UriPath: "test"}, &FakeProxy{Body: []byte(testBody), Error: nil, StatusCode: http.StatusInternalServerError})

  if result != testBody {
      t.Errorf("got %q want %q", result, testBody)
  }

  if StatusCode != http.StatusInternalServerError {
    t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
  }
}
