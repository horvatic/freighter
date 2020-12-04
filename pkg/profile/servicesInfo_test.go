package profile

import (
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/test/mock"
	"net/http"
	"testing"
)

func TestGetAllServices(t *testing.T) {
	testService := datastore.Service{ServiceName: "test", Host: "1.1.1.1", Port: "8080"}
	datastore := &mock.MockDataStore{Service: testService, Error: nil}
	profile := NewProfile(datastore)

	message, returnCode := getAllServices(profile)

	want := "[{\"serviceName\":\"test\",\"host\":\"1.1.1.1\",\"port\":\"8080\"}]"
	if want != message {
		t.Errorf("got %q want %q", message, want)
	}
	if http.StatusOK != returnCode {
		t.Errorf("got %q want %q", returnCode, http.StatusOK)
	}
}
