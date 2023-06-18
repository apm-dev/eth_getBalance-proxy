package scheme

import (
	"strings"

	"github.com/apm-dev/eth_getBalance-proxy/src/config"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/apm-dev/eth_getBalance-proxy/src/scheme/eth"
	"github.com/pkg/errors"
)

func CreateScheme(config *config.Config, blockchain string) (domain.Scheme, error) {
	switch strings.ToLower(blockchain) {
	case ETH:
		return eth.NewEthScheme(config), nil
	default:
		return nil, errors.Wrapf(domain.ErrInvalidArgument, "not supported scheme '%s'", blockchain)
	}
}
