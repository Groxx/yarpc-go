// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/v2 (interfaces: Chooser,Dialer,Identifier,InboundCodec,List,Peer,Router,RouterMiddleware,Stream,StreamCloser,StreamEncodingHandler,StreamInboundTransportMiddleware,StreamOutbound,StreamOutboundTransportMiddleware,StreamTransportHandler,Subscriber,UnaryEncodingHandler,UnaryInboundEncodingMiddleware,UnaryInboundTransportMiddleware,UnaryOutbound,UnaryOutboundTransportMiddleware,UnaryTransportHandler)

// Copyright (c) 2018 Uber Technologies, Inc.
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

// Package yarpctest is a generated GoMock package.
package yarpctest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v2 "go.uber.org/yarpc/v2"
	reflect "reflect"
)

// MockChooser is a mock of Chooser interface
type MockChooser struct {
	ctrl     *gomock.Controller
	recorder *MockChooserMockRecorder
}

// MockChooserMockRecorder is the mock recorder for MockChooser
type MockChooserMockRecorder struct {
	mock *MockChooser
}

// NewMockChooser creates a new mock instance
func NewMockChooser(ctrl *gomock.Controller) *MockChooser {
	mock := &MockChooser{ctrl: ctrl}
	mock.recorder = &MockChooserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChooser) EXPECT() *MockChooserMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockChooser) Choose(arg0 context.Context, arg1 *v2.Request) (v2.Peer, func(error), error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(v2.Peer)
	ret1, _ := ret[1].(func(error))
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Choose indicates an expected call of Choose
func (mr *MockChooserMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockChooser)(nil).Choose), arg0, arg1)
}

// Name mocks base method
func (m *MockChooser) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockChooserMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockChooser)(nil).Name))
}

// MockDialer is a mock of Dialer interface
type MockDialer struct {
	ctrl     *gomock.Controller
	recorder *MockDialerMockRecorder
}

// MockDialerMockRecorder is the mock recorder for MockDialer
type MockDialerMockRecorder struct {
	mock *MockDialer
}

// NewMockDialer creates a new mock instance
func NewMockDialer(ctrl *gomock.Controller) *MockDialer {
	mock := &MockDialer{ctrl: ctrl}
	mock.recorder = &MockDialerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDialer) EXPECT() *MockDialerMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockDialer) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockDialerMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDialer)(nil).Name))
}

// ReleasePeer mocks base method
func (m *MockDialer) ReleasePeer(arg0 v2.Identifier, arg1 v2.Subscriber) error {
	ret := m.ctrl.Call(m, "ReleasePeer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleasePeer indicates an expected call of ReleasePeer
func (mr *MockDialerMockRecorder) ReleasePeer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleasePeer", reflect.TypeOf((*MockDialer)(nil).ReleasePeer), arg0, arg1)
}

// RetainPeer mocks base method
func (m *MockDialer) RetainPeer(arg0 v2.Identifier, arg1 v2.Subscriber) (v2.Peer, error) {
	ret := m.ctrl.Call(m, "RetainPeer", arg0, arg1)
	ret0, _ := ret[0].(v2.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetainPeer indicates an expected call of RetainPeer
func (mr *MockDialerMockRecorder) RetainPeer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetainPeer", reflect.TypeOf((*MockDialer)(nil).RetainPeer), arg0, arg1)
}

// MockIdentifier is a mock of Identifier interface
type MockIdentifier struct {
	ctrl     *gomock.Controller
	recorder *MockIdentifierMockRecorder
}

// MockIdentifierMockRecorder is the mock recorder for MockIdentifier
type MockIdentifierMockRecorder struct {
	mock *MockIdentifier
}

// NewMockIdentifier creates a new mock instance
func NewMockIdentifier(ctrl *gomock.Controller) *MockIdentifier {
	mock := &MockIdentifier{ctrl: ctrl}
	mock.recorder = &MockIdentifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentifier) EXPECT() *MockIdentifierMockRecorder {
	return m.recorder
}

// Identifier mocks base method
func (m *MockIdentifier) Identifier() string {
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identifier indicates an expected call of Identifier
func (mr *MockIdentifierMockRecorder) Identifier() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockIdentifier)(nil).Identifier))
}

// MockInboundCodec is a mock of InboundCodec interface
type MockInboundCodec struct {
	ctrl     *gomock.Controller
	recorder *MockInboundCodecMockRecorder
}

