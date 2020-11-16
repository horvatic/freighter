package gateway

import (
	"github.com/horvatic/freighter/pkg/config"
	"github.com/horvatic/freighter/pkg/datastore"
	"github.com/horvatic/freighter/pkg/proxy"
	"github.com/horvatic/freighter/pkg/request"
	"github.com/horvatic/freighter/pkg/router"
	"io"
	"io/ioutil"
	"strings"
	"fmt"
)

func ProcessRequest(req *request.Request, p proxy.Proxy, d datastore.DataStore, config *config.GatewayConfig) (io.ReadCloser, int) {
	key := getApiKeyFromHeader(req.Headers)
	fmt.Println(req.Headers)
	if !hasValidApiKey(config, key) {
		return ioutil.NopCloser(strings.NewReader("Api key failed to validate")), 403
	}
	req.Headers = removeApiKeyHeader(req.Headers)
	return router.Route(req, p, d)
}
