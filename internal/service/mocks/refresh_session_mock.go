// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nekruzrabiev/simple-app/internal/service (interfaces: RefreshSession)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/nekruzrabiev/simple-app/internal/service"
)

// MockRefreshSession is a mock of RefreshSession interface.
type MockRefreshSession struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshSessionMockRecorder
}

// MockRefreshSessionMockRecorder is the mock recorder for MockRefreshSession.
type MockRefreshSessionMockRecorder struct {
	mock *MockRefreshSession
}

// NewMockRefreshSession creates a new mock instance.
func NewMockRefreshSession(ctrl *gomock.Controller) *MockRefreshSession {
	mock := &MockRefreshSession{ctrl: ctrl}
	mock.recorder = &MockRefreshSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshSession) EXPECT() *MockRefreshSessionMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRefreshSession) Create(arg0 context.Context, arg1 int) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRefreshSessionMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRefreshSession)(nil).Create), arg0, arg1)
}

// Update mocks base method.
func (m *MockRefreshSession) Update(arg0 context.Context, arg1 string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRefreshSessionMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRefreshSession)(nil).Update), arg0, arg1)
}