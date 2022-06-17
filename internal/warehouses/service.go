package warehouses

import "fmt"

//go:generate mockgen -source=./service.go -destination=./mock/service_mock.go
type Service interface {
	Create(address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s *service) GetById(id int) (Warehouse, error) {
	warehouse, err := s.repository.GetById(id)
	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) GetAll() ([]Warehouse, error) {
	warehouses, err := s.repository.GetAll()
	if err != nil {
		return []Warehouse{}, err
	}
	return warehouses, nil
}

func (s *service) Create(
	address,
	telephone,
	warehouseCode string,
	minimunCapacity,
	minimunTemperature int,
) (*Warehouse, error) {

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

func (s *service) Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error) {

	warehouse, err := s.repository.GetById(id)
	if err != nil {
		return Warehouse{}, err
	}

	if warehouse.WarehouseCode == warehouseCode {
		return s.repository.Update(id, address, telephone, warehouseCode, minimunCapacity, minimunTemperature)
	}

	whList, err := s.repository.GetAll()
	if err != nil {
		return Warehouse{}, err
	}

	for _, warehouse := range whList {
		if warehouse.WarehouseCode == warehouseCode {
			return Warehouse{}, fmt.Errorf("this warehouse already exists")
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

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
