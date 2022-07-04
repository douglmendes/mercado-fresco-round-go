package domain

import "fmt"

type Section struct {
	Id                 int `json:"id,omitempty"`
	SectionNumber      int `json:"section_number,omitempty"`
	CurrentTemperature int `json:"current_temperature,omitempty"`
	MinimumTemperature int `json:"minimum_temperature,omitempty"`
	CurrentCapacity    int `json:"current_capacity,omitempty"`
	MinimumCapacity    int `json:"minimum_capacity,omitempty"`
	MaximumCapacity    int `json:"maximum_capacity,omitempty"`
	WarehouseId        int `json:"warehouse_id,omitempty"`
	ProductTypeId      int `json:"product_type_id,omitempty"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (*Section, error)
	LastID() (int, error)
	Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (*Section, error)
	Exists(id int) error
	Update(id int, args map[string]int) (*Section, error)
	Delete(id int) (*Section, error)
}

type Service interface {
	GetAll() ([]Section, error)
	GetById(id int) (*Section, error)
	Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (*Section, error)
	Update(id int, args map[string]int) (*Section, error)
	Delete(id int) (*Section, error)
}

type ErrorNotFound struct {
	Id int
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("section %d not found in database", e.Id)
}

type ErrorConflict struct {
	SectionNumber int
}

func (e *ErrorConflict) Error() string {
	return fmt.Sprintf("a section with number %d already exists", e.SectionNumber)
}
