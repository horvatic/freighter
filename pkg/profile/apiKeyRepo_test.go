package profile

import (
	"errors"
	"fmt"
	"github.com/horvatic/freighter/test/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSetApiKey(t *testing.T) {
	datastore := &mock.MockDataStore{Error: nil, ApiKey: ""}
	profile := NewProfile(datastore)
	body := ioutil.NopCloser(strings.NewReader("{ \"apiKey\" : \"a1S1Fvfve\" }"))
	defer body.Close()

	message, returnCode := setApiKey(body, profile)

	want := fmt.Sprintf("{\"apiKey\":\"%+v\"}", "a1S1Fvfve")
	if want != message {
		t.Errorf("got %q want %q", message, want)
	}
	if http.StatusOK != returnCode {
		t.Errorf("got %q want %q", returnCode, http.StatusOK)
	}
}

func TestSetApiKeyError(t *testing.T) {
	datastore := &mock.MockDataStore{Error: errors.New("error")}
	profile := NewProfile(datastore)
	body := ioutil.NopCloser(strings.NewReader("bdfdsdfffwefwefweewffwe"))
	defer body.Close()

	_, returnCode := setApiKey(body, profile)

	if http.StatusBadRequest != returnCode {
		t.Errorf("got %q want %q", returnCode, http.StatusBadRequest)
	}
}

func TestGetApiKey(t *testing.T) {
	datastore := &mock.MockDataStore{ApiKey: "dwerweffr32fw", Error: nil}
	profile := NewProfile(datastore)

	message, returnCode := getApiKey(profile)

	want := "{\"apiKey\":\"dwerweffr32fw\"}"
	if want != message {
		t.Errorf("got %q want %q", message, want)
	}
	if http.StatusOK != returnCode {
		t.Errorf("got %q want %q", returnCode, http.StatusOK)
	}
}
