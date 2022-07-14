package service

import (
	"context"
	"fmt"

	pbRepo "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
	productRepo "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	sectionRepo "github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
)

type service struct {
	productBatchesRepository pbRepo.ProductBatchesRepository
	productRepo              productRepo.ProductRepository
	sectionRepo              sectionRepo.Repository
}

func NewService(pbr pbRepo.ProductBatchesRepository, pr productRepo.ProductRepository, sr sectionRepo.Repository) pbRepo.ProductBatchesService {
	return &service{
		productBatchesRepository: pbr,
		productRepo:              pr,
		sectionRepo:              sr,
	}
}

func (s *service) Create(ctx context.Context, batchNumber, currentQuantity, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour, minimumTemperature, productId, sectionId int) (*pbRepo.ProductBatch, error) {
	productBatches, err := s.productBatchesRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, productBatch := range productBatches {
		if productBatch.BatchNumber == batchNumber {
			return nil, fmt.Errorf("a product batch with the batch_number %d already exists", batchNumber)
		}
	}

	product, err := s.productRepo.GetById(ctx, productId)
	if product.Id == 0 || err != nil {
		return nil, fmt.Errorf("product %d not found", productId)
	}

	section, err := s.sectionRepo.GetById(sectionId)
	if section == nil || section.Id == 0 || err != nil {
		return nil, fmt.Errorf("section %d not found", productId)
	}

	return s.productBatchesRepository.Create(ctx, batchNumber, currentQuantity, currentTemperature, dueDate, initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productId, sectionId)
}

func (s *service) GetBySectionId(ctx context.Context, sectionId int) ([]pbRepo.SectionRecords, error) {
	if sectionId != 0 {
		_, err := s.sectionRepo.GetById(sectionId)
		if err != nil {
			return nil, err
		}
	}

	records, err := s.productBatchesRepository.GetBySectionId(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	return records, nil
}
