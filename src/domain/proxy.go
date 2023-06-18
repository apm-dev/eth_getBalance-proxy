package domain

import (
	"context"
	"time"
)

type RpcProxyService interface {
	SendRequest(c context.Context, blockchain string, req *JsonRpcRequest) (*JsonRpcResponse, error)
}

type RpcProxyCache interface {
	GetCachedResponse(blockchain string, req *JsonRpcRequest) (*JsonRpcResponse, bool)
	CacheResponse(blockchain string, req *JsonRpcRequest, rsp *JsonRpcResponse, ttl time.Duration)
}
