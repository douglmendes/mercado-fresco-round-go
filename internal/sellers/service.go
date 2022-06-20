package sellers

import "fmt"

//go:generate mockgen -source=./service.go -destination=./mock/service_mock.go
type Service interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Create(cid int, commpanyName, address, telephone string) (Seller, error)
	Update(id, cid int, companyname, address, telephone string) (Seller, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service {
		repository: r,
	}
	
}

func (s service) GetAll() ([]Seller, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return []Seller{}, err
	}
	return sl, nil
	
}

func (s service) GetById(id int) (Seller, error) {
	sl, err := s.repository.GetById(id)
	if err != nil {
		return Seller{}, err
	}
	return sl, nil
	
}

func (s service) Create(cid int, commpanyName, address, telephone string) (Seller, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Seller{}, err
	}

	sl, err := s.repository.GetAll()
	if err != nil {
		return Seller{}, err
	}

	for i := range sl {
		if sl[i].Cid == cid {
				return Seller{}, fmt.Errorf("this seller already exists")
		}
	}

	lastID++

	seller, err := s.repository.Create(lastID, cid, commpanyName, address, telephone)

	if err != nil {
		return Seller{}, err
	}

	return seller, nil
}

func (s service) Update(id, cid int, companyName, address, telephone string) (Seller, error) {
	sl, err := s.repository.GetAll()
	if err != nil {
		return Seller{}, err
	}

	for i := range sl {
		if sl[i].Cid == cid {
				return Seller{}, fmt.Errorf("this seller already exists")
		}
	}

	seller, err := s.repository.Update(id, cid, companyName, address, telephone)
	if err != nil {
		return Seller{}, err
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