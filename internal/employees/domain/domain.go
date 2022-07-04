package domain

type Employee struct {
	Id           int64  `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type Repository interface {
	GetAll() ([]Employee, error)
	GetById(id int64) (*Employee, error)
	Create(cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Delete(id int64) error
}

type Service interface {
	GetAll() ([]Employee, error)
	GetById(id int64) (*Employee, error)
	Create(cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*Employee, error)
	Delete(id int64) error
}
