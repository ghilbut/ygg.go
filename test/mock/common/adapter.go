// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../../common/adapter.go

package mock_common

import (
	. "github.com/ghilbut/ygg.go/common"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Adapter interface
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *_MockAdapterRecorder
}

// Recorder for MockAdapter (not exported)
type _MockAdapterRecorder struct {
	mock *MockAdapter
}

func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &_MockAdapterRecorder{mock}
	return mock
}

func (_m *MockAdapter) EXPECT() *_MockAdapterRecorder {
	return _m.recorder
}

func (_m *MockAdapter) BindDelegate(delegate AdapterDelegate) {
	_m.ctrl.Call(_m, "BindDelegate", delegate)
}

func (_mr *_MockAdapterRecorder) BindDelegate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BindDelegate", arg0)
}

func (_m *MockAdapter) UnbindDelegate() {
	_m.ctrl.Call(_m, "UnbindDelegate")
}

func (_mr *_MockAdapterRecorder) UnbindDelegate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnbindDelegate")
}

func (_m *MockAdapter) SetCtrlProxy(proxy *CtrlProxy) {
	_m.ctrl.Call(_m, "SetCtrlProxy", proxy)
}

func (_mr *_MockAdapterRecorder) SetCtrlProxy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetCtrlProxy", arg0)
}

func (_m *MockAdapter) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockAdapterRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of AdapterDelegate interface
type MockAdapterDelegate struct {
	ctrl     *gomock.Controller
	recorder *_MockAdapterDelegateRecorder
}

// Recorder for MockAdapterDelegate (not exported)
type _MockAdapterDelegateRecorder struct {
	mock *MockAdapterDelegate
}

func NewMockAdapterDelegate(ctrl *gomock.Controller) *MockAdapterDelegate {
	mock := &MockAdapterDelegate{ctrl: ctrl}
	mock.recorder = &_MockAdapterDelegateRecorder{mock}
	return mock
}

func (_m *MockAdapterDelegate) EXPECT() *_MockAdapterDelegateRecorder {
	return _m.recorder
}

func (_m *MockAdapterDelegate) OnAdapterClosed(adapter Adapter) {
	_m.ctrl.Call(_m, "OnAdapterClosed", adapter)
}

func (_mr *_MockAdapterDelegateRecorder) OnAdapterClosed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnAdapterClosed", arg0)
}
