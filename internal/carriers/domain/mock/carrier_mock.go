// Code generated by MockGen. DO NOT EDIT.
// Source: ./carrier.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	domain "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockCarrierRepository is a mock of CarrierRepository interface.
type MockCarrierRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierRepositoryMockRecorder
}

// MockCarrierRepositoryMockRecorder is the mock recorder for MockCarrierRepository.
type MockCarrierRepositoryMockRecorder struct {
	mock *MockCarrierRepository
}

// NewMockCarrierRepository creates a new mock instance.
func NewMockCarrierRepository(ctrl *gomock.Controller) *MockCarrierRepository {
	mock := &MockCarrierRepository{ctrl: ctrl}
	mock.recorder = &MockCarrierRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierRepository) EXPECT() *MockCarrierRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCarrierRepository) Create(ctx context.Context, cid, companyName, address, telephone string, localityId int) (domain.Carrier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, cid, companyName, address, telephone, localityId)
	ret0, _ := ret[0].(domain.Carrier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCarrierRepositoryMockRecorder) Create(ctx, cid, companyName, address, telephone, localityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCarrierRepository)(nil).Create), ctx, cid, companyName, address, telephone, localityId)
}

// GetAll mocks base method.
func (m *MockCarrierRepository) GetAll(ctx context.Context) ([]domain.Carrier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]domain.Carrier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCarrierRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCarrierRepository)(nil).GetAll), ctx)
}

// MockCarrierService is a mock of CarrierService interface.
type MockCarrierService struct {
	ctrl     *gomock.Controller
	recorder *MockCarrierServiceMockRecorder
}

// MockCarrierServiceMockRecorder is the mock recorder for MockCarrierService.
type MockCarrierServiceMockRecorder struct {
	mock *MockCarrierService
}

// NewMockCarrierService creates a new mock instance.
func NewMockCarrierService(ctrl *gomock.Controller) *MockCarrierService {
	mock := &MockCarrierService{ctrl: ctrl}
	mock.recorder = &MockCarrierServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarrierService) EXPECT() *MockCarrierServiceMockRecorder {
	return m.recorder
}

// CreateCarrier mocks base method.
func (m *MockCarrierService) CreateCarrier(ctx context.Context, cid, companyName, address, telephone string, localityId int) (domain.Carrier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCarrier", ctx, cid, companyName, address, telephone, localityId)
	ret0, _ := ret[0].(domain.Carrier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCarrier indicates an expected call of CreateCarrier.
func (mr *MockCarrierServiceMockRecorder) CreateCarrier(ctx, cid, companyName, address, telephone, localityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCarrier", reflect.TypeOf((*MockCarrierService)(nil).CreateCarrier), ctx, cid, companyName, address, telephone, localityId)
}
