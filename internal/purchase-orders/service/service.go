package service

import (
	"context"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(ctx context.Context, orderNumber string, orderDate string, trackingCode string, buyerId int, carrierId int, productRecordId int, orderStatusId int) (*domain.PurchaseOrder, error) {
	po, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for p := range po {
		if po[p].OrderNumber == orderNumber {
			return nil, fmt.Errorf("order number already exists")
		}
	}
	por, err := s.repository.Create(ctx, orderNumber, orderDate, trackingCode, buyerId, carrierId, productRecordId, orderStatusId)
	if err != nil {
		return nil, err
	}
	return por, nil
}
