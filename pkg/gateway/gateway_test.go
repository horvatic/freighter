package gateway

import (
	"errors"
	"github.com/horvatic/freighter/test/mock"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRoute(t *testing.T) {
	testBody := ioutil.NopCloser(strings.NewReader("test123"))
	defer testBody.Close()
	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: testBody, Error: nil, StatusCode: http.StatusOK},
		&mock.MockDataStore{})
	defer result.Close()
	sResult := getString(result)
	if sResult != "test123" {
		t.Errorf("got %q want %q", sResult, "test123")
	}

	if StatusCode != http.StatusOK {
		t.Errorf("got %q want %q", http.StatusOK, StatusCode)
	}
}

func TestRoute404(t *testing.T) {
	testBody := ioutil.NopCloser(strings.NewReader("test123"))
	defer testBody.Close()

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: testBody, Error: nil, StatusCode: http.StatusNotFound},
		&mock.MockDataStore{})
	defer result.Close()
	sResult := getString(result)
	if sResult != "test123" {
		t.Errorf("got %q want %q", sResult, "test123")
	}

	if StatusCode != http.StatusNotFound {
		t.Errorf("got %q want %q", http.StatusNotFound, StatusCode)
	}
}

func TestRoute500(t *testing.T) {
	testBody := ioutil.NopCloser(strings.NewReader("test123"))
	defer testBody.Close()

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: testBody, Error: nil, StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{})
	defer result.Close()
	sResult := getString(result)
	if sResult != "test123" {
		t.Errorf("got %q want %q", sResult, "test123")
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}

func TestRoute500CallingProxy(t *testing.T) {
	testBody := ioutil.NopCloser(strings.NewReader("Could not complete request"))
	defer testBody.Close()

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: nil, Error: errors.New("error"), StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{})

	defer result.Close()
	sResult := getString(result)
	if sResult != "Could not complete request" {
		t.Errorf("got %q want %q", sResult, "Could not complete request")
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}

func TestRouteCantFindService(t *testing.T) {
	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: nil, Error: nil, StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{Error: errors.New("error")})
	defer result.Close()
	sResult := getString(result)
	if sResult != "error" {
		t.Errorf("got %q want %q", sResult, "error")
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}

func getString(stream io.ReadCloser) string {
	s, _ := ioutil.ReadAll(stream)
	return string(s)
}
