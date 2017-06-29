// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/go.uber.org/thriftrw/protocol/protocol.go

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

package thrift

import (
	gomock "github.com/golang/mock/gomock"
	wire "go.uber.org/thriftrw/wire"
	io "io"
)

// MockProtocol is a mock of Protocol interface
type MockProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolMockRecorder
}

// MockProtocolMockRecorder is the mock recorder for MockProtocol
type MockProtocolMockRecorder struct {
	mock *MockProtocol
}

// NewMockProtocol creates a new mock instance
func NewMockProtocol(ctrl *gomock.Controller) *MockProtocol {
	mock := &MockProtocol{ctrl: ctrl}
	mock.recorder = &MockProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockProtocol) EXPECT() *MockProtocolMockRecorder {
	return _m.recorder
}

// Encode mocks base method
func (_m *MockProtocol) Encode(v wire.Value, w io.Writer) error {
	ret := _m.ctrl.Call(_m, "Encode", v, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode
func (_mr *MockProtocolMockRecorder) Encode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Encode", arg0, arg1)
}

// EncodeEnveloped mocks base method
func (_m *MockProtocol) EncodeEnveloped(e wire.Envelope, w io.Writer) error {
	ret := _m.ctrl.Call(_m, "EncodeEnveloped", e, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncodeEnveloped indicates an expected call of EncodeEnveloped
func (_mr *MockProtocolMockRecorder) EncodeEnveloped(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EncodeEnveloped", arg0, arg1)
}

// Decode mocks base method
func (_m *MockProtocol) Decode(r io.ReaderAt, t wire.Type) (wire.Value, error) {
	ret := _m.ctrl.Call(_m, "Decode", r, t)
	ret0, _ := ret[0].(wire.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (_mr *MockProtocolMockRecorder) Decode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Decode", arg0, arg1)
}

// DecodeEnveloped mocks base method
func (_m *MockProtocol) DecodeEnveloped(r io.ReaderAt) (wire.Envelope, error) {
	ret := _m.ctrl.Call(_m, "DecodeEnveloped", r)
	ret0, _ := ret[0].(wire.Envelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeEnveloped indicates an expected call of DecodeEnveloped
func (_mr *MockProtocolMockRecorder) DecodeEnveloped(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeEnveloped", arg0)
}
