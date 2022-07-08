package service

import (
	"context"
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	localityD "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
)

type service struct {
	repository domain.Repository
	localityRepository localityD.LocalityRepository
}

func NewService(r domain.Repository, rL localityD.LocalityRepository) domain.Service {
	return &service{
		repository: r,
		localityRepository: rL,
	}
}

func (s service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	sl, err := s.repository.GetAll(ctx)
	if err != nil {
		return []domain.Seller{}, err
	}
	return sl, nil

}

func (s service) GetById(ctx context.Context, id int) (domain.Seller, error) {
	sl, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Seller{}, err
	}
	return sl, nil

}

func (s service) Create(ctx context.Context, cid int, companyName, address, telephone string, localityId int) (domain.Seller, error) {

	sl, err := s.repository.GetAll(ctx)
	if err != nil {
		return domain.Seller{}, err
	}

	for i := range sl {
		if sl[i].Cid == cid {
			return domain.Seller{}, fmt.Errorf("this seller already exists")
		}
	}

	_, err = s.localityRepository.GetById(ctx, localityId)
	if err != nil {
		return domain.Seller{}, fmt.Errorf("locality %d not found", localityId)
	}

	seller, err := s.repository.Create(ctx, cid, companyName, address, telephone, localityId)

	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

func (s service) Update(ctx context.Context, id, cid int, companyName, address, telephone string, localityId int) (domain.Seller, error) {
	sl, err := s.repository.GetAll(ctx)
	if err != nil {
		return domain.Seller{}, err
	}

	for i := range sl {
		if sl[i].Cid == cid {
			return domain.Seller{}, fmt.Errorf("this seller already exists")
		}
	}

	seller, err := s.repository.Update(ctx, id, cid, companyName, address, telephone, localityId)
	if err != nil {
		return domain.Seller{}, err
	}

	return seller, err
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
