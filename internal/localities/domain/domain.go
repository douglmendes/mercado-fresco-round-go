package domain

import "context"

type Locality struct {
	Id int `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName string `json:"country_name"`
}
//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type LocalityRepository interface {
	GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	Create(ctx context.Context, localityName, provinceName, countryName string) (Locality, error)
	Update(ctx context.Context, id int, localityName, provinceName, countryName string) (Locality, error)
	Delete(ctx context.Context, id int) error
}

type LocalityService interface {
	GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	Create(ctx context.Context, localityName, provinceName, countryName string) (Locality, error)
	Update(ctx context.Context, id int, localityName, provinceName, countryName string) (Locality, error)
	Delete(ctx context.Context, id int) error
}