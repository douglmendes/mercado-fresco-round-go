package domain

type Warehouse struct {
	Id                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

//go:generate mockgen -source=./warehouse.go -destination=./mock/warehouse_mock.go
type WarehouseService interface {
	Create(address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (*Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error)
	Delete(id int) error
}

type WarehouseRepository interface {
	Create(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error)
	LastID() (int, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	Update(id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature int) (Warehouse, error)
	Delete(id int) error
}
