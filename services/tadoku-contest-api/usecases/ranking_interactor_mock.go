// Code generated by MockGen. DO NOT EDIT.
// Source: ranking_interactor.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/api/domain"
	reflect "reflect"
)

// MockRankingInteractor is a mock of RankingInteractor interface
type MockRankingInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockRankingInteractorMockRecorder
}

// MockRankingInteractorMockRecorder is the mock recorder for MockRankingInteractor
type MockRankingInteractorMockRecorder struct {
	mock *MockRankingInteractor
}

// NewMockRankingInteractor creates a new mock instance
func NewMockRankingInteractor(ctrl *gomock.Controller) *MockRankingInteractor {
	mock := &MockRankingInteractor{ctrl: ctrl}
	mock.recorder = &MockRankingInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRankingInteractor) EXPECT() *MockRankingInteractorMockRecorder {
	return m.recorder
}

// CreateRanking mocks base method
func (m *MockRankingInteractor) CreateRanking(contestID, userID uint64, languages domain.LanguageCodes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRanking", contestID, userID, languages)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRanking indicates an expected call of CreateRanking
func (mr *MockRankingInteractorMockRecorder) CreateRanking(contestID, userID, languages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRanking", reflect.TypeOf((*MockRankingInteractor)(nil).CreateRanking), contestID, userID, languages)
}

// CreateLog mocks base method
func (m *MockRankingInteractor) CreateLog(log domain.ContestLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLog", log)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLog indicates an expected call of CreateLog
func (mr *MockRankingInteractorMockRecorder) CreateLog(log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLog", reflect.TypeOf((*MockRankingInteractor)(nil).CreateLog), log)
}

// UpdateLog mocks base method
func (m *MockRankingInteractor) UpdateLog(log domain.ContestLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLog", log)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLog indicates an expected call of UpdateLog
func (mr *MockRankingInteractorMockRecorder) UpdateLog(log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLog", reflect.TypeOf((*MockRankingInteractor)(nil).UpdateLog), log)
}

// DeleteLog mocks base method
func (m *MockRankingInteractor) DeleteLog(logID, userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLog", logID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLog indicates an expected call of DeleteLog
func (mr *MockRankingInteractorMockRecorder) DeleteLog(logID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLog", reflect.TypeOf((*MockRankingInteractor)(nil).DeleteLog), logID, userID)
}

// UpdateRanking mocks base method
func (m *MockRankingInteractor) UpdateRanking(contestID, userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRanking", contestID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRanking indicates an expected call of UpdateRanking
func (mr *MockRankingInteractorMockRecorder) UpdateRanking(contestID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRanking", reflect.TypeOf((*MockRankingInteractor)(nil).UpdateRanking), contestID, userID)
}

// RankingsForRegistration mocks base method
func (m *MockRankingInteractor) RankingsForRegistration(contestID, userID uint64) (domain.Rankings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RankingsForRegistration", contestID, userID)
	ret0, _ := ret[0].(domain.Rankings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RankingsForRegistration indicates an expected call of RankingsForRegistration
func (mr *MockRankingInteractorMockRecorder) RankingsForRegistration(contestID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RankingsForRegistration", reflect.TypeOf((*MockRankingInteractor)(nil).RankingsForRegistration), contestID, userID)
}

// RankingsForContest mocks base method
func (m *MockRankingInteractor) RankingsForContest(contestID uint64, languageCode domain.LanguageCode) (domain.Rankings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RankingsForContest", contestID, languageCode)
	ret0, _ := ret[0].(domain.Rankings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RankingsForContest indicates an expected call of RankingsForContest
func (mr *MockRankingInteractorMockRecorder) RankingsForContest(contestID, languageCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RankingsForContest", reflect.TypeOf((*MockRankingInteractor)(nil).RankingsForContest), contestID, languageCode)
}

// CurrentRegistration mocks base method
func (m *MockRankingInteractor) CurrentRegistration(userID uint64) (domain.RankingRegistration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentRegistration", userID)
	ret0, _ := ret[0].(domain.RankingRegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentRegistration indicates an expected call of CurrentRegistration
func (mr *MockRankingInteractorMockRecorder) CurrentRegistration(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentRegistration", reflect.TypeOf((*MockRankingInteractor)(nil).CurrentRegistration), userID)
}

// ContestLogs mocks base method
func (m *MockRankingInteractor) ContestLogs(contestID, userID uint64) (domain.ContestLogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContestLogs", contestID, userID)
	ret0, _ := ret[0].(domain.ContestLogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContestLogs indicates an expected call of ContestLogs
func (mr *MockRankingInteractorMockRecorder) ContestLogs(contestID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContestLogs", reflect.TypeOf((*MockRankingInteractor)(nil).ContestLogs), contestID, userID)
}

// RecentContestLogs mocks base method
func (m *MockRankingInteractor) RecentContestLogs(contestID, limit uint64) (domain.ContestLogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecentContestLogs", contestID, limit)
	ret0, _ := ret[0].(domain.ContestLogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecentContestLogs indicates an expected call of RecentContestLogs
func (mr *MockRankingInteractorMockRecorder) RecentContestLogs(contestID, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecentContestLogs", reflect.TypeOf((*MockRankingInteractor)(nil).RecentContestLogs), contestID, limit)
}
