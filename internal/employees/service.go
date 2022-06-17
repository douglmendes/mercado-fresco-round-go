package employees

import "fmt"

//go:generate mockgen -source=./service.go -destination=./mock/service_mock.go
type Service interface {
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	Create(cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error)
	Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}

}

func (s service) GetAll() ([]Employee, error) {
	emp, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return emp, nil

}

func (s service) GetById(id int) (Employee, error) {
	sl, err := s.repository.GetById(id)
	if err != nil {
		return Employee{}, err
	}
	return sl, nil

}

func (s service) Create(cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error) {
	emp, err := s.repository.GetAll()

	if err != nil {
		return Employee{}, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return Employee{}, fmt.Errorf("this card number id already exists")
		}
	}
	lastID, err := s.repository.LastID()
	if err != nil {
		return Employee{}, err
	}

	sl, err := s.repository.GetAll()
	if err != nil {
		return Employee{}, err
	}

	for i := range sl {
		if sl[i].CardNumberId == cardNumberId {
			return Employee{}, fmt.Errorf("this card number id already exists")
		}
	}

	lastID++

	employee, err := s.repository.Create(lastID, cardNumberId, firstName, lastName, warehouseId)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s service) Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error) {
	emp, err := s.repository.GetAll()

	if err != nil {
		return Employee{}, err
	}

	for e := range emp {
		if emp[e].CardNumberId == cardNumberId {
			return Employee{}, fmt.Errorf("this employee already exists")
		}
	}
	employee, err := s.repository.Update(id, cardNumberId, firstName, lastName, warehouseId)
	if err != nil {
		return Employee{}, err
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
