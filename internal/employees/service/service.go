package service

import (
	"context"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	"log"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}

}

func (s service) GetAll(ctx context.Context) ([]domain.Employee, error) {

	emp, err := s.repository.GetAll(ctx)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	return emp, nil

}

func (s service) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	emp, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return emp, nil

}

func (s service) Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	emp, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this card number id already exists")
		}
	}
	/*lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Employee{}, err
	}

	lastID++
	*/
	employee, err := s.repository.Create(ctx, cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s service) Update(ctx context.Context, id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	emp, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this employee already exists")
		}
	}
	employee, err := s.repository.Update(ctx, id, cardNumberId, firstName, lastName, warehouseId)
	if err != nil {
		return nil, err
	}
	return employee, err
}

func (s service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