// MockInboundCodecMockRecorder is the mock recorder for MockInboundCodec
type MockInboundCodecMockRecorder struct {
	mock *MockInboundCodec
}

// NewMockInboundCodec creates a new mock instance
func NewMockInboundCodec(ctrl *gomock.Controller) *MockInboundCodec {
	mock := &MockInboundCodec{ctrl: ctrl}
	mock.recorder = &MockInboundCodecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInboundCodec) EXPECT() *MockInboundCodecMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockInboundCodec) Decode(arg0 *v2.Buffer) (interface{}, error) {
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockInboundCodecMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockInboundCodec)(nil).Decode), arg0)
}

// Encode mocks base method
func (m *MockInboundCodec) Encode(arg0 interface{}) (*v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Encode", arg0)
	ret0, _ := ret[0].(*v2.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockInboundCodecMockRecorder) Encode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockInboundCodec)(nil).Encode), arg0)
}

// MockList is a mock of List interface
type MockList struct {
	ctrl     *gomock.Controller
	recorder *MockListMockRecorder
}

// MockListMockRecorder is the mock recorder for MockList
type MockListMockRecorder struct {
	mock *MockList
}

// NewMockList creates a new mock instance
func NewMockList(ctrl *gomock.Controller) *MockList {
	mock := &MockList{ctrl: ctrl}
	mock.recorder = &MockListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockList) EXPECT() *MockListMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockList) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockListMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockList)(nil).Name))
}

// Update mocks base method
func (m *MockList) Update(arg0 v2.ListUpdates) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockListMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockList)(nil).Update), arg0)
}

// MockPeer is a mock of Peer interface
type MockPeer struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMockRecorder
}

// MockPeerMockRecorder is the mock recorder for MockPeer
type MockPeerMockRecorder struct {
	mock *MockPeer
}

// NewMockPeer creates a new mock instance
func NewMockPeer(ctrl *gomock.Controller) *MockPeer {
	mock := &MockPeer{ctrl: ctrl}
	mock.recorder = &MockPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeer) EXPECT() *MockPeerMockRecorder {
	return m.recorder
}

// EndRequest mocks base method
func (m *MockPeer) EndRequest() {
	m.ctrl.Call(m, "EndRequest")
}

// EndRequest indicates an expected call of EndRequest
func (mr *MockPeerMockRecorder) EndRequest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndRequest", reflect.TypeOf((*MockPeer)(nil).EndRequest))
}

// Identifier mocks base method
func (m *MockPeer) Identifier() string {
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identifier indicates an expected call of Identifier
func (mr *MockPeerMockRecorder) Identifier() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockPeer)(nil).Identifier))
}

// StartRequest mocks base method
func (m *MockPeer) StartRequest() {
	m.ctrl.Call(m, "StartRequest")
}

// StartRequest indicates an expected call of StartRequest
func (mr *MockPeerMockRecorder) StartRequest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRequest", reflect.TypeOf((*MockPeer)(nil).StartRequest))
}

// Status mocks base method
func (m *MockPeer) Status() v2.Status {
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(v2.Status)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockPeerMockRecorder) Status() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockPeer)(nil).Status))
}

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouter) Choose(arg0 context.Context, arg1 *v2.Request) (v2.TransportHandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(v2.TransportHandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouterMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouter)(nil).Choose), arg0, arg1)
}

// Procedures mocks base method
func (m *MockRouter) Procedures() []v2.TransportProcedure {
	ret := m.ctrl.Call(m, "Procedures")
	ret0, _ := ret[0].([]v2.TransportProcedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouterMockRecorder) Procedures() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouter)(nil).Procedures))
}

// MockRouterMiddleware is a mock of RouterMiddleware interface
type MockRouterMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMiddlewareMockRecorder
}

// MockRouterMiddlewareMockRecorder is the mock recorder for MockRouterMiddleware
type MockRouterMiddlewareMockRecorder struct {
	mock *MockRouterMiddleware
}

// NewMockRouterMiddleware creates a new mock instance
func NewMockRouterMiddleware(ctrl *gomock.Controller) *MockRouterMiddleware {
	mock := &MockRouterMiddleware{ctrl: ctrl}
	mock.recorder = &MockRouterMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouterMiddleware) EXPECT() *MockRouterMiddlewareMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouterMiddleware) Choose(arg0 context.Context, arg1 *v2.Request, arg2 v2.Router) (v2.TransportHandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1, arg2)
	ret0, _ := ret[0].(v2.TransportHandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouterMiddlewareMockRecorder) Choose(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouterMiddleware)(nil).Choose), arg0, arg1, arg2)
}

