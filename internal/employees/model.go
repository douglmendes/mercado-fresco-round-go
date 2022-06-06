package buyers

type Employee struct {
	Id           int    `json:"id"`
	CardNumberId int    `json:"card_number"`
	FirstName    string `json:"firs_name"`
	LastName     string `json:"last_name"`
	WarehouseId  int    `json:"warehouse_id"`
}
