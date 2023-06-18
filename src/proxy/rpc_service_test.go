package proxy_test

import (
	"context"
	"testing"

	"github.com/apm-dev/eth_getBalance-proxy/src/config"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain"
	"github.com/apm-dev/eth_getBalance-proxy/src/domain/mocks"
	"github.com/apm-dev/eth_getBalance-proxy/src/proxy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_SendRequest(t *testing.T) {
	config := config.NewConfig()
	mockCache := new(mocks.RpcProxyCache)
	mockNodeRepo := new(mocks.NodeRepository)
	rpcProxy := proxy.NewRpcProxyService(config, mockCache, mockNodeRepo)

	normalRpcRequest, _ := domain.NewJsonRpcRequest(
		1, "eth_getBalance",
		[]string{"0x974CaA59e49682CdA0AD2bbe82983419A2ECC400", "latest"},
	)
	normalRpcResponse := &domain.JsonRpcResponse{
		ID:      1,
		JsonRpc: "2.0",
		Result:  "0x1e287677609cdef347d",
		Error:   nil,
	}

	t.Run("not supported scheme", func(t *testing.T) {
		resp, err := rpcProxy.SendRequest(context.Background(), "ApmChain", normalRpcRequest)
		assert.Nil(t, resp)
		assert.ErrorIs(t, err, domain.ErrInvalidArgument)
		assert.ErrorContains(t, err, "not supported scheme 'ApmChain'")
	})

	t.Run("invalid method format", func(t *testing.T) {
		req, _ := domain.NewJsonRpcRequest(13, "eth_get_Balance", nil)
		resp, err := rpcProxy.SendRequest(context.Background(), "eth", req)
		assert.NoError(t, err)
		assert.EqualValues(t,
			&domain.JsonRpcResponse{
				ID:      13,
				JsonRpc: "2.0",
				Result:  "",
				Error: &domain.JsonRpcError{
					Code:    -32601,
					Message: "the method 'eth_get_Balance' does not exist/is not available",
				},
			},
			resp,
		)
	})

	t.Run("blacklisted method", func(t *testing.T) {
		req, _ := domain.NewJsonRpcRequest(2, "eth_mining", nil)
		resp, err := rpcProxy.SendRequest(context.Background(), "eth", req)
		assert.NoError(t, err)
		assert.EqualValues(t,
			&domain.JsonRpcResponse{
				ID:      2,
				JsonRpc: "2.0",
				Result:  "",
				Error: &domain.JsonRpcError{
					Code:    -32601,
					Message: "the method 'eth_mining' does not exist/is not available",
				},
			},
			resp,
		)
	})

	t.Run("blacklisted prefix", func(t *testing.T) {
		req, _ := domain.NewJsonRpcRequest(12, "engine_getBalance", nil)
		resp, err := rpcProxy.SendRequest(context.Background(), "eth", req)
		assert.NoError(t, err)
		assert.EqualValues(t,
			&domain.JsonRpcResponse{
				ID:      12,
				JsonRpc: "2.0",
				Result:  "",
				Error: &domain.JsonRpcError{
					Code:    -32601,
					Message: "the method 'engine_getBalance' does not exist/is not available",
				},
			},
			resp,
		)
	})

	t.Run("return response from cache when available", func(t *testing.T) {
		mockCache.On("GetCachedResponse", "eth", normalRpcRequest).
			Return(normalRpcResponse, true).Once()

		resp, err := rpcProxy.SendRequest(context.Background(), "eth", normalRpcRequest)

		assert.NoError(t, err)
		assert.EqualValues(t, normalRpcResponse, resp)

		mockCache.AssertExpectations(t)
		mockNodeRepo.AssertNotCalled(t, "GetNodesByBlockchain", "eth")
	})

	t.Run("no available nodes for blockchain", func(t *testing.T) {
		mockCache.On("GetCachedResponse", mock.Anything, mock.Anything).
			Return(nil, false).Once()
		mockNodeRepo.On("GetNodesByBlockchain", "eth").
			Return(nil, nil).Once()

		resp, err := rpcProxy.SendRequest(context.Background(), "eth", normalRpcRequest)
		assert.Nil(t, resp)
		assert.ErrorIs(t, err, domain.ErrInternalServer)

		mockCache.AssertExpectations(t)
		mockNodeRepo.AssertExpectations(t)
	})

	// above test are just Proof Of Concept, there are many more tests to have
}
