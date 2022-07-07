package service

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
)

type service struct {
	repository domain.WarehouseRepository
}

func NewService(r domain.WarehouseRepository) domain.WarehouseService {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(id int64) (domain.Warehouse, error) {
	warehouse, err := s.repository.GetById(id)
	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) GetAll() ([]domain.Warehouse, error) {
	warehouses, err := s.repository.GetAll()
	if err != nil {
		return []domain.Warehouse{}, err
	}
	return warehouses, nil
}

func (s *service) Create(
	address,
	telephone,
	warehouseCode string,
	localityId int64,
) (*domain.Warehouse, error) {

	whList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, warehouse := range whList {
		if warehouse.WarehouseCode == warehouseCode {
			return nil, fmt.Errorf("this warehouse already exists")
		}
	}

	warehouse, err := s.repository.Create(address, telephone, warehouseCode, localityId)
	if err != nil {
		return nil, err
	}

	return &warehouse, nil

}

func (s *service) Update(
	id int64,
	address,
	telephone,
	warehouseCode string,
	localityId int64,
) (domain.Warehouse, error) {

	warehouse, err := s.repository.GetById(id)
	if err != nil {
		return domain.Warehouse{}, err
	}

	if warehouse.WarehouseCode == warehouseCode {
		return s.repository.Update(id, address, telephone, warehouseCode, localityId)
	}

	whList, err := s.repository.GetAll()
	if err != nil {
		return domain.Warehouse{}, err
	}

	for _, warehouse := range whList {
		if warehouse.WarehouseCode == warehouseCode {
			return domain.Warehouse{}, fmt.Errorf("this warehouse already exists")
		}
	}

	return s.repository.Update(id, address, telephone, warehouseCode, localityId)
}

func (s *service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
