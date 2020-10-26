package datastore

type Service struct {
	ServiceName string `json:"serviceName"`
	Host        string `json:"host"`
	Port        string `json:"port"`
}
