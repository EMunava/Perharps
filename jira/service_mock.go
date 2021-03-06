// Code generated by MockGen. DO NOT EDIT.
// Source: ../jira/service.go

// Package jira is a generated GoMock package.
package jira

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateJira mocks base method
func (m *MockService) CreateJira(ctx context.Context, chatId uint32, title, description, name string) {
	m.ctrl.Call(m, "CreateJira", ctx, chatId, title, description, name)
}

// CreateJira indicates an expected call of CreateJira
func (mr *MockServiceMockRecorder) CreateJira(ctx, chatId, title, description, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJira", reflect.TypeOf((*MockService)(nil).CreateJira), ctx, chatId, title, description, name)
}
