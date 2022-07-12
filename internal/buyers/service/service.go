package service

import (
	"context"
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

func (s service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	buy, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return buy, nil

}

func (s service) GetById(ctx context.Context, id int) (*domain.Buyer, error) {
	buy, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return buy, nil

}

func (s service) GetOrdersByBuyers(ctx context.Context, id int) ([]domain.OrdersByBuyers, error) {
	lc, err := s.repository.GetOrdersByBuyers(ctx, id)
	if err != nil {
		return []domain.OrdersByBuyers{}, err
	}
	return lc, nil
}

func (s service) Create(ctx context.Context, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	buy, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	for b := range buy {
		if buy[b].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this card number id already exists")
		}
	}
	buyer, _ := s.repository.Create(ctx, cardNumberId, firstName, lastName)

	return buyer, nil
}

func (s service) Update(ctx context.Context, id int, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	sl, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := range sl {
		if sl[i].CardNumberId == cardNumberId {
			return nil, fmt.Errorf("this Buyer already exists")
		}
	}

	Buyer, err := s.repository.Update(ctx, id, cardNumberId, firstName, lastName)
	if err != nil {
		return nil, err
	}

	return Buyer, err
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return err
}
