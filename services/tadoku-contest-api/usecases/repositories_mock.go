// Code generated by MockGen. DO NOT EDIT.
// Source: repositories.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/api/domain"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockUserRepository) Store(user domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockUserRepositoryMockRecorder) Store(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUserRepository)(nil).Store), user)
}

// FindByID mocks base method
func (m *MockUserRepository) FindByID(id uint64) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockUserRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), id)
}

// FindByEmail mocks base method
func (m *MockUserRepository) FindByEmail(email string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// MockContestRepository is a mock of ContestRepository interface
type MockContestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockContestRepositoryMockRecorder
}

// MockContestRepositoryMockRecorder is the mock recorder for MockContestRepository
type MockContestRepositoryMockRecorder struct {
	mock *MockContestRepository
}

// NewMockContestRepository creates a new mock instance
func NewMockContestRepository(ctrl *gomock.Controller) *MockContestRepository {
	mock := &MockContestRepository{ctrl: ctrl}
	mock.recorder = &MockContestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContestRepository) EXPECT() *MockContestRepositoryMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockContestRepository) Store(contest domain.Contest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", contest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockContestRepositoryMockRecorder) Store(contest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockContestRepository)(nil).Store), contest)
}

// GetOpenContests mocks base method
func (m *MockContestRepository) GetOpenContests() ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenContests")
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenContests indicates an expected call of GetOpenContests
func (mr *MockContestRepositoryMockRecorder) GetOpenContests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenContests", reflect.TypeOf((*MockContestRepository)(nil).GetOpenContests))
}
