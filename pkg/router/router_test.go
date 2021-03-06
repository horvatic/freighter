package router

import (
	"errors"
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/request"
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
	query := make(map[string][]string)
	rawQuery := "hello=world&hello=444"
	headers := make(map[string][]string)
	query["hello"] = []string{"world", "444"}
	headers["head"] = []string{"headerinfo"}
	body := ioutil.NopCloser(strings.NewReader("test body"))
	defer body.Close()
	result, StatusCode := Route(request.NewRequest("POST", "servicename/test", query, rawQuery, headers, body),
		&mock.MockProxy{Body: testBody, Error: nil, StatusCode: http.StatusOK},
		&mock.MockDataStore{Service: datastore.Service{ServiceName: "test", Host: "1.1.1.1", Port: "8080"}, Error: nil})
	defer result.Close()
	sResult := getString(result)
	if sResult != "test123" {
		t.Errorf("got %q want %q", sResult, "test123")
	}

	if "1.1.1.1:8080/test?hello=world&hello=444" != mock.ProxyUriRequest {
		t.Errorf("got %q want %q", mock.ProxyUriRequest, "1.1.1.1:8080/test?hello=world&hello=444")
	}

	if StatusCode != http.StatusOK {
		t.Errorf("got %q want %q", http.StatusOK, StatusCode)
	}

	if mock.ProxyHeaderRequest["head"][0] != "headerinfo" {
		t.Errorf("got %q want %q", mock.ProxyHeaderRequest["head"][0], "headerinfo")
	}

	if mock.ProxyBodyRequest != "test body" {
		t.Errorf("got %q want %q", mock.ProxyBodyRequest, "test body")
	}

	if mock.ProxyMethod != "POST" {
		t.Errorf("got %q want %q", mock.ProxyMethod, "POST")
	}
}

func TestRoute404(t *testing.T) {
	testBody := ioutil.NopCloser(strings.NewReader("test123"))
	defer testBody.Close()

	result, StatusCode := Route(&request.Request{UriPath: "test", Query: nil},
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

	result, StatusCode := Route(&request.Request{UriPath: "test", Query: nil},
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

	result, StatusCode := Route(&request.Request{UriPath: "test", Query: nil},
		&mock.MockProxy{Body: nil, Error: errors.New("error"), StatusCode: http.StatusInternalServerError},
		&mock.MockDataStore{})

	defer result.Close()
	sResult := getString(result)
	if sResult != "error" {
		t.Errorf("got %q want %q", sResult, "error")
	}

	if StatusCode != http.StatusInternalServerError {
		t.Errorf("got %q want %q", http.StatusInternalServerError, StatusCode)
	}
}

func TestRouteCantFindService(t *testing.T) {
	result, StatusCode := Route(&request.Request{UriPath: "test", Query: nil},
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
