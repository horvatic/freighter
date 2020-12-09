package mock

import (
	"github.com/horvatic/freighter/pkg/datastore"
)

type MockDataStore struct {
	Service datastore.Service
	ApiKey  string
	Error   error
}

func (m *MockDataStore) GetService(serviceName string) (*datastore.Service, error) {
	return &m.Service, m.Error
}

func (m *MockDataStore) SetService(service *datastore.Service) error {
	m.Service = *service
	return m.Error
}

func (m *MockDataStore) RemoveService(serviceName string) error {
	return m.Error
}

func (m *MockDataStore) GetAllServices() []datastore.Service {
	return []datastore.Service{m.Service}
}

func (m *MockDataStore) GetApiKey() (string, error) {
	return m.ApiKey, nil
}

func (m *MockDataStore) SetApiKey(apiKey string) error {
	m.ApiKey = apiKey
	return nil
}
