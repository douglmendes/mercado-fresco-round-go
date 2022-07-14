package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.WarehouseRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(ctx context.Context, id int) (warehouse domain.Warehouse, err error) {
	row := r.db.QueryRowContext(ctx, sqlGetById, id)
	if err := row.Scan(
		&warehouse.Id,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.WarehouseCode,
		&warehouse.LocalityId,
	); err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, fmt.Errorf("warehouse not found")
	}
	return

}

func (r *repository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse

	rows, err := r.db.QueryContext(ctx, sqlGetAll)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
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
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			return warehouses, err
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (r *repository) Create(ctx context.Context, address, telephone, warehouseCode string, localityId int) (domain.Warehouse, error) {
	warehouse := domain.Warehouse{
		Address:       address,
		Telephone:     telephone,
		WarehouseCode: warehouseCode,
		LocalityId:    localityId,
	}

	result, err := r.db.ExecContext(ctx, sqlCreate, &address, &telephone, &warehouseCode, &localityId)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, err
	}

	incrementId, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, err
	}

	warehouse.Id = int(incrementId)

	return warehouse, nil
}

func (r *repository) Update(
	ctx context.Context,
	id int,
	address,
	telephone,
	warehouseCode string,
	localityId int,
) (domain.Warehouse, error) {

	warehouse, err := r.GetById(ctx, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
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

	result, err := r.db.ExecContext(
		ctx,
		sqlUpdate,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.WarehouseCode,
		&warehouse.LocalityId,
		id,
	)

	affected, err := result.RowsAffected()
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Warehouse{}, err
	}
	log.Println(affected)

	return warehouse, nil
}

func (r *repository) Delete(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, sqlDelete, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), "warehouse not found")
		return fmt.Errorf("warehouse with id %d not found", id)
	}
	return
}
