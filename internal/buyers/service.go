package buyers

import "fmt"

type Service interface {
	GetById(id int) (Buyer, error)
	GetAll() ([]Buyer, error)
	Create(cardNumberId, firstName, lastName string) (Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (Buyer, error)
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

func (s service) GetAll() ([]Buyer, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return sl, nil

}

func (s service) GetById(id int) (Buyer, error) {
	sl, err := s.repository.GetById(id)
	if err != nil {
		return Buyer{}, err
	}
	return sl, nil

}

func (s service) Create(cardNumberId, firstName, lastName string) (Buyer, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Buyer{}, err
	}

	sl, err := s.repository.GetAll()
	if err != nil {
		return Buyer{}, err
	}

	for i := range sl {
		if sl[i].CardNumberId == cardNumberId {
			return Buyer{}, fmt.Errorf("this Buyer already exists")
		}
	}

	lastID++

	Buyer, err := s.repository.Create(lastID, cardNumberId, firstName, lastName)

	if err != nil {
		return Buyer, err
	}

	return Buyer, nil
}

func (s service) Update(id int, cardNumberId, firstName, lastName string) (Buyer, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return Buyer{}, err
	}

	for i := range sl {
		if sl[i].CardNumberId == cardNumberId {
			return Buyer{}, fmt.Errorf("this Buyer already exists")
		}
	}

	Buyer, err := s.repository.Update(id, cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer, err
	}

	return Buyer, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
