// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/apm-dev/eth_getBalance-proxy/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// NodeRepository is an autogenerated mock type for the NodeRepository type
type NodeRepository struct {
	mock.Mock
}

// GetNodesByBlockchain provides a mock function with given fields: blockchain
func (_m *NodeRepository) GetNodesByBlockchain(blockchain string) ([]domain.Node, error) {
	ret := _m.Called(blockchain)

	var r0 []domain.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Node, error)); ok {
		return rf(blockchain)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Node); ok {
		r0 = rf(blockchain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(blockchain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewNodeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewNodeRepository creates a new instance of NodeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNodeRepository(t mockConstructorTestingTNewNodeRepository) *NodeRepository {
	mock := &NodeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
