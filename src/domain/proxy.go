package domain

import "context"

type RpcProxyService interface {
	SendRequest(c context.Context, blockchain string, req *RpcRequest) (*RpcResponse, error)
}
