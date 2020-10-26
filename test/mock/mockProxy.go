package mock

type MockProxy struct {
	Body       []byte
	Error      error
	StatusCode int
}

func (p *MockProxy) GetRequest(uri string) ([]byte, error, int) {
	return p.Body, p.Error, p.StatusCode
}
