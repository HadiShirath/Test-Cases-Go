// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpConnector is a mock of HttpConnector interface.
type MockHttpConnector struct {
	ctrl     *gomock.Controller
	recorder *MockHttpConnectorMockRecorder
}

// MockHttpConnectorMockRecorder is the mock recorder for MockHttpConnector.
type MockHttpConnectorMockRecorder struct {
	mock *MockHttpConnector
}

// NewMockHttpConnector creates a new mock instance.
func NewMockHttpConnector(ctrl *gomock.Controller) *MockHttpConnector {
	mock := &MockHttpConnector{ctrl: ctrl}
	mock.recorder = &MockHttpConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpConnector) EXPECT() *MockHttpConnectorMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHttpConnector) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHttpConnectorMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpConnector)(nil).Do), req)
}
