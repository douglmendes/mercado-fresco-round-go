package service

import (
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

/*
func (s service) GetById(id int) (domain.Employee, error) {
	sl, err := s.repository.GetById(id)
	if err != nil {
		return domain.Employee{}, err
	}
	return sl, nil

}

func (s service) Create(cardNumberId string, firstName string, lastName string, warehouseId int) (domain.Employee, error) {
	emp, err := s.repository.GetAll()

	if err != nil {
		return domain.Employee{}, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return domain.Employee{}, fmt.Errorf("this card number id already exists")
		}
	}
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Employee{}, err
	}

	lastID++

	employee, err := s.repository.Create(lastID, cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (s service) Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (domain.Employee, error) {
	emp, err := s.repository.GetAll()

	if err != nil {
		return domain.Employee{}, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return domain.Employee{}, fmt.Errorf("this employee already exists")
		}
	}
	employee, err := s.repository.Update(id, cardNumberId, firstName, lastName, warehouseId)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
*/
