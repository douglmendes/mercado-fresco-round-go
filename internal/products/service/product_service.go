package service

import (
	"context"
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
)

type service struct {
	repository domain.ProductRepository
}

func NewService(r domain.ProductRepository) domain.ProductService {
	return &service{repository: r}
}

func (s service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s service) GetById(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s service) Create(ctx context.Context, arg domain.Product) (domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return domain.Product{}, err
	}

	for _, product := range products {
		if product.ProductCode == arg.ProductCode {
			return domain.Product{}, fmt.Errorf("the product with code \"%s\" already exists", arg.ProductCode)
		}
	}

	product, err := s.repository.Create(ctx, arg)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s service) productCodeExists(ctx context.Context, arg domain.Product) (bool, error) {
	products, err := s.repository.GetAll(ctx)
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

func (s service) updateProduct(ctx context.Context, product, arg domain.Product) (
	domain.Product,
	error,
) {
	if arg.ProductCode != "" {
		validProductCode, err := s.productCodeExists(ctx, arg)
		if err != nil {
			return domain.Product{}, err
		}

		if validProductCode {
			return domain.Product{}, fmt.Errorf("the product with code \"%s\" already exists", arg.ProductCode)
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

func (s service) Update(ctx context.Context, arg domain.Product) (domain.Product, error) {
	foundProduct, err := s.repository.GetById(ctx, arg.Id)
	if err != nil {
		return domain.Product{}, err
	}

	updatedProduct, err := s.updateProduct(ctx, foundProduct, arg)
	if err != nil {
		return domain.Product{}, err
	}

	updatedProduct, err = s.repository.Update(ctx, updatedProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
