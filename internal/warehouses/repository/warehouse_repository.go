package repository

import (
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"log"
)

const (
	sqlCreate  = "INSERT INTO warehouse (address, telephone, warehouse_code, locality_id) VALUES (?, ?, ?, ?)"
	sqlGetAll  = "SELECT id, address, telephone, warehouse_code, locality_id FROM warehouse"
	sqlGetById = "SELECT id, address, telephone, warehouse_code, locality_id FROM warehouse WHERE id = ?"
	sqlDelete  = "DELETE FROM warehouse WHERE id = ?"
	sqlUpdate  = "UPDATE warehouse SET address = ?, telephone = ?, warehouse_code = ?, locality_id = ? WHERE id = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.WarehouseRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(id int64) (warehouse domain.Warehouse, err error) {
	row := r.db.QueryRow(sqlGetById, id)
	if err := row.Scan(
		&warehouse.Id,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.WarehouseCode,
		&warehouse.LocalityId,
	); err != nil {
		return domain.Warehouse{}, fmt.Errorf("warehouse not found")
	}
	return

}

func (r *repository) GetAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	rows, err := r.db.Query(sqlGetAll)
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

func (r *repository) Create(address, telephone, warehouseCode string, localityId int64) (domain.Warehouse, error) {
	warehouse := domain.Warehouse{
		Address:       address,
		Telephone:     telephone,
		WarehouseCode: warehouseCode,
		LocalityId:    localityId,
	}

	result, err := r.db.Exec(sqlCreate, &address, &telephone, &warehouseCode, &localityId)
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

func (r *repository) Update(
	id int64,
	address,
	telephone,
	warehouseCode string,
	localityId int64,
) (domain.Warehouse, error) {

	warehouse, err := r.GetById(id)
	if err != nil {
		return domain.Warehouse{}, fmt.Errorf("warehouse not found")
	}

	if address != "" {
		warehouse.Address = address
	}

	if telephone != "" {
		warehouse.Telephone = telephone
	}

	if warehouseCode != "" {
		warehouse.WarehouseCode = warehouseCode
	}

	if localityId != 0 {
		warehouse.LocalityId = localityId
	}

	result, err := r.db.Exec(
		sqlUpdate,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.WarehouseCode,
		&warehouse.LocalityId,
		id,
	)

	affected, err := result.RowsAffected()
	if err != nil {
		return domain.Warehouse{}, err
	}
	log.Println(affected)

	return warehouse, nil
}

func (r *repository) Delete(id int64) (err error) {
	_, err = r.db.Exec(sqlDelete, id)
	if err != nil {
		return fmt.Errorf("warehouse with id %d not found", id)
	}
	return
}
