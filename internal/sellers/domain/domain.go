package domain

import "context"

type Seller struct {
	ID          int    `json:"id"`
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int `json:"locality_id"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type Repository interface {
	GetAll(ctx context.Context) ([]Seller, error)
	GetById(ctx context.Context, id int) (Seller, error)
	Create(ctx context.Context, cid int, commpanyName, address, telephone string, localityId int) (Seller, error)
	Update(ctx context.Context, id, cid int, commpanyName, address, telephone string, localityId int) (Seller, error)
	Delete(ctx context.Context, id int) error
}

type Service interface {
	GetAll(ctx context.Context) ([]Seller, error)
	GetById(ctx context.Context, id int) (Seller, error)
	Create(ctx context.Context, cid int, commpanyName, address, telephone string, localityId int) (Seller, error)
	Update(ctx context.Context, id, cid int, companyname, address, telephone string, localityId int) (Seller, error)
	Delete(ctx context.Context, id int) error
}
