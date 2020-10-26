package gateway

import (
	"errors"
	"github.com/horvatic/freighter/test/mock"
	"net/http"
	"testing"
)

func TestRoute(t *testing.T) {
	testBody := "test123"

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: []byte(testBody), Error: nil, StatusCode: http.StatusOK},
		&mock.MockDataStore{})

	if result != testBody {
		t.Errorf("got %q want %q", result, testBody)
	}

	if StatusCode != http.StatusOK {
		t.Errorf("got %q want %q", http.StatusOK, StatusCode)
	}
}

func TestRoute404(t *testing.T) {
	testBody := "test123"

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: []byte(testBody), Error: nil, StatusCode: http.StatusNotFound},
		&mock.MockDataStore{})

	if result != testBody {
		t.Errorf("got %q want %q", result, testBody)
	}

	if StatusCode != http.StatusNotFound {
		t.Errorf("got %q want %q", http.StatusNotFound, StatusCode)
	}
}

func TestRoute500(t *testing.T) {
	testBody := "test123"

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: []byte(testBody), Error: nil, StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{})

	if result != testBody {
		t.Errorf("got %q want %q", result, testBody)
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}

func TestRoute500CallingProxy(t *testing.T) {
	testBody := "Could not complete request"

	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: nil, Error: errors.New("error"), StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{})

	if result != testBody {
		t.Errorf("got %q want %q", result, testBody)
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}

func TestRouteCantFindService(t *testing.T) {
	result, StatusCode := route(&Request{UriPath: "test"},
		&mock.MockProxy{Body: nil, Error: nil, StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{Error: errors.New("error")})

	if result != "error" {
		t.Errorf("got %q want %q", result, "error")
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}
