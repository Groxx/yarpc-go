// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/transport (interfaces: ClientConfig,ClientConfigProvider)

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package transporttest

import (
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/api/transport"
)

// MockClientConfig is a mock of ClientConfig interface
type MockClientConfig struct {
	ctrl     *gomock.Controller
	recorder *MockClientConfigMockRecorder
}

// MockClientConfigMockRecorder is the mock recorder for MockClientConfig
type MockClientConfigMockRecorder struct {
	mock *MockClientConfig
}

// NewMockClientConfig creates a new mock instance
func NewMockClientConfig(ctrl *gomock.Controller) *MockClientConfig {
	mock := &MockClientConfig{ctrl: ctrl}
	mock.recorder = &MockClientConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockClientConfig) EXPECT() *MockClientConfigMockRecorder {
	return _m.recorder
}

// Caller mocks base method
func (_m *MockClientConfig) Caller() string {
	ret := _m.ctrl.Call(_m, "Caller")
	ret0, _ := ret[0].(string)
	return ret0
}

// Caller indicates an expected call of Caller
func (_mr *MockClientConfigMockRecorder) Caller() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Caller")
}

// GetOnewayOutbound mocks base method
func (_m *MockClientConfig) GetOnewayOutbound() transport.OnewayOutbound {
	ret := _m.ctrl.Call(_m, "GetOnewayOutbound")
	ret0, _ := ret[0].(transport.OnewayOutbound)
	return ret0
}

// GetOnewayOutbound indicates an expected call of GetOnewayOutbound
func (_mr *MockClientConfigMockRecorder) GetOnewayOutbound() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOnewayOutbound")
}

// GetUnaryOutbound mocks base method
func (_m *MockClientConfig) GetUnaryOutbound() transport.UnaryOutbound {
	ret := _m.ctrl.Call(_m, "GetUnaryOutbound")
	ret0, _ := ret[0].(transport.UnaryOutbound)
	return ret0
}

// GetUnaryOutbound indicates an expected call of GetUnaryOutbound
func (_mr *MockClientConfigMockRecorder) GetUnaryOutbound() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUnaryOutbound")
}

// Service mocks base method
func (_m *MockClientConfig) Service() string {
	ret := _m.ctrl.Call(_m, "Service")
	ret0, _ := ret[0].(string)
	return ret0
}

// Service indicates an expected call of Service
func (_mr *MockClientConfigMockRecorder) Service() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Service")
}

// MockClientConfigProvider is a mock of ClientConfigProvider interface
type MockClientConfigProvider struct {
	ctrl     *gomock.Controller
	recorder *MockClientConfigProviderMockRecorder
}

// MockClientConfigProviderMockRecorder is the mock recorder for MockClientConfigProvider
type MockClientConfigProviderMockRecorder struct {
	mock *MockClientConfigProvider
}

// NewMockClientConfigProvider creates a new mock instance
func NewMockClientConfigProvider(ctrl *gomock.Controller) *MockClientConfigProvider {
	mock := &MockClientConfigProvider{ctrl: ctrl}
	mock.recorder = &MockClientConfigProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockClientConfigProvider) EXPECT() *MockClientConfigProviderMockRecorder {
	return _m.recorder
}

// ClientConfig mocks base method
func (_m *MockClientConfigProvider) ClientConfig(_param0 string) transport.ClientConfig {
	ret := _m.ctrl.Call(_m, "ClientConfig", _param0)
	ret0, _ := ret[0].(transport.ClientConfig)
	return ret0
}

// ClientConfig indicates an expected call of ClientConfig
func (_mr *MockClientConfigProviderMockRecorder) ClientConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClientConfig", arg0)
}
