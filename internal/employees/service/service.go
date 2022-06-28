package service

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}

}

func (s service) GetAll() ([]domain.Employee, error) {
	emp, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return emp, nil

}

func (s service) GetById(id int64) (*domain.Employee, error) {
	emp, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return emp, nil

}

func (s service) Create(cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	emp, err := s.repository.GetAll()

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
	employee, err := s.repository.Create(cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s service) Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	emp, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this employee already exists")
		}
	}
	employee, err := s.repository.Update(id, cardNumberId, firstName, lastName, warehouseId)
	if err != nil {
		return nil, err
	}
	return employee, err
}

/*
func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
*/
