package domain

import "context"

type PurchaseOrder struct {
	Id              int    `json:"id"`
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerId         int    `json:"buyer_id"`
	ProductRecordId int    `json:"product_record_id"`
	OrderStatusId   int    `json:"order_status_id"`
}

type Repository interface {
	Create(ctx context.Context, OrderNumber string, OrderDate string, TrackingCode string, BuyerId int, ProductRecordId int, OrderStatusId int) (*PurchaseOrder, error)
	GetAll(ctx context.Context) ([]PurchaseOrder, error)
}

type Service interface {
	Create(ctx context.Context, OrderNumber string, OrderDate string, TrackingCode string, BuyerId int, ProductRecordId int, OrderStatusId int) (*PurchaseOrder, error)
}
