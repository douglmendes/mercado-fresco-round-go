package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
)

const (
	sqlCreate = "INSERT INTO warehouse (address, telephone, warehouse_code, locality_id) VALUES (?, ?, ?, ?)"
	sqlGetAll = "SELECT id, address, telephone, warehouse_code, locality_id FROM warehouse"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.WarehouseRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	//if err := r.db.Read(&warehouses); err != nil {
	//	return domain.Warehouse{}, err
	//}

	for _, warehouse := range warehouses {
		if warehouse.Id == id {
			return warehouse, nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("warehouse not found")
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		return []domain.Warehouse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var warehouse domain.Warehouse
		if err := rows.Scan(
			&warehouse.Id,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.WarehouseCode,
			&warehouse.LocalityId,
		); err != nil {
			return warehouses, err
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (r *repository) Create(ctx context.Context, address, telephone, warehouseCode string, localityId int64) (domain.Warehouse, error) {
	warehouse := domain.Warehouse{
		Address:       address,
		Telephone:     telephone,
		WarehouseCode: warehouseCode,
		LocalityId:    localityId,
	}

	result, err := r.db.ExecContext(ctx, sqlCreate, &address, &telephone, &warehouseCode, &localityId)
	if err != nil {
		return domain.Warehouse{}, nil
	}

	incrementId, err := result.LastInsertId()
	if err != nil {
		return domain.Warehouse{}, err
	}

	warehouse.Id = incrementId

	return warehouse, nil
}

func (r *repository) LastID() (int64, error) {

	var warehouses []domain.Warehouse

	//if err := r.db.Read(&warehouses); err != nil {
	//	return 0, err
	//}

	if len(warehouses) == 0 {
		return 0, nil
	}
	return warehouses[len(warehouses)-1].Id, nil
}

func (r *repository) Update(ctx context.Context, id int64, address, telephone, warehouseCode string, localityId int64) (domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	//if err := r.db.Read(&warehouses); err != nil {
	//	return domain.Warehouse{}, err
	//}

	wh := domain.Warehouse{
		Address:       address,
		Telephone:     telephone,
		WarehouseCode: warehouseCode,
		LocalityId:    localityId,
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

			if localityId != 0 {
				wh.LocalityId = localityId
			}

			//if minimunTemperature != 0 {
			//	wh.MinimunTemperature = minimunTemperature
			//}
			//
			//if minimunCapacity != 0 {
			//	wh.MinimunCapacity = minimunCapacity
			//}

			warehouses[i] = wh
			updated = true
			//if err := r.db.Write(warehouses); err != nil {
			//	return domain.Warehouse{}, err
			//}
		}
	}
	if !updated {
		return domain.Warehouse{}, fmt.Errorf("warehouse with id %d not found", id)
	}
	return wh, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {

	var warehouses []domain.Warehouse
	//if err := r.db.Read(&warehouses); err != nil {
	//	return err
	//}

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

	//if err := r.db.Write(warehouses); err != nil {
	//	return err
	//}

	return nil
}
