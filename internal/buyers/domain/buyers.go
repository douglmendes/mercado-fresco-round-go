package domain

import "golang.org/x/net/context"

type Buyer struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type OrdersByBuyers struct {
	Id                  int    `json:"id"`
	CardNumberId        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

//go:generate mockgen -source=./buyers.go -destination=./mock/buyers_mock.go
type Repository interface {
	GetById(ctx context.Context, id int) (*Buyer, error)
	GetAll(ctx context.Context,) ([]Buyer, error)
	GetOrdersByBuyers(ctx context.Context, id int) ([]OrdersByBuyers, error)
	Create(ctx context.Context, cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(ctx context.Context, id int, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int) error
}

type Service interface {
	GetById(ctx context.Context, id int) (*Buyer, error)
	GetAll(ctx context.Context) ([]Buyer, error)
	GetOrdersByBuyers(ctx context.Context, id int) ([]OrdersByBuyers, error)
	Create(ctx context.Context, cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(ctx context.Context, id int, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(ctx context.Context, id int) error
}
