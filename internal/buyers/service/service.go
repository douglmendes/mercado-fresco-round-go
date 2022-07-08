package service

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]domain.Buyer, error) {
	buy, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buy, nil

}

func (s service) GetById(id int) (*domain.Buyer, error) {
	buy, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return buy, nil

}

func (s service) Create(cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	buy, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	for b := range buy {
		if buy[b].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this card number id already exists")
		}
	}
	buyer, _ := s.repository.Create(cardNumberId, firstName, lastName)

	return buyer, nil
}

func (s service) Update(id int, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range sl {
		if sl[i].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this Buyer already exists")
		}
	}

	Buyer, err := s.repository.Update(id, cardNumberId, firstName, lastName)
	if err != nil {
		return nil, err
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
