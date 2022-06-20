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

func (s service) Update(arg Product) (Product, error) {
	updatedProduct, err := s.repository.Update(arg)
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
