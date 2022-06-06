package warehouses

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	Create(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error)
	LastID() (int, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetById(id int) (Warehouse, error) {
	var warehouses []Warehouse

	if err := r.db.Read(&warehouses); err != nil {
		return Warehouse{}, err
	}

	for _, warehouse := range warehouses {
		if warehouse.Id == id {
			return warehouse, nil
		}
	}
	return Warehouse{}, fmt.Errorf("warehouse not found")
}

func (r *repository) GetAll() ([]Warehouse, error) {
	var warehouses []Warehouse

	if err := r.db.Read(&warehouses); err != nil {
		return []Warehouse{}, nil
	}

	return warehouses, nil
}

func (r *repository) Create(
	id int,
	address,
	telephone,
	warehouseCode string,
	minimunCapacity,
	minimunTemperature int,
) (Warehouse, error) {

	var warehouses []Warehouse
	if err := r.db.Read(&warehouses); err != nil {
		return Warehouse{}, nil
	}

	warehouse := Warehouse{
		id,
		address,
		telephone,
		warehouseCode,
		minimunCapacity,
		minimunTemperature,
	}

	warehouses = append(warehouses, warehouse)

	if err := r.db.Write(warehouses); err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil
}

func (r *repository) LastID() (int, error) {

	var warehouses []Warehouse

	if err := r.db.Read(&warehouses); err != nil {
		return 0, err
	}

	if len(warehouses) == 0 {
		return 0, nil
	}
	return warehouses[len(warehouses)-1].Id, nil
}

func (r *repository) Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error) {

	var warehouses []Warehouse
	if err := r.db.Read(&warehouses); err != nil {
		return Warehouse{}, err
	}

	wh := Warehouse{
		Address:            address,
		Telephone:          telephone,
		WarehouseCode:      warehouseCode,
		MinimunCapacity:    minimunCapacity,
		MinimunTemperature: minimunTemperature,
	}

	updated := false
	for i := range warehouses {
		if warehouses[i].Id == id {
			wh.Id = id
			warehouses[i] = wh
			updated = true
		}
	}
	if !updated {
		return Warehouse{}, fmt.Errorf("warehouse with id %d not found", id)
	}

	if err := r.db.Write(warehouses); err != nil {
		return Warehouse{}, err
	}

	return wh, nil
}

func (r *repository) Delete(id int) error {

	var warehouses []Warehouse
	if err := r.db.Read(&warehouses); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range warehouses {
		if warehouses[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("warehouse with id %d not found", id)
	}

	warehouses = append(warehouses[:index], warehouses[index+1:]...)

	if err := r.db.Write(warehouses); err != nil {
		return err
	}

	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
