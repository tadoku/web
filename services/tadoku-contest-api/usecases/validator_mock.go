// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// ValidateStruct mocks base method
func (m *MockValidator) ValidateStruct(s interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStruct", s)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateStruct indicates an expected call of ValidateStruct
func (mr *MockValidatorMockRecorder) ValidateStruct(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStruct", reflect.TypeOf((*MockValidator)(nil).ValidateStruct), s)
}
