package products

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Create(arg Product) (Product, error)
	LastID() (int, error)
	Update(arg Product) (Product, error)
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

func (r *repository) Create(arg Product) (Product, error) {
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return Product{}, err
	}

	newProduct := Product{
		arg.Id,
		arg.ProductCode,
		arg.Description,
		arg.Width,
		arg.Height,
		arg.Length,
		arg.NetWeight,
		arg.ExpirationRate,
		arg.RecommendedFreezingTemperature,
		arg.FreezingRate,
		arg.ProductTypeId,
		arg.SellerId,
	}
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

func (r *repository) updateProduct(product, arg Product) (Product, error) {
	if arg.ProductCode != "" {
		if r.productCodeExists(arg.ProductCode) {
			return Product{}, fmt.Errorf("the product with code \"%s\" already exists", arg.ProductCode)
		}

		product.ProductCode = arg.ProductCode
	}

	if arg.Description != "" {
		product.Description = arg.Description
	}

	if arg.Width != 0 {
		product.Width = arg.Width
	}

	if arg.Height != 0 {
		product.Height = arg.Height
	}

	if arg.Length != 0 {
		product.Length = arg.Length
	}

	if arg.NetWeight != 0 {
		product.NetWeight = arg.NetWeight
	}

	if arg.ExpirationRate != 0 {
		product.ExpirationRate = arg.ExpirationRate
	}

	if arg.RecommendedFreezingTemperature != 0 {
		product.RecommendedFreezingTemperature = arg.RecommendedFreezingTemperature
	}

	if arg.FreezingRate != 0 {
		product.FreezingRate = arg.FreezingRate
	}

	if arg.ProductTypeId != 0 {
		product.ProductTypeId = arg.ProductTypeId
	}

	if arg.SellerId != 0 {
		product.SellerId = arg.SellerId
	}

	return product, nil
}

func (r *repository) Update(arg Product) (Product, error) {
	var products []Product
	var updatedProduct Product
	updated := false

	if err := r.db.Read(&products); err != nil {
		return Product{}, err
	}

	for index, product := range products {
		if product.Id == arg.Id {
			patchedProduct, err := r.updateProduct(product, arg)
			if err != nil {
				return Product{}, err
			}

			products[index] = patchedProduct
			updated = true
			updatedProduct = product
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("product (%d) not found", arg.Id)
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