// Procedures mocks base method
func (m *MockRouterMiddleware) Procedures(arg0 v2.Router) []v2.TransportProcedure {
	ret := m.ctrl.Call(m, "Procedures", arg0)
	ret0, _ := ret[0].([]v2.TransportProcedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouterMiddlewareMockRecorder) Procedures(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouterMiddleware)(nil).Procedures), arg0)
}

// MockStream is a mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockStream) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockStreamMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStream)(nil).Context))
}

// ReceiveMessage mocks base method
func (m *MockStream) ReceiveMessage(arg0 context.Context) (*v2.StreamMessage, error) {
	ret := m.ctrl.Call(m, "ReceiveMessage", arg0)
	ret0, _ := ret[0].(*v2.StreamMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage
func (mr *MockStreamMockRecorder) ReceiveMessage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockStream)(nil).ReceiveMessage), arg0)
}

// Request mocks base method
func (m *MockStream) Request() *v2.Request {
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*v2.Request)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockStreamMockRecorder) Request() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockStream)(nil).Request))
}

// SendMessage mocks base method
func (m *MockStream) SendMessage(arg0 context.Context, arg1 *v2.StreamMessage) error {
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockStreamMockRecorder) SendMessage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockStream)(nil).SendMessage), arg0, arg1)
}

// MockStreamCloser is a mock of StreamCloser interface
type MockStreamCloser struct {
	ctrl     *gomock.Controller
	recorder *MockStreamCloserMockRecorder
}

// MockStreamCloserMockRecorder is the mock recorder for MockStreamCloser
type MockStreamCloserMockRecorder struct {
	mock *MockStreamCloser
}

// NewMockStreamCloser creates a new mock instance
func NewMockStreamCloser(ctrl *gomock.Controller) *MockStreamCloser {
	mock := &MockStreamCloser{ctrl: ctrl}
	mock.recorder = &MockStreamCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamCloser) EXPECT() *MockStreamCloserMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStreamCloser) Close(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStreamCloserMockRecorder) Close(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStreamCloser)(nil).Close), arg0)
}

// Context mocks base method
func (m *MockStreamCloser) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockStreamCloserMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStreamCloser)(nil).Context))
}

// ReceiveMessage mocks base method
func (m *MockStreamCloser) ReceiveMessage(arg0 context.Context) (*v2.StreamMessage, error) {
	ret := m.ctrl.Call(m, "ReceiveMessage", arg0)
	ret0, _ := ret[0].(*v2.StreamMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage
func (mr *MockStreamCloserMockRecorder) ReceiveMessage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockStreamCloser)(nil).ReceiveMessage), arg0)
}

// Request mocks base method
func (m *MockStreamCloser) Request() *v2.Request {
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*v2.Request)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockStreamCloserMockRecorder) Request() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockStreamCloser)(nil).Request))
}

// SendMessage mocks base method
func (m *MockStreamCloser) SendMessage(arg0 context.Context, arg1 *v2.StreamMessage) error {
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockStreamCloserMockRecorder) SendMessage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockStreamCloser)(nil).SendMessage), arg0, arg1)
}

// MockStreamEncodingHandler is a mock of StreamEncodingHandler interface
type MockStreamEncodingHandler struct {
	ctrl     *gomock.Controller
	recorder *MockStreamEncodingHandlerMockRecorder
}

// MockStreamEncodingHandlerMockRecorder is the mock recorder for MockStreamEncodingHandler
type MockStreamEncodingHandlerMockRecorder struct {
	mock *MockStreamEncodingHandler
}

// NewMockStreamEncodingHandler creates a new mock instance
func NewMockStreamEncodingHandler(ctrl *gomock.Controller) *MockStreamEncodingHandler {
	mock := &MockStreamEncodingHandler{ctrl: ctrl}
	mock.recorder = &MockStreamEncodingHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamEncodingHandler) EXPECT() *MockStreamEncodingHandlerMockRecorder {
	return m.recorder
}

// HandleStream mocks base method
func (m *MockStreamEncodingHandler) HandleStream(arg0 *v2.ServerStream) error {
	ret := m.ctrl.Call(m, "HandleStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStream indicates an expected call of HandleStream
func (mr *MockStreamEncodingHandlerMockRecorder) HandleStream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStream", reflect.TypeOf((*MockStreamEncodingHandler)(nil).HandleStream), arg0)
}

// MockStreamInboundTransportMiddleware is a mock of StreamInboundTransportMiddleware interface
type MockStreamInboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockStreamInboundTransportMiddlewareMockRecorder
}

