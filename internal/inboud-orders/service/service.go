package service

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (*domain.InboudOrder, error) {
	in, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range in {
		if in[i].OrderNumber == orderNumber {
			return nil, fmt.Errorf("this order number already exists")
		}
	}
	io, err := s.repository.Create(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return nil, err
	}
	return io, nil
}

func (s service) GetByEmployee(id int64) ([]domain.EmployeeInboudOrder, error) {
	io, err := s.repository.GetByEmployee(id)
	if err != nil {
		return nil, err
	}
	return io, nil
}
