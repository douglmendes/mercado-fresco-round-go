package service

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
)

type service struct {
	repository domain.WarehouseRepository
}

func (s *service) GetById(id int) (domain.Warehouse, error) {
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
	minimunCapacity,
	minimunTemperature int,
) (*domain.Warehouse, error) {

	lastID, err := s.repository.LastID()
	if err != nil {
		return nil, err
	}

	whList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, warehouse := range whList {
		if warehouse.WarehouseCode == warehouseCode {
			return nil, fmt.Errorf("this warehouse already exists")
		}
	}

	lastID++

	warehouse, err := s.repository.Create(lastID, address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
	if err != nil {
		return nil, err
	}

	return &warehouse, nil

}

func (s *service) Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (domain.Warehouse, error) {

	warehouse, err := s.repository.GetById(id)
	if err != nil {
		return domain.Warehouse{}, err
	}

	if warehouse.WarehouseCode == warehouseCode {
		return s.repository.Update(id, address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
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

	return s.repository.Update(id, address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}

func NewService(r domain.WarehouseRepository) domain.WarehouseService {
	return &service{
		repository: r,
	}
}
