package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllServices(p *profile) (string, int) {
	services := p.store.GetAllServices()
	result, err := json.Marshal(services)
	if err != nil {
		return fmt.Sprintf("{ \"error\": \"" + err.Error() + "\"}"), http.StatusInternalServerError
	} else {
		return string(result), http.StatusOK
	}
}
