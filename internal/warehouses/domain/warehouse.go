package domain

import "context"

type Warehouse struct {
	Id            int64  `json:"id"`
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	WarehouseCode string `json:"warehouse_code"`
	LocalityId    int64  `json:"locality_id"`
}

//go:generate mockgen -source=./warehouse.go -destination=./mock/warehouse_mock.go
type WarehouseService interface {
	Create(ctx context.Context, address, telephone, warehouseCode string, localityId int64) (*Warehouse, error)
	GetAll(ctx context.Context) ([]Warehouse, error)
	GetById(ctx context.Context, id int64) (Warehouse, error)
	Update(ctx context.Context, id int64, address, telephone, warehouseCode string, localityId int64) (Warehouse, error)
	Delete(ctx context.Context, id int64) error
}

type WarehouseRepository interface {
	Create(ctx context.Context, address, telephone, warehouseCode string, localityId int64) (Warehouse, error)
	LastID() (int64, error)
	GetAll(ctx context.Context) ([]Warehouse, error)
	GetById(ctx context.Context, id int64) (Warehouse, error)
	Update(ctx context.Context, id int64, address, telephone, warehouseCode string, localityId int64) (Warehouse, error)
	Delete(ctx context.Context, id int64) error
}
