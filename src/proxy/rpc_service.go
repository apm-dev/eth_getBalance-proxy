package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/apm-dev/eth_getBalance-proxy/src/config"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/apm-dev/eth_getBalance-proxy/src/scheme"
)

type rpcProxyService struct {
	config   *config.Config
	cache    domain.RpcProxyCache
	nodeRepo domain.NodeRepository
}

func NewRpcProxyService(
	config *config.Config,
	cache domain.RpcProxyCache,
	nr domain.NodeRepository,
) domain.RpcProxyService {
	return &rpcProxyService{
		config:   config,
		cache:    cache,
		nodeRepo: nr,
	}
}

func (s *rpcProxyService) SendRequest(c context.Context, blockchain string, req *domain.JsonRpcRequest) (*domain.JsonRpcResponse, error) {
	// create scheme
	scheme, err := scheme.CreateScheme(s.config, blockchain)
	if err != nil {
		return nil, err
	}
	// parse request
	fastReply, err := scheme.ParseRequest(req)
	if err != nil {
		return nil, err
	}
	if fastReply != nil {
		return fastReply, nil
	}
	// reply from cache
	cacheTTL, isCacheSupported := scheme.IsCacheSupported(req)
	if isCacheSupported {
		item, ok := s.cache.GetCachedResponse(blockchain, req)
		if ok {
			return item, nil
		}
	}
	// get available nodes for requested blockchain
	nodes, err := s.nodeRepo.GetNodesByBlockchain(blockchain)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, errors.Wrapf(domain.ErrInternalServer, "there is no available nodes for '%s'", blockchain)
	}
	// get the fastest node (assumes that nodes are sorted by weight)
	node := nodes[0]
	// prepare request to node
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServer, err.Error())
	}
	ctx, cancel := context.WithTimeout(c, s.config.App.OutgoingRequestTimeout)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, node.Url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServer, err.Error())
	}
	request.Header.Add("Content-Type", "application/json")
	// send request to node and parse response
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServer, err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		rsp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrapf(domain.ErrInternalServer, "request to '%s' node '%s' failed with status: '%d'", blockchain, node.Url, resp.StatusCode)
		} else {
			return nil, errors.Wrapf(domain.ErrInternalServer, "request to '%s' node '%s' failed with status: '%d', and message '%s'", blockchain, node.Url, resp.StatusCode, string(rsp))
		}
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServer, err.Error())
	}
	var rpcResp domain.JsonRpcResponse
	err = json.Unmarshal(respBody, &rpcResp)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternalServer, err.Error())
	}
	// validate response
	if !scheme.IsJsonRpcResponseValid(&rpcResp) {
		return nil, errors.Wrap(domain.ErrInternalServer, err.Error())
	}
	// cache the response
	if isCacheSupported && rpcResp.Result != "" {
		s.cache.CacheResponse(blockchain, req, &rpcResp, cacheTTL)
	}
	return &rpcResp, nil
}
