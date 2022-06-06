package products

import (
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error)
	LastID() (int, error)
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

func (r *repository) LastID() (int, error) {
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}
