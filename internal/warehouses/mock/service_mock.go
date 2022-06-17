// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_warehouses is a generated GoMock package.
package mock_warehouses

import (
	reflect "reflect"

	warehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockService) Create(address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (*warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
	ret0, _ := ret[0].(*warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(address, telephone, warehouseCode, minimunCapacity, minimunTemperature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
}

// Delete mocks base method.
func (m *MockService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockService) GetAll() ([]warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockService) GetById(id int) (warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockServiceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockService)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockService) Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
	ret0, _ := ret[0].(warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(id, address, telephone, warehouseCode, minimunCapacity, minimunTemperature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), id, address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
}
