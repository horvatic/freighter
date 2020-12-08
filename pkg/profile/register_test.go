package profile

import (
	"errors"
	"github.com/horvatic/freighter/test/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	datastore := &mock.MockDataStore{Error: nil}
	profile := NewProfile(datastore)
	body := ioutil.NopCloser(strings.NewReader("{ \"serviceName\" : \"test\", \"host\" : \"1.1.1.1\", \"port\": \"8080\" }"))
	defer body.Close()

	message, returnCode := register(body, profile)

	want := "{\"serviceName\":\"test\",\"host\":\"1.1.1.1\",\"port\":\"8080\"}"
	if want != message {
		t.Errorf("got %q want %q", message, want)
	}
	if http.StatusOK != returnCode {
		t.Errorf("got %q want %q", returnCode, http.StatusOK)
	}
}

func TestRegisterError(t *testing.T) {
	datastore := &mock.MockDataStore{Error: errors.New("error")}
	profile := NewProfile(datastore)
	body := ioutil.NopCloser(strings.NewReader("bdfdsdfffwefwefweewffwe"))
	defer body.Close()

	_, returnCode := register(body, profile)

	if http.StatusBadRequest != returnCode {
		t.Errorf("got %q want %q", returnCode, http.StatusBadRequest)
	}
}
