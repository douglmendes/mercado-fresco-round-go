package repository

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type repository struct {
	db store.Store
}

func (r *repository) GetById(id int) (domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	if err := r.db.Read(&warehouses); err != nil {
		return domain.Warehouse{}, err
	}

	for _, warehouse := range warehouses {
		if warehouse.Id == id {
			return warehouse, nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("warehouse not found")
}

func (r *repository) GetAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	if err := r.db.Read(&warehouses); err != nil {
		return []domain.Warehouse{}, err
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
) (domain.Warehouse, error) {

	var warehouses []domain.Warehouse
	if err := r.db.Read(&warehouses); err != nil {
		return domain.Warehouse{}, err
	}

	warehouse := domain.Warehouse{
		id,
		address,
		telephone,
		warehouseCode,
		minimunCapacity,
		minimunTemperature,
	}

	warehouses = append(warehouses, warehouse)

	if err := r.db.Write(warehouses); err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (r *repository) LastID() (int, error) {

	var warehouses []domain.Warehouse

	if err := r.db.Read(&warehouses); err != nil {
		return 0, err
	}

	if len(warehouses) == 0 {
		return 0, nil
	}
	return warehouses[len(warehouses)-1].Id, nil
}

func (r *repository) Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	if err := r.db.Read(&warehouses); err != nil {
		return domain.Warehouse{}, err
	}

	wh := domain.Warehouse{
		Address:            address,
		Telephone:          telephone,
		WarehouseCode:      warehouseCode,
		MinimunCapacity:    minimunCapacity,
		MinimunTemperature: minimunTemperature,
	}

	updated := false
	for i := range warehouses {
		if warehouses[i].Id == id {
			wh = warehouses[i]

			if address != "" {
				wh.Address = address
			}

			if telephone != "" {
				wh.Telephone = telephone
			}

			if warehouseCode != "" {
				wh.WarehouseCode = warehouseCode
			}

			if minimunTemperature != 0 {
				wh.MinimunTemperature = minimunTemperature
			}

			if minimunCapacity != 0 {
				wh.MinimunCapacity = minimunCapacity
			}

			warehouses[i] = wh
			updated = true
			if err := r.db.Write(warehouses); err != nil {
				return domain.Warehouse{}, err
			}
		}
	}
	if !updated {
		return domain.Warehouse{}, fmt.Errorf("warehouse with id %d not found", id)
	}
	return wh, nil
}

func (r *repository) Delete(id int) error {

	var warehouses []domain.Warehouse
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

func NewRepository(db store.Store) domain.WarehouseRepository {
	return &repository{
		db: db,
	}
}
