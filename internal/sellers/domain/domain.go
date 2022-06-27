package domain

type Seller struct {
	ID int `json:"id"`
	Cid int `json:"cid"`
	CompanyName string `json:"company_name"`
	Address string `json:"address"`
	Telephone string `json:"telephone"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type Repository interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Create(id, cid int, commpanyName, address, telephone string) (Seller, error)
	LastID() (int, error)
	Update(id, cid int, commpanyName, address, telephone string) (Seller, error)
	Delete(id int) error
}

type Service interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Create(cid int, commpanyName, address, telephone string) (Seller, error)
	Update(id, cid int, companyname, address, telephone string) (Seller, error)
	Delete(id int) error
}