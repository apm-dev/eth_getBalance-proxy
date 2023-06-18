package eth

import (
	"sync"

	"github.com/apm-dev/eth_getBalance-proxy/src/config"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/apm-dev/eth_getBalance-proxy/src/scheme/jsonrpc"
)

type ethScheme struct {
	*jsonrpc.JsonRpcScheme
	config *config.Config
}

var scheme *ethScheme

var once = &sync.Once{}

func NewEthScheme(config *config.Config) domain.Scheme {
	once.Do(func() {
		scheme = &ethScheme{
			config:        config,
			JsonRpcScheme: jsonrpc.NewJsonRpcScheme(cachingMethods, blacklistedMethods, blacklistedPrefixes),
		}
	})
	return scheme
}
