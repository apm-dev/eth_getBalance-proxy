// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/apm-dev/eth_getBalance-proxy/src/domain"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// RpcProxyCache is an autogenerated mock type for the RpcProxyCache type
type RpcProxyCache struct {
	mock.Mock
}

// CacheResponse provides a mock function with given fields: blockchain, req, rsp, ttl
func (_m *RpcProxyCache) CacheResponse(blockchain string, req *domain.JsonRpcRequest, rsp *domain.JsonRpcResponse, ttl time.Duration) {
	_m.Called(blockchain, req, rsp, ttl)
}

// GetCachedResponse provides a mock function with given fields: blockchain, req
func (_m *RpcProxyCache) GetCachedResponse(blockchain string, req *domain.JsonRpcRequest) (*domain.JsonRpcResponse, bool) {
	ret := _m.Called(blockchain, req)

	var r0 *domain.JsonRpcResponse
	var r1 bool
	if rf, ok := ret.Get(0).(func(string, *domain.JsonRpcRequest) (*domain.JsonRpcResponse, bool)); ok {
		return rf(blockchain, req)
	}
	if rf, ok := ret.Get(0).(func(string, *domain.JsonRpcRequest) *domain.JsonRpcResponse); ok {
		r0 = rf(blockchain, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.JsonRpcResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *domain.JsonRpcRequest) bool); ok {
		r1 = rf(blockchain, req)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

type mockConstructorTestingTNewRpcProxyCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewRpcProxyCache creates a new instance of RpcProxyCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRpcProxyCache(t mockConstructorTestingTNewRpcProxyCache) *RpcProxyCache {
	mock := &RpcProxyCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
