package domain

import "context"

type PurchaseOrder struct {
	Id              int
	OrderNumber     string
	OrderDate       string
	TrackingCode    string
	BuyerId         int
	ProductRecordId int
	OrderStatusId   int
}

type Repository interface {
	Create(ctx context.Context, OrderNumber string, OrderDate string, TrackingCode string, BuyerId int, ProductRecordId int, OrderStatusId int) (*PurchaseOrder, error)
	GetAll(ctx context.Context) ([]PurchaseOrder, error)
}

type Service interface {
	Create(ctx context.Context, OrderNumber string, OrderDate string, TrackingCode string, BuyerId int, ProductRecordId int, OrderStatusId int) (*PurchaseOrder, error)
}
