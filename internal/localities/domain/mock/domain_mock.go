// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	domain "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockLocalityRepository is a mock of LocalityRepository interface.
type MockLocalityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLocalityRepositoryMockRecorder
}

// MockLocalityRepositoryMockRecorder is the mock recorder for MockLocalityRepository.
type MockLocalityRepositoryMockRecorder struct {
	mock *MockLocalityRepository
}

// NewMockLocalityRepository creates a new mock instance.
func NewMockLocalityRepository(ctrl *gomock.Controller) *MockLocalityRepository {
	mock := &MockLocalityRepository{ctrl: ctrl}
	mock.recorder = &MockLocalityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalityRepository) EXPECT() *MockLocalityRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLocalityRepository) Create(ctx context.Context, localityName, provinceName, countryName string) (domain.Locality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, localityName, provinceName, countryName)
	ret0, _ := ret[0].(domain.Locality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLocalityRepositoryMockRecorder) Create(ctx, localityName, provinceName, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLocalityRepository)(nil).Create), ctx, localityName, provinceName, countryName)
}

// Delete mocks base method.
func (m *MockLocalityRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLocalityRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLocalityRepository)(nil).Delete), ctx, id)
}

// GetById mocks base method.
func (m *MockLocalityRepository) GetById(ctx context.Context, id int) (domain.Locality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(domain.Locality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockLocalityRepositoryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockLocalityRepository)(nil).GetById), ctx, id)
}

// GetBySellers mocks base method.
func (m *MockLocalityRepository) GetBySellers(ctx context.Context, id int) ([]domain.SellersByLocality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySellers", ctx, id)
	ret0, _ := ret[0].([]domain.SellersByLocality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySellers indicates an expected call of GetBySellers.
func (mr *MockLocalityRepositoryMockRecorder) GetBySellers(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySellers", reflect.TypeOf((*MockLocalityRepository)(nil).GetBySellers), ctx, id)
}

// Update mocks base method.
func (m *MockLocalityRepository) Update(ctx context.Context, id int, localityName, provinceName, countryName string) (domain.Locality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, localityName, provinceName, countryName)
	ret0, _ := ret[0].(domain.Locality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLocalityRepositoryMockRecorder) Update(ctx, id, localityName, provinceName, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLocalityRepository)(nil).Update), ctx, id, localityName, provinceName, countryName)
}

// MockLocalityService is a mock of LocalityService interface.
type MockLocalityService struct {
	ctrl     *gomock.Controller
	recorder *MockLocalityServiceMockRecorder
}

// MockLocalityServiceMockRecorder is the mock recorder for MockLocalityService.
type MockLocalityServiceMockRecorder struct {
	mock *MockLocalityService
}

// NewMockLocalityService creates a new mock instance.
func NewMockLocalityService(ctrl *gomock.Controller) *MockLocalityService {
	mock := &MockLocalityService{ctrl: ctrl}
	mock.recorder = &MockLocalityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalityService) EXPECT() *MockLocalityServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLocalityService) Create(ctx context.Context, localityName, provinceName, countryName string) (domain.Locality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, localityName, provinceName, countryName)
	ret0, _ := ret[0].(domain.Locality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLocalityServiceMockRecorder) Create(ctx, localityName, provinceName, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLocalityService)(nil).Create), ctx, localityName, provinceName, countryName)
}

// Delete mocks base method.
func (m *MockLocalityService) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLocalityServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLocalityService)(nil).Delete), ctx, id)
}

// GetById mocks base method.
func (m *MockLocalityService) GetById(ctx context.Context, id int) (domain.Locality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(domain.Locality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockLocalityServiceMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockLocalityService)(nil).GetById), ctx, id)
}

// GetBySellers mocks base method.
func (m *MockLocalityService) GetBySellers(ctx context.Context, id int) ([]domain.SellersByLocality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySellers", ctx, id)
	ret0, _ := ret[0].([]domain.SellersByLocality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySellers indicates an expected call of GetBySellers.
func (mr *MockLocalityServiceMockRecorder) GetBySellers(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySellers", reflect.TypeOf((*MockLocalityService)(nil).GetBySellers), ctx, id)
}

// Update mocks base method.
func (m *MockLocalityService) Update(ctx context.Context, id int, localityName, provinceName, countryName string) (domain.Locality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, localityName, provinceName, countryName)
	ret0, _ := ret[0].(domain.Locality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLocalityServiceMockRecorder) Update(ctx, id, localityName, provinceName, countryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLocalityService)(nil).Update), ctx, id, localityName, provinceName, countryName)
}
