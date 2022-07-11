package domain

import "context"

type Locality struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type SellersByLocality struct {
	LocalityId   int    `json:"id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}

type CarriersByLocality struct {
	LocalityId    int    `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount int    `json:"carriers_count"`
}

//go:generate mockgen -source=./domain.go -destination=./mock/domain_mock.go
type LocalityRepository interface {
	// GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	GetBySellers(ctx context.Context, id int) ([]SellersByLocality, error)
	GetByCarriers(ctx context.Context, id int) ([]CarriersByLocality, error)
	Create(ctx context.Context, localityName, provinceName, countryName string) (Locality, error)
	Update(ctx context.Context, id int, localityName, provinceName, countryName string) (Locality, error)
	Delete(ctx context.Context, id int) error
}

type LocalityService interface {
	// GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id int) (Locality, error)
	GetBySellers(ctx context.Context, id int) ([]SellersByLocality, error)
	GetByCarriers(ctx context.Context, id int) ([]CarriersByLocality, error)
	Create(ctx context.Context, localityName, provinceName, countryName string) (Locality, error)
	Update(ctx context.Context, id int, localityName, provinceName, countryName string) (Locality, error)
	Delete(ctx context.Context, id int) error
}
