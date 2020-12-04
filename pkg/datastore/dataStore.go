package datastore

type DataStore interface {
	SetService(service *Service) error
	GetService(serviceName string) (*Service, error)
	RemoveService(serviceName string) error
	GetAllServices() []Service
}
