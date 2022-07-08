package service

import "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"

type service struct {
	repository domain.ProductRecordRepository
}

func NewProductRecordService(repository domain.ProductRecordRepository) domain.ProductRecordService {
	return &service{repository}
}

func (s service) GetByProductId(productId int) ([]domain.ProductRecordCount, error) {
	// TODO: implementation
	return []domain.ProductRecordCount{}, nil
}

func (s service) Create(arg domain.ProductRecord) (domain.ProductRecord, error) {
	// TODO: implementation
	return domain.ProductRecord{}, nil
}
