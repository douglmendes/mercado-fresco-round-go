package service

import (
	"context"

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

// func (s service) GetAll(ctx context.Context) ([]domain.Locality, error) {
// 	lc, err := s.repository.GetAll(ctx)
// 	if err != nil {
// 		return []domain.Locality{}, err
// 	}
// 	return lc, nil
// }

func (s service) GetById(ctx context.Context, id int) (domain.Locality, error) {
	lc, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Locality{}, err
	}
	return lc, nil
}

func (s service) GetBySellers(ctx context.Context, id int) ([]domain.SellersByLocality, error) {
	lc, err := s.repository.GetBySellers(ctx, id)
	if err != nil {
		return []domain.SellersByLocality{}, err
	}
	return lc, nil
}

func (s service) Create(ctx context.Context, localityname, provinceName, countryName string) (domain.Locality, error) {
	// lc , err := s.repository.GetAll(ctx)
	// if err != nil {
	// 	return domain.Locality{}, err
	// }

	// for i := range lc {
	// 	if lc[i].Id ==
	// }

	locality, err := s.repository.Create(ctx, localityname, provinceName, countryName)

	if err != nil {
		return domain.Locality{}, err
	}

	return locality, nil
}

func (s service) Update(ctx context.Context, id int, localityname, provinceName, countryName string) (domain.Locality, error) {
	locality, err := s.repository.Update(ctx, id, localityname, provinceName, countryName)
	if err != nil {
		return domain.Locality{}, err
	}

	return locality, err
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
