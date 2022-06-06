package products

import (
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product

	err := r.db.Read(&products)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) Store(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error) {
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return Product{}, err
	}

	newProduct := Product{id, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId}
	products = append(products, newProduct)

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return newProduct, nil
}