// MockStreamInboundTransportMiddlewareMockRecorder is the mock recorder for MockStreamInboundTransportMiddleware
type MockStreamInboundTransportMiddlewareMockRecorder struct {
	mock *MockStreamInboundTransportMiddleware
}

// NewMockStreamInboundTransportMiddleware creates a new mock instance
func NewMockStreamInboundTransportMiddleware(ctrl *gomock.Controller) *MockStreamInboundTransportMiddleware {
	mock := &MockStreamInboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockStreamInboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamInboundTransportMiddleware) EXPECT() *MockStreamInboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// HandleStream mocks base method
func (m *MockStreamInboundTransportMiddleware) HandleStream(arg0 *v2.ServerStream, arg1 v2.StreamTransportHandler) error {
	ret := m.ctrl.Call(m, "HandleStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStream indicates an expected call of HandleStream
func (mr *MockStreamInboundTransportMiddlewareMockRecorder) HandleStream(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStream", reflect.TypeOf((*MockStreamInboundTransportMiddleware)(nil).HandleStream), arg0, arg1)
}

// MockStreamOutbound is a mock of StreamOutbound interface
type MockStreamOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockStreamOutboundMockRecorder
}

// MockStreamOutboundMockRecorder is the mock recorder for MockStreamOutbound
type MockStreamOutboundMockRecorder struct {
	mock *MockStreamOutbound
}

// NewMockStreamOutbound creates a new mock instance
func NewMockStreamOutbound(ctrl *gomock.Controller) *MockStreamOutbound {
	mock := &MockStreamOutbound{ctrl: ctrl}
	mock.recorder = &MockStreamOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamOutbound) EXPECT() *MockStreamOutboundMockRecorder {
	return m.recorder
}

// CallStream mocks base method
func (m *MockStreamOutbound) CallStream(arg0 context.Context, arg1 *v2.Request) (*v2.ClientStream, error) {
	ret := m.ctrl.Call(m, "CallStream", arg0, arg1)
	ret0, _ := ret[0].(*v2.ClientStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallStream indicates an expected call of CallStream
func (mr *MockStreamOutboundMockRecorder) CallStream(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallStream", reflect.TypeOf((*MockStreamOutbound)(nil).CallStream), arg0, arg1)
}

// MockStreamOutboundTransportMiddleware is a mock of StreamOutboundTransportMiddleware interface
type MockStreamOutboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockStreamOutboundTransportMiddlewareMockRecorder
}

// MockStreamOutboundTransportMiddlewareMockRecorder is the mock recorder for MockStreamOutboundTransportMiddleware
type MockStreamOutboundTransportMiddlewareMockRecorder struct {
	mock *MockStreamOutboundTransportMiddleware
}

// NewMockStreamOutboundTransportMiddleware creates a new mock instance
func NewMockStreamOutboundTransportMiddleware(ctrl *gomock.Controller) *MockStreamOutboundTransportMiddleware {
	mock := &MockStreamOutboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockStreamOutboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamOutboundTransportMiddleware) EXPECT() *MockStreamOutboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// CallStream mocks base method
func (m *MockStreamOutboundTransportMiddleware) CallStream(arg0 context.Context, arg1 *v2.Request, arg2 v2.StreamOutbound) (*v2.ClientStream, error) {
	ret := m.ctrl.Call(m, "CallStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v2.ClientStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallStream indicates an expected call of CallStream
func (mr *MockStreamOutboundTransportMiddlewareMockRecorder) CallStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallStream", reflect.TypeOf((*MockStreamOutboundTransportMiddleware)(nil).CallStream), arg0, arg1, arg2)
}

// MockStreamTransportHandler is a mock of StreamTransportHandler interface
type MockStreamTransportHandler struct {
	ctrl     *gomock.Controller
	recorder *MockStreamTransportHandlerMockRecorder
}

// MockStreamTransportHandlerMockRecorder is the mock recorder for MockStreamTransportHandler
type MockStreamTransportHandlerMockRecorder struct {
	mock *MockStreamTransportHandler
}

// NewMockStreamTransportHandler creates a new mock instance
func NewMockStreamTransportHandler(ctrl *gomock.Controller) *MockStreamTransportHandler {
	mock := &MockStreamTransportHandler{ctrl: ctrl}
	mock.recorder = &MockStreamTransportHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamTransportHandler) EXPECT() *MockStreamTransportHandlerMockRecorder {
	return m.recorder
}

// HandleStream mocks base method
func (m *MockStreamTransportHandler) HandleStream(arg0 *v2.ServerStream) error {
	ret := m.ctrl.Call(m, "HandleStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStream indicates an expected call of HandleStream
func (mr *MockStreamTransportHandlerMockRecorder) HandleStream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStream", reflect.TypeOf((*MockStreamTransportHandler)(nil).HandleStream), arg0)
}

// MockSubscriber is a mock of Subscriber interface
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// NotifyStatusChanged mocks base method
func (m *MockSubscriber) NotifyStatusChanged(arg0 v2.Identifier) {
	m.ctrl.Call(m, "NotifyStatusChanged", arg0)
}

// NotifyStatusChanged indicates an expected call of NotifyStatusChanged
func (mr *MockSubscriberMockRecorder) NotifyStatusChanged(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyStatusChanged", reflect.TypeOf((*MockSubscriber)(nil).NotifyStatusChanged), arg0)
}

// MockUnaryEncodingHandler is a mock of UnaryEncodingHandler interface
type MockUnaryEncodingHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryEncodingHandlerMockRecorder
}

// MockUnaryEncodingHandlerMockRecorder is the mock recorder for MockUnaryEncodingHandler
type MockUnaryEncodingHandlerMockRecorder struct {
	mock *MockUnaryEncodingHandler
}

// NewMockUnaryEncodingHandler creates a new mock instance
func NewMockUnaryEncodingHandler(ctrl *gomock.Controller) *MockUnaryEncodingHandler {
	mock := &MockUnaryEncodingHandler{ctrl: ctrl}
	mock.recorder = &MockUnaryEncodingHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryEncodingHandler) EXPECT() *MockUnaryEncodingHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockUnaryEncodingHandler) Handle(arg0 context.Context, arg1 interface{}) (interface{}, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle
func (mr *MockUnaryEncodingHandlerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUnaryEncodingHandler)(nil).Handle), arg0, arg1)
}

