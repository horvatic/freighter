package mock

import (
	"github.com/horvatic/freighter/pkg/datastore"
)

type MockDataStore struct {
	Service datastore.Service
	Error   error
}

func (m *MockDataStore) GetService(serviceName string) (*datastore.Service, error) {
	return &m.Service, m.Error
}

func (m *MockDataStore) SetService(service *datastore.Service) error {
	return m.Error
}

func (m *MockDataStore) RemoveService(serviceName string) error {
	return m.Error
}
