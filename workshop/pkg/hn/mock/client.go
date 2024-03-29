// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	hn "workshop-starter/pkg/hn"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// MaxItem mocks base method
func (m *MockClient) MaxItem() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxItem")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxItem indicates an expected call of MaxItem
func (mr *MockClientMockRecorder) MaxItem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxItem", reflect.TypeOf((*MockClient)(nil).MaxItem))
}

// GetItem mocks base method
func (m *MockClient) GetItem(itemID int) (hn.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", itemID)
	ret0, _ := ret[0].(hn.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem
func (mr *MockClientMockRecorder) GetItem(itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockClient)(nil).GetItem), itemID)
}
