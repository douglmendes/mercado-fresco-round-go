package products

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error)
	LastID() (int, error)
	Update(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error)
	Delete(id int) error
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

func (r *repository) GetById(id int) (Product, error) {
	var products []Product

	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}

	return Product{}, fmt.Errorf("product (%d) not found", id)
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

func (r *repository) productCodeExists(productCode string) bool {
	var products []Product

	r.db.Read(&products)

	for _, product := range products {
		if product.ProductCode == productCode {
			return true
		}
	}

	return false
}

func (r *repository) updateProduct(product Product, id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error) {
	if productCode != "" {
		if r.productCodeExists(productCode) {
			return Product{}, fmt.Errorf("the product with code \"%s\" already exists", productCode)
		}

		product.ProductCode = productCode
	}

	if description != "" {
		product.Description = description
	}

	if width != 0 {
		product.Width = width
	}

	if height != 0 {
		product.Height = height
	}

	if length != 0 {
		product.Length = length
	}

	if netWeight != 0 {
		product.NetWeight = netWeight
	}

	if expirationRate != 0 {
		product.ExpirationRate = expirationRate
	}

	if recommendedFreezingTemperature != 0 {
		product.RecommendedFreezingTemperature = recommendedFreezingTemperature
	}

	if freezingRate != 0 {
		product.FreezingRate = freezingRate
	}

	if productTypeId != 0 {
		product.ProductTypeId = productTypeId
	}

	if sellerId != 0 {
		product.SellerId = sellerId
	}

	return product, nil
}

func (r *repository) Update(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error) {
	var products []Product
	var updatedProduct Product
	updated := false

	if err := r.db.Read(&products); err != nil {
		return Product{}, err
	}

	for index, product := range products {
		if product.Id == id {
			patchedProduct, err := r.updateProduct(product, id, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)
			if err != nil {
				return Product{}, err
			}

			products[index] = patchedProduct
			updated = true
			updatedProduct = product
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("product (%d) not found", id)
	}

	r.db.Write(products)

	return updatedProduct, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	foundIndex := 0
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return err
	}

	for index, product := range products {
		if product.Id == id {
			deleted = true
			foundIndex = index
		}
	}

	if !deleted {
		return fmt.Errorf("product (%d) not found", id)
	}

	products = append(products[:foundIndex], products[foundIndex+1:]...)

	r.db.Write(products)

	return nil
}
