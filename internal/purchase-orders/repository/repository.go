package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.PurchaseOrder, error) {
	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return nil, err
	}
	po := make([]domain.PurchaseOrder, 0)
	for rows.Next() {
		var p domain.PurchaseOrder
		err := rows.Scan(
			&p.Id,
			&p.OrderNumber,
			&p.OrderDate,
			&p.TrackingCode,
			&p.BuyerId,
			&p.CarrierId,
			&p.ProductRecordId,
			&p.OrderStatusId,
		)
		if err != nil {
			return nil, err
		}
		po = append(po, p)
	}
	return po, nil
}

func (r *repository) Create(ctx context.Context, OrderNumber string, OrderDate string, TrackingCode string, BuyerId int, CarrierId int, ProductRecordId int, OrderStatusId int) (*domain.PurchaseOrder, error) {
	result, err := r.db.ExecContext(ctx, queryCreate, OrderNumber, OrderDate, TrackingCode, BuyerId, CarrierId, ProductRecordId, OrderStatusId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error retrieving id %d", id)
	}

	i := domain.PurchaseOrder{Id: int(id),
		OrderNumber:     OrderNumber,
		OrderDate:       OrderDate,
		TrackingCode:    TrackingCode,
		BuyerId:         BuyerId,
		CarrierId:       CarrierId,
		ProductRecordId: ProductRecordId,
		OrderStatusId:   OrderStatusId,
	}
	return &i, nil
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
