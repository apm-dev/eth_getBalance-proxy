package jsonrpc

import (
	"fmt"
	"strings"
	"time"

	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
)

type JsonRpcScheme struct {
	cachingMethods      map[string]time.Duration
	blacklistedMethods  map[string]bool
	blacklistedPrefixes map[string]bool
}

func NewJsonRpcScheme(
	cachingMethods map[string]time.Duration,
	blacklistedMethods map[string]bool,
	blacklistedPrefixes map[string]bool,
) *JsonRpcScheme {
	return &JsonRpcScheme{
		cachingMethods:      cachingMethods,
		blacklistedMethods:  blacklistedMethods,
		blacklistedPrefixes: blacklistedPrefixes,
	}
}

func (s *JsonRpcScheme) ParseRequest(req *domain.JsonRpcRequest) (fastReply *domain.JsonRpcResponse, err error) {
	if s.IsSupportedRpcMethod(req.Method) {
		return nil, nil
	}
	return s.BuildUnsupportedMethodResponse(req), nil
}

func (s *JsonRpcScheme) IsSupportedRpcMethod(method string) bool {
	methodData := strings.Split(method, "_")
	if len(methodData) != 2 {
		return false
	}
	if _, ok := s.blacklistedPrefixes[methodData[0]]; ok {
		return false
	}
	if _, ok := s.blacklistedMethods[method]; ok {
		return false
	}
	return true
}

func (s *JsonRpcScheme) IsCacheSupported(req *domain.JsonRpcRequest) (time.Duration, bool) {
	ttl, ok := s.cachingMethods[req.Method]
	return ttl, ok
}

func (s *JsonRpcScheme) IsJsonRpcResponseValid(resp *domain.JsonRpcResponse) bool {
	if resp.JsonRpc != "2.0" {
		return false
	}
	return (resp.Result != "" && resp.Error == nil) || (resp.Result == "" && resp.Error != nil)
}

func (s *JsonRpcScheme) BuildUnsupportedMethodResponse(req *domain.JsonRpcRequest) *domain.JsonRpcResponse {
	return &domain.JsonRpcResponse{
		ID:      req.ID,
		JsonRpc: "2.0",
		Error: &domain.JsonRpcError{
			Code:    -32601,
			Message: fmt.Sprintf("the method %s does not exist/is not available", req.Method),
		},
	}
}
