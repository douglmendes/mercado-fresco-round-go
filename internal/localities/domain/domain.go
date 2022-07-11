package domain

import "context"

type Locality struct {
	Id           int    `json:"id"`
	ZipCode      string `json:"zip_code"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type SellersByLocality struct {
	LocalityId   int    `json:"id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type LocalityRepository interface {
	GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	GetBySellers(ctx context.Context, id int) ([]SellersByLocality, error)
	Create(ctx context.Context, zipCode, localityName, provinceName, countryName string) (Locality, error)
}

type LocalityService interface {
	GetBySellers(ctx context.Context, id int) ([]SellersByLocality, error)
	Create(ctx context.Context, zipCode, localityName, provinceName, countryName string) (Locality, error)
}
