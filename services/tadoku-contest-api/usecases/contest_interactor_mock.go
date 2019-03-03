// Code generated by MockGen. DO NOT EDIT.
// Source: contest_interactor.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/api/domain"
	reflect "reflect"
)

// MockContestInteractor is a mock of ContestInteractor interface
type MockContestInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockContestInteractorMockRecorder
}

// MockContestInteractorMockRecorder is the mock recorder for MockContestInteractor
type MockContestInteractorMockRecorder struct {
	mock *MockContestInteractor
}

// NewMockContestInteractor creates a new mock instance
func NewMockContestInteractor(ctrl *gomock.Controller) *MockContestInteractor {
	mock := &MockContestInteractor{ctrl: ctrl}
	mock.recorder = &MockContestInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContestInteractor) EXPECT() *MockContestInteractorMockRecorder {
	return m.recorder
}

// CreateContest mocks base method
func (m *MockContestInteractor) CreateContest(contest domain.Contest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContest", contest)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateContest indicates an expected call of CreateContest
func (mr *MockContestInteractorMockRecorder) CreateContest(contest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContest", reflect.TypeOf((*MockContestInteractor)(nil).CreateContest), contest)
}

// UpdateContest mocks base method
func (m *MockContestInteractor) UpdateContest(contest domain.Contest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContest", contest)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContest indicates an expected call of UpdateContest
func (mr *MockContestInteractorMockRecorder) UpdateContest(contest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContest", reflect.TypeOf((*MockContestInteractor)(nil).UpdateContest), contest)
}
