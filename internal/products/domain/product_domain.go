package domain

import "context"

type Product struct {
	Id                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, arg Product) (Product, error)
	Update(ctx context.Context, arg Product) (Product, error)
	Delete(ctx context.Context, id int) error
}

type ProductService interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, arg Product) (Product, error)
	Update(ctx context.Context, arg Product) (Product, error)
	Delete(ctx context.Context, id int) error
}
