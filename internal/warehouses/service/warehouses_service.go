package service

import (
	"context"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type service struct {
	repository domain.WarehouseRepository
}

func NewService(r domain.WarehouseRepository) domain.WarehouseService {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(ctx context.Context, id int) (domain.Warehouse, error) {
	warehouse, err := s.repository.GetById(ctx, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	warehouses, err := s.repository.GetAll(ctx)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return []domain.Warehouse{}, err
	}
	return warehouses, nil
}

func (s *service) Create(
	ctx context.Context,
	address,
	telephone,
	warehouseCode string,
	localityId int,
) (*domain.Warehouse, error) {

	whList, err := s.repository.GetAll(ctx)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return nil, err
	}

	for _, warehouse := range whList {
		if warehouse.WarehouseCode == warehouseCode {
			return nil, fmt.Errorf("this warehouse already exists")
		}
	}

	warehouse, err := s.repository.Create(ctx, address, telephone, warehouseCode, localityId)
	if err != nil {
		return nil, err
	}

	return &warehouse, nil

}

func (s *service) Update(
	ctx context.Context,
	id int,
	address,
	telephone,
	warehouseCode string,
	localityId int,
) (domain.Warehouse, error) {

	warehouse, err := s.repository.GetById(ctx, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, err
	}

	if warehouse.WarehouseCode == warehouseCode {
		return s.repository.Update(ctx, id, address, telephone, warehouseCode, localityId)
	}

	whList, err := s.repository.GetAll(ctx)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, err
	}

	for _, warehouse := range whList {
		if warehouse.WarehouseCode == warehouseCode {
			return domain.Warehouse{}, fmt.Errorf("this warehouse already exists")
		}
	}

	return s.repository.Update(ctx, id, address, telephone, warehouseCode, localityId)
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return err
	}
	return nil
}
