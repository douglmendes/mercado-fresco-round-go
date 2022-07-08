package domain

type Buyer struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

//go:generate mockgen -source=./repository.go -destination=./mock/repository_mock.go
type Repository interface {
	GetById(id int) (*Buyer, error)
	GetAll() ([]Buyer, error)
	//LastID() (int, error)
	Create(cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(id int) error
}

type Service interface {
	GetById(id int) (*Buyer, error)
	GetAll() ([]Buyer, error)
	Create(cardNumberId, firstName, lastName string) (*Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (*Buyer, error)
	Delete(id int) error
}
