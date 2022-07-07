package domain

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

type Repository interface {
	Create(string, string, int, int, int) (*InboudOrder, error)
	GetAll() ([]InboudOrder, error)
	GetByEmployee(employee int64) ([]EmployeeInboudOrder, error)
}

type Service interface {
	Create(string, string, int, int, int) (*InboudOrder, error)
	GetByEmployee(employee int64) ([]EmployeeInboudOrder, error)
}