// MockUnaryInboundEncodingMiddleware is a mock of UnaryInboundEncodingMiddleware interface
type MockUnaryInboundEncodingMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryInboundEncodingMiddlewareMockRecorder
}

// MockUnaryInboundEncodingMiddlewareMockRecorder is the mock recorder for MockUnaryInboundEncodingMiddleware
type MockUnaryInboundEncodingMiddlewareMockRecorder struct {
	mock *MockUnaryInboundEncodingMiddleware
}

// NewMockUnaryInboundEncodingMiddleware creates a new mock instance
func NewMockUnaryInboundEncodingMiddleware(ctrl *gomock.Controller) *MockUnaryInboundEncodingMiddleware {
	mock := &MockUnaryInboundEncodingMiddleware{ctrl: ctrl}
	mock.recorder = &MockUnaryInboundEncodingMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryInboundEncodingMiddleware) EXPECT() *MockUnaryInboundEncodingMiddlewareMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockUnaryInboundEncodingMiddleware) Handle(arg0 context.Context, arg1 interface{}, arg2 v2.UnaryEncodingHandler) (interface{}, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle
func (mr *MockUnaryInboundEncodingMiddlewareMockRecorder) Handle(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUnaryInboundEncodingMiddleware)(nil).Handle), arg0, arg1, arg2)
}

// MockUnaryInboundTransportMiddleware is a mock of UnaryInboundTransportMiddleware interface
type MockUnaryInboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryInboundTransportMiddlewareMockRecorder
}

// MockUnaryInboundTransportMiddlewareMockRecorder is the mock recorder for MockUnaryInboundTransportMiddleware
type MockUnaryInboundTransportMiddlewareMockRecorder struct {
	mock *MockUnaryInboundTransportMiddleware
}

// NewMockUnaryInboundTransportMiddleware creates a new mock instance
func NewMockUnaryInboundTransportMiddleware(ctrl *gomock.Controller) *MockUnaryInboundTransportMiddleware {
	mock := &MockUnaryInboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockUnaryInboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryInboundTransportMiddleware) EXPECT() *MockUnaryInboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockUnaryInboundTransportMiddleware) Handle(arg0 context.Context, arg1 *v2.Request, arg2 *v2.Buffer, arg3 v2.UnaryTransportHandler) (*v2.Response, *v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v2.Response)
	ret1, _ := ret[1].(*v2.Buffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Handle indicates an expected call of Handle
func (mr *MockUnaryInboundTransportMiddlewareMockRecorder) Handle(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUnaryInboundTransportMiddleware)(nil).Handle), arg0, arg1, arg2, arg3)
}

