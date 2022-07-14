package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	productDomain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
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

func (s service) GetByProductId(ctx context.Context, productId int) ([]domain.ProductRecordCount, error) {
	if productId != 0 {
		_, err := s.productRepository.GetById(ctx, productId)
		if err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())

			return []domain.ProductRecordCount{}, err
		}
	}

	productRecords, err := s.productRecordRepository.GetByProductId(ctx, productId)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())

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

func (s service) Create(ctx context.Context, arg domain.ProductRecord) (
	domain.ProductRecord,
	error,
) {
	productRecord := domain.ProductRecord{}

	if !isValidDate(arg.LastUpdateDate) {
		logger.Error(ctx, store.GetPathWithLine(), "invalid date")

		return productRecord,
			errors.New("last update date must be valid date (ex.: 2020-02-20) and greater than or equal current date")
	}

	_, err := s.productRepository.GetById(ctx, arg.ProductId)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())

		return productRecord,
			fmt.Errorf("product with id (%v) not found", arg.ProductId)
	}

	productRecord, err = s.productRecordRepository.Create(ctx, arg)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())

		return productRecord,
			fmt.Errorf("failed to create product record")
	}

	return productRecord, nil
}
