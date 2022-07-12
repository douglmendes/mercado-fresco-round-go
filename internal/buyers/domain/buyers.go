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
	GetById(id int) (*Buyer, error)
	GetAll() ([]Buyer, error)
	GetOrdersByBuyers(ctx context.Context, id int) ([]OrdersByBuyers, error)
	Create(cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(id int) error
}

type Service interface {
	GetById(id int) (*Buyer, error)
	GetAll() ([]Buyer, error)
	GetOrdersByBuyers(ctx context.Context, id int) ([]OrdersByBuyers, error)
	Create(cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(id int) error
}
