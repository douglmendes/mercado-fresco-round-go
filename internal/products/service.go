package products

import "fmt"

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Create(arg Product) (Product, error)
	Update(arg Product) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s service) GetById(id int) (Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s service) Create(arg Product) (Product, error) {
	lastId, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	products, err := s.repository.GetAll()
	if err != nil {
		return Product{}, err
	}

	for _, product := range products {
		if product.ProductCode == arg.ProductCode {
			return Product{}, fmt.Errorf("the product with code \"%s\" already exists", arg.ProductCode)
		}
	}

	arg.Id = lastId + 1

	product, err := s.repository.Create(arg)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s service) productCodeExists(arg Product) (bool, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return true, err
	}

	for _, product := range products {
		if product.Id != arg.Id && product.ProductCode == arg.ProductCode {
			return true, nil
		}
	}

	return false, nil
}

func (s service) updateProduct(product, arg Product) (Product, error) {
	if arg.ProductCode != "" {
		validProductCode, err := s.productCodeExists(arg)
		if err != nil {
			return Product{}, err
		}

		if validProductCode {
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

func (s service) Update(arg Product) (Product, error) {
	foundProduct, err := s.repository.GetById(arg.Id)
	if err != nil {
		return Product{}, err
	}

	updatedProduct, err := s.updateProduct(foundProduct, arg)
	if err != nil {
		return Product{}, err
	}

	updatedProduct, err = s.repository.Update(updatedProduct)
	if err != nil {
		return Product{}, err
	}

	return updatedProduct, nil
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
