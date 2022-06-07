package products

import "fmt"

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error)
	Update(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error)
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

func (s service) Store(productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error) {
	lastId, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	products, err := s.repository.GetAll()
	if err != nil {
		return Product{}, err
	}

	for _, product := range products {
		if product.ProductCode == productCode {
			return Product{}, fmt.Errorf("the product with code \"%s\" already exists", productCode)
		}
	}

	product, err := s.repository.Store(lastId+1, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s service) Update(id int, productCode, description string, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate float64, productTypeId, sellerId int) (Product, error) {
	updatedProduct, err := s.repository.Update(id, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)
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
