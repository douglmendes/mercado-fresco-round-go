package service

import (
	"context"
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
)

type service struct {
	repository domain.LocalityRepository
}

func NewService(r domain.LocalityRepository) domain.LocalityService {
	return &service{
		repository: r,
	}
}

func (s service) GetBySellers(ctx context.Context, id int) ([]domain.SellersByLocality, error) {
	lc, err := s.repository.GetBySellers(ctx, id)
	if err != nil {
		return []domain.SellersByLocality{}, err
	}
	return lc, nil
}

func (s service) Create(ctx context.Context, zipCode, localityname, provinceName, countryName string) (domain.Locality, error) {

	sl, err := s.repository.GetAll(ctx)
	if err != nil {
		return domain.Locality{}, err
	}

	for i := range sl {
		if sl[i].ZipCode == zipCode {
			return domain.Locality{}, fmt.Errorf("this locality already exists")
		}
	}

	locality, err := s.repository.Create(ctx, zipCode, localityname, provinceName, countryName)

	if err != nil {
		return domain.Locality{}, err
	}

	return locality, nil
}