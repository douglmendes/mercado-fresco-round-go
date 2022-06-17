// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock_employees is a generated GoMock package.
package mock_employees

import (
	reflect "reflect"

	employees "github.com/douglmendes/mercado-fresco-round-go/internal/employees"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(id int, cardNumberId, firstName, lastName string, warehouseId int) (employees.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", id, cardNumberId, firstName, lastName, warehouseId)
	ret0, _ := ret[0].(employees.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(id, cardNumberId, firstName, lastName, warehouseId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), id, cardNumberId, firstName, lastName, warehouseId)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll() ([]employees.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]employees.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockRepository) GetById(id int) (employees.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(employees.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), id)
}

// LastID mocks base method.
func (m *MockRepository) LastID() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockRepositoryMockRecorder) LastID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockRepository)(nil).LastID))
}

// Update mocks base method.
func (m *MockRepository) Update(id int, cardNumberId, firstName, lastName string, warehouseId int) (employees.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, cardNumberId, firstName, lastName, warehouseId)
	ret0, _ := ret[0].(employees.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(id, cardNumberId, firstName, lastName, warehouseId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), id, cardNumberId, firstName, lastName, warehouseId)
}
