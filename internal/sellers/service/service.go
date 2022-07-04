package service

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service {
		repository: r,
	}
	
}

func (s service) GetAll() ([]domain.Seller, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return []domain.Seller{}, err
	}
	return sl, nil
	
}

func (s service) GetById(id int) (domain.Seller, error) {
	sl, err := s.repository.GetById(id)
	if err != nil {
		return domain.Seller{}, err
	}
	return sl, nil
	
}

func (s service) Create(cid int, companyName, address, telephone string) (domain.Seller, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Seller{}, err
	}

	sl, err := s.repository.GetAll()
	if err != nil {
		return domain.Seller{}, err
	}

	for i := range sl {
		if sl[i].Cid == cid {
				return domain.Seller{}, fmt.Errorf("this seller already exists")
		}
	}

	lastID++

	seller, err := s.repository.Create(lastID, cid, companyName, address, telephone)

	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

func (s service) Update(id, cid int, companyName, address, telephone string) (domain.Seller, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return domain.Seller{}, err
	}

	for i := range sl {
		if sl[i].Cid == cid {
				return domain.Seller{}, fmt.Errorf("this seller already exists")
		}
	}

	seller, err := s.repository.Update(id, cid, companyName, address, telephone)
	if err != nil {
		return domain.Seller{}, err
	}

	return seller, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}