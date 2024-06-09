// Code generated by MockGen. DO NOT EDIT.
// Source: graphblog/internal/user (interfaces: DAO)
//
// Generated by this command:
//
//	mockgen -destination=dao_mock.go -package=user . DAO
//

// Package user is a generated GoMock package.
package user

import (
	model "graphblog/graph/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockDAO is a mock of DAO interface.
type MockDAO struct {
	ctrl     *gomock.Controller
	recorder *MockDAOMockRecorder
}

// MockDAOMockRecorder is the mock recorder for MockDAO.
type MockDAOMockRecorder struct {
	mock *MockDAO
}

// NewMockDAO creates a new mock instance.
func NewMockDAO(ctrl *gomock.Controller) *MockDAO {
	mock := &MockDAO{ctrl: ctrl}
	mock.recorder = &MockDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAO) EXPECT() *MockDAOMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDAO) Create(arg0 model.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDAOMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDAO)(nil).Create), arg0)
}

// GetAll mocks base method.
func (m *MockDAO) GetAll() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDAOMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDAO)(nil).GetAll))
}
