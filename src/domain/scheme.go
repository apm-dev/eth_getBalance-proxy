package domain

import "time"

type Scheme interface {
	ParseRequest(req *JsonRpcRequest) (fastReply *JsonRpcResponse, err error)
	IsSupportedRpcMethod(method string) bool
	IsCacheSupported(req *JsonRpcRequest) (ttl time.Duration, ok bool)
	IsJsonRpcResponseValid(resp *JsonRpcResponse) bool
}
