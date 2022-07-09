package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	productDomain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
)

type service struct {
	productRecordRepository domain.ProductRecordRepository
	productRepository       productDomain.ProductRepository
}

func NewProductRecordService(
	repository domain.ProductRecordRepository,
	productRepository productDomain.ProductRepository,
) domain.ProductRecordService {
	return &service{
		productRecordRepository: repository,
		productRepository:       productRepository,
	}
}

func (s service) GetByProductId(productId int) ([]domain.ProductRecordCount, error) {
	// TODO: implementation
	return []domain.ProductRecordCount{}, nil
}

func isValidDate(dateString string) (bool, error) {
	layout := "2020-02-20"

	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return false, err
	}

	currentDate := time.Now()

	return parsedDate.Sub(currentDate) < 0, nil
}

func (s service) Create(arg domain.ProductRecord) (
	domain.ProductRecord,
	error,
) {
	productRecord := domain.ProductRecord{}

	validDate, _ := isValidDate(arg.LastUpdateDate)
	if !validDate {
		return productRecord,
			errors.New("last update date must be valid date (ex.: 2020-02-20) and greater than or equal current date")
	}

	_, err := s.productRepository.GetById(arg.ProductId)
	if err != nil {
		return productRecord,
			fmt.Errorf("product with id (%v) not found", arg.ProductId)
	}

	productRecord, err = s.productRecordRepository.Create(arg)
	if err != nil {
		return productRecord,
			fmt.Errorf("failed to create product record")
	}

	return productRecord, nil
}
