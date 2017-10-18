// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sclevine/cflocal/cf/cmd (interfaces: LocalApp)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockLocalApp is a mock of LocalApp interface
type MockLocalApp struct {
	ctrl     *gomock.Controller
	recorder *MockLocalAppMockRecorder
}

// MockLocalAppMockRecorder is the mock recorder for MockLocalApp
type MockLocalAppMockRecorder struct {
	mock *MockLocalApp
}

// NewMockLocalApp creates a new mock instance
func NewMockLocalApp(ctrl *gomock.Controller) *MockLocalApp {
	mock := &MockLocalApp{ctrl: ctrl}
	mock.recorder = &MockLocalAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalApp) EXPECT() *MockLocalAppMockRecorder {
	return m.recorder
}

// Tar mocks base method
func (m *MockLocalApp) Tar(arg0 string) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "Tar", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tar indicates an expected call of Tar
func (mr *MockLocalAppMockRecorder) Tar(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tar", reflect.TypeOf((*MockLocalApp)(nil).Tar), arg0)
}