package service

import (
	"context"
	"fmt"
	repositoryEmployee "github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
)

type service struct {
	repository         domain.Repository
	repositoryEmployee repositoryEmployee.Repository
}

func NewService(r domain.Repository, re repositoryEmployee.Repository) domain.Service {
	return &service{
		repository:         r,
		repositoryEmployee: re,
	}
}

func (s service) Create(ctx context.Context, orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (*domain.InboudOrder, error) {

	_, err := s.repositoryEmployee.GetById(ctx, int64(employeeId))
	if err != nil {
		return nil, fmt.Errorf("employee number not found")
	}
	in, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for i := range in {
		if in[i].OrderNumber == orderNumber {
			return nil, fmt.Errorf("order number already exists")
		}
	}
	io, err := s.repository.Create(ctx, orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return nil, err
	}
	return io, nil
}

func (s service) GetByEmployee(ctx context.Context, id int64) ([]domain.EmployeeInboudOrder, error) {
	io, err := s.repository.GetByEmployee(ctx, id)
	if err != nil {
		return nil, err
	}
	return io, nil
}
