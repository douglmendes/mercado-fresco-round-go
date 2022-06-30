package domain

type Warehouse struct {
	Id            int64  `json:"id"`
	Address       string `json:"address"`
	Telephone     string `json:"telephone"`
	WarehouseCode string `json:"warehouse_code"`
	LocalityId    int64  `json:"locality_id"`
}

//go:generate mockgen -source=./warehouse.go -destination=./mock/warehouse_mock.go
type WarehouseService interface {
	Create(address, telephone, warehouseCode string, localityId int64) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int64) (Warehouse, error)
	Update(id int64, address, telephone, warehouseCode string, localityId int64) (Warehouse, error)
	Delete(id int64) error
}

type WarehouseRepository interface {
	Create(address, telephone, warehouseCode string, localityId int64) (Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int64) (Warehouse, error)
	Update(id int64, address, telephone, warehouseCode string, localityId int64) (Warehouse, error)
	Delete(id int64) error
}
