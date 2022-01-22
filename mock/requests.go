// Code generated by MockGen. DO NOT EDIT.
// Source: requests.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"
	requests "requests"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIRequests is a mock of IRequests interface.
type MockIRequests struct {
	ctrl     *gomock.Controller
	recorder *MockIRequestsMockRecorder
}

// MockIRequestsMockRecorder is the mock recorder for MockIRequests.
type MockIRequestsMockRecorder struct {
	mock *MockIRequests
}

// NewMockIRequests creates a new mock instance.
func NewMockIRequests(ctrl *gomock.Controller) *MockIRequests {
	mock := &MockIRequests{ctrl: ctrl}
	mock.recorder = &MockIRequestsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRequests) EXPECT() *MockIRequestsMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIRequests) Delete(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIRequestsMockRecorder) Delete(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRequests)(nil).Delete), varargs...)
}

// Do mocks base method.
func (m *MockIRequests) Do(method, toUrl string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{method, toUrl}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockIRequestsMockRecorder) Do(method, toUrl interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{method, toUrl}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockIRequests)(nil).Do), varargs...)
}

// Get mocks base method.
func (m *MockIRequests) Get(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIRequestsMockRecorder) Get(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRequests)(nil).Get), varargs...)
}

// Head mocks base method.
func (m *MockIRequests) Head(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Head", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockIRequestsMockRecorder) Head(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockIRequests)(nil).Head), varargs...)
}

// Options mocks base method.
func (m *MockIRequests) Options(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Options", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Options indicates an expected call of Options.
func (mr *MockIRequestsMockRecorder) Options(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockIRequests)(nil).Options), varargs...)
}

// Patch mocks base method.
func (m *MockIRequests) Patch(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockIRequestsMockRecorder) Patch(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockIRequests)(nil).Patch), varargs...)
}

// Post mocks base method.
func (m *MockIRequests) Post(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Post", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockIRequestsMockRecorder) Post(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockIRequests)(nil).Post), varargs...)
}

// Put mocks base method.
func (m *MockIRequests) Put(url string, opts ...requests.Option) (*requests.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{url}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*requests.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockIRequestsMockRecorder) Put(url interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{url}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIRequests)(nil).Put), varargs...)
}

// SetClient mocks base method.
func (m *MockIRequests) SetClient(client *http.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", client)
}

// SetClient indicates an expected call of SetClient.
func (mr *MockIRequestsMockRecorder) SetClient(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockIRequests)(nil).SetClient), client)
}

// SetProxyUrl mocks base method.
func (m *MockIRequests) SetProxyUrl(proxyUrl string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProxyUrl", proxyUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProxyUrl indicates an expected call of SetProxyUrl.
func (mr *MockIRequestsMockRecorder) SetProxyUrl(proxyUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProxyUrl", reflect.TypeOf((*MockIRequests)(nil).SetProxyUrl), proxyUrl)
}

// SetTimeout mocks base method.
func (m *MockIRequests) SetTimeout(t time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeout", t)
}

// SetTimeout indicates an expected call of SetTimeout.
func (mr *MockIRequestsMockRecorder) SetTimeout(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockIRequests)(nil).SetTimeout), t)
}
