// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	domain "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
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
func (m *MockRepository) Create(arg0 context.Context, arg1, arg2 string, arg3, arg4, arg5 int) (*domain.InboudOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*domain.InboudOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll(arg0 context.Context) ([]domain.InboudOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]domain.InboudOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), arg0)
}

// GetByEmployee mocks base method.
func (m *MockRepository) GetByEmployee(ctx context.Context, employee int64) ([]domain.EmployeeInboudOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmployee", ctx, employee)
	ret0, _ := ret[0].([]domain.EmployeeInboudOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmployee indicates an expected call of GetByEmployee.
func (mr *MockRepositoryMockRecorder) GetByEmployee(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmployee", reflect.TypeOf((*MockRepository)(nil).GetByEmployee), ctx, employee)
}

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
func (m *MockService) Create(arg0 context.Context, arg1, arg2 string, arg3, arg4, arg5 int) (*domain.InboudOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*domain.InboudOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetByEmployee mocks base method.
func (m *MockService) GetByEmployee(ctx context.Context, employee int64) ([]domain.EmployeeInboudOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmployee", ctx, employee)
	ret0, _ := ret[0].([]domain.EmployeeInboudOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmployee indicates an expected call of GetByEmployee.
func (mr *MockServiceMockRecorder) GetByEmployee(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmployee", reflect.TypeOf((*MockService)(nil).GetByEmployee), ctx, employee)
}