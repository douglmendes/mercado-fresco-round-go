package domain

import "context"

type Employee struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type Repository interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Update(ctx context.Context, id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Delete(ctx context.Context, id int64) error
}

type Service interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int64) (*Employee, error)
	Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Update(ctx context.Context, id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Delete(ctx context.Context, id int64) error
}
