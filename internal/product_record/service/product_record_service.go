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
	if productId != 0 {
		_, err := s.productRepository.GetById(productId)
		if err != nil {
			return []domain.ProductRecordCount{}, err
		}
	}

	productRecords, err := s.productRecordRepository.GetByProductId(productId)
	if err != nil {
		return []domain.ProductRecordCount{}, err
	}

	return productRecords, nil
}

func isValidDate(dateString string) bool {
	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return false
	}

	currentDate := time.Now()
	currentDateString := currentDate.String()[:10]
	parsedCurrentDate, _ := time.Parse(layout, currentDateString)

	diff := parsedDate.Sub(parsedCurrentDate)

	return diff >= 0
}

func (s service) Create(arg domain.ProductRecord) (
	domain.ProductRecord,
	error,
) {
	productRecord := domain.ProductRecord{}

	if !isValidDate(arg.LastUpdateDate) {
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
