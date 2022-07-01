package domain

import "context"

type Warehouse struct {
	Id            int    `json:"id"`
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	WarehouseCode string `json:"warehouse_code"`
	LocalityId    int    `json:"locality_id"`
}

//go:generate mockgen -source=./warehouse.go -destination=./mock/warehouse_mock.go
type WarehouseService interface {
	Create(ctx context.Context, address, telephone, warehouseCode string, localityId int) (*Warehouse, error)
	GetAll(ctx context.Context) ([]Warehouse, error)
	GetById(ctx context.Context, id int) (Warehouse, error)
	Update(ctx context.Context, id int, address, telephone, warehouseCode string, localityId int) (Warehouse, error)
	Delete(ctx context.Context, id int) error
}

type WarehouseRepository interface {
	Create(ctx context.Context, address, telephone, warehouseCode string, localityId int) (Warehouse, error)
	GetAll(ctx context.Context) ([]Warehouse, error)
	GetById(ctx context.Context, id int) (Warehouse, error)
	Update(ctx context.Context, id int, address, telephone, warehouseCode string, localityId int) (Warehouse, error)
	Delete(ctx context.Context, id int) error
}
