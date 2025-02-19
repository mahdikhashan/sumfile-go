// Code generated by MockGen. DO NOT EDIT.
// Source: sumfile.go
//
// Generated by this command:
//
//	mockgen -source=sumfile.go -destination=mock_sumfile.go -package=main
//

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockioReader is a mock of ioReader interface.
type MockioReader struct {
	ctrl     *gomock.Controller
	recorder *MockioReaderMockRecorder
	isgomock struct{}
}

// MockioReaderMockRecorder is the mock recorder for MockioReader.
type MockioReaderMockRecorder struct {
	mock *MockioReader
}

// NewMockioReader creates a new mock instance.
func NewMockioReader(ctrl *gomock.Controller) *MockioReader {
	mock := &MockioReader{ctrl: ctrl}
	mock.recorder = &MockioReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockioReader) EXPECT() *MockioReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockioReader) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockioReaderMockRecorder) Read(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockioReader)(nil).Read), p)
}
