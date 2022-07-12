package domain

import "context"

type InboudOrder struct {
	Id             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id"`
	ProductBatchId int    `json:"product_batch_id"`
	WarehouseId    int    `json:"warehouse_id"`
}

type EmployeeInboudOrder struct {
	Id               int64  `json:"id"`
	CardNumberId     string `json:"card_number_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	WarehouseId      int    `json:"warehouse_id"`
	InboudOrderCount int    `json:"inboud_order_count"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type Repository interface {
	Create(context.Context, string, string, int, int, int) (*InboudOrder, error)
	GetAll(context.Context) ([]InboudOrder, error)
	GetByEmployee(ctx context.Context, employee int64) ([]EmployeeInboudOrder, error)
}

type Service interface {
	Create(context.Context, string, string, int, int, int) (*InboudOrder, error)
	GetByEmployee(ctx context.Context, employee int64) ([]EmployeeInboudOrder, error)
}