// Name mocks base method
func (m *MockUnaryInboundTransportMiddleware) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockUnaryInboundTransportMiddlewareMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnaryInboundTransportMiddleware)(nil).Name))
}

// MockUnaryOutbound is a mock of UnaryOutbound interface
type MockUnaryOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryOutboundMockRecorder
}

// MockUnaryOutboundMockRecorder is the mock recorder for MockUnaryOutbound
type MockUnaryOutboundMockRecorder struct {
	mock *MockUnaryOutbound
}

// NewMockUnaryOutbound creates a new mock instance
func NewMockUnaryOutbound(ctrl *gomock.Controller) *MockUnaryOutbound {
	mock := &MockUnaryOutbound{ctrl: ctrl}
	mock.recorder = &MockUnaryOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryOutbound) EXPECT() *MockUnaryOutboundMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockUnaryOutbound) Call(arg0 context.Context, arg1 *v2.Request, arg2 *v2.Buffer) (*v2.Response, *v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v2.Response)
	ret1, _ := ret[1].(*v2.Buffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Call indicates an expected call of Call
func (mr *MockUnaryOutboundMockRecorder) Call(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockUnaryOutbound)(nil).Call), arg0, arg1, arg2)
}

// MockUnaryOutboundTransportMiddleware is a mock of UnaryOutboundTransportMiddleware interface
type MockUnaryOutboundTransportMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryOutboundTransportMiddlewareMockRecorder
}

// MockUnaryOutboundTransportMiddlewareMockRecorder is the mock recorder for MockUnaryOutboundTransportMiddleware
type MockUnaryOutboundTransportMiddlewareMockRecorder struct {
	mock *MockUnaryOutboundTransportMiddleware
}

// NewMockUnaryOutboundTransportMiddleware creates a new mock instance
func NewMockUnaryOutboundTransportMiddleware(ctrl *gomock.Controller) *MockUnaryOutboundTransportMiddleware {
	mock := &MockUnaryOutboundTransportMiddleware{ctrl: ctrl}
	mock.recorder = &MockUnaryOutboundTransportMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryOutboundTransportMiddleware) EXPECT() *MockUnaryOutboundTransportMiddlewareMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockUnaryOutboundTransportMiddleware) Call(arg0 context.Context, arg1 *v2.Request, arg2 *v2.Buffer, arg3 v2.UnaryOutbound) (*v2.Response, *v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v2.Response)
	ret1, _ := ret[1].(*v2.Buffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Call indicates an expected call of Call
func (mr *MockUnaryOutboundTransportMiddlewareMockRecorder) Call(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockUnaryOutboundTransportMiddleware)(nil).Call), arg0, arg1, arg2, arg3)
}

// Name mocks base method
func (m *MockUnaryOutboundTransportMiddleware) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockUnaryOutboundTransportMiddlewareMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnaryOutboundTransportMiddleware)(nil).Name))
}

// MockUnaryTransportHandler is a mock of UnaryTransportHandler interface
type MockUnaryTransportHandler struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryTransportHandlerMockRecorder
}

// MockUnaryTransportHandlerMockRecorder is the mock recorder for MockUnaryTransportHandler
type MockUnaryTransportHandlerMockRecorder struct {
	mock *MockUnaryTransportHandler
}

// NewMockUnaryTransportHandler creates a new mock instance
func NewMockUnaryTransportHandler(ctrl *gomock.Controller) *MockUnaryTransportHandler {
	mock := &MockUnaryTransportHandler{ctrl: ctrl}
	mock.recorder = &MockUnaryTransportHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryTransportHandler) EXPECT() *MockUnaryTransportHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockUnaryTransportHandler) Handle(arg0 context.Context, arg1 *v2.Request, arg2 *v2.Buffer) (*v2.Response, *v2.Buffer, error) {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v2.Response)
	ret1, _ := ret[1].(*v2.Buffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Handle indicates an expected call of Handle
func (mr *MockUnaryTransportHandlerMockRecorder) Handle(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUnaryTransportHandler)(nil).Handle), arg0, arg1, arg2)
}
