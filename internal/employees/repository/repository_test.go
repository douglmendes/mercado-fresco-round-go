package repository

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func callMock(t *testing.T) (*mock_domain.MockService, *EmployeesController) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_domain.NewMockRepository(ctrl)
	handler := NewRepository(service)

	return service, handler
}
