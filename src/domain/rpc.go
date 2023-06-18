package domain

type RpcRequest struct {
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	JsonRpc string      `json:"jsonrpc"`
	Params  interface{} `json:"params,omitempty"`
}

type RpcResponse struct {
	ID      int      `json:"id"`
	JsonRpc string   `json:"jsonrpc"`
	Result  string   `json:"result,omitempty"`
	Error   RpcError `json:"error,omitempty"`
}

type RpcError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
