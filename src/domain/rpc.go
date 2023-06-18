package domain

import (
	"github.com/pkg/errors"
)

type JsonRpcRequest struct {
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	JsonRpc string      `json:"jsonrpc"`
	Params  interface{} `json:"params,omitempty"`
}

func NewJsonRpcRequest(id int, method string, params interface{}) (*JsonRpcRequest, error) {
	if id < 0 {
		return nil, errors.Wrap(ErrInvalidArgument, "json rpc request 'id' should be greater than zero")
	}
	if len(method) == 0 {
		return nil, errors.Wrap(ErrInvalidArgument, "json rpc request 'method' should not be empty")
	}
	return &JsonRpcRequest{
		ID:      id,
		Method:  method,
		JsonRpc: "2.0",
		Params:  params,
	}, nil
}

type JsonRpcResponse struct {
	ID      int           `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Result  string        `json:"result,omitempty"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

type JsonRpcError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
