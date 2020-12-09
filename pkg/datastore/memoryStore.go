package datastore

import (
	"errors"
	"github.com/horvatic/freighter/pkg/config"
)

type MemoryDataStore struct {
	services []Service
	apiKey   string
}

func NewMemoryDataStore(config *config.GatewayConfig) *MemoryDataStore {
	m := MemoryDataStore{services: []Service{}, apiKey: config.InitAuthKey}

	return &m
}

func (m *MemoryDataStore) GetService(serviceName string) (*Service, error) {
	for _, v := range m.services {
		if v.ServiceName == serviceName {
			return &v, nil
		}
	}

	return &Service{}, errors.New("No service found")
}

func (m *MemoryDataStore) SetService(service *Service) error {
	_, err := m.GetService(service.ServiceName)
	if err == nil {
		return errors.New("Service already in the list")
	}
	m.services = append(m.services, *service)
	return nil
}

func (m *MemoryDataStore) RemoveService(serviceName string) error {
	removeIndex := -1
	for i, v := range m.services {
		if v.ServiceName == serviceName {
			removeIndex = i
			break
		}
	}
	if removeIndex == -1 {
		return errors.New("Service not found in the list")
	}
	m.services = append(m.services[:removeIndex], m.services[removeIndex+1:]...)
	return nil
}

func (m *MemoryDataStore) GetAllServices() []Service {
	return m.services
}

func (m *MemoryDataStore) GetApiKey() (string, error) {
	return m.apiKey, nil
}

func (m *MemoryDataStore) SetApiKey(apiKey string) error {
	m.apiKey = apiKey
	return nil
}
