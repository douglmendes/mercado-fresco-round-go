package service

import (
	"context"
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain/mock"
	products_domain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	products_mock "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain/mock"
	sections_domain "github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
	sections_mock "github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain/mock"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

var (
	sampleBatch = domain.ProductBatch{
		Id:                 1,
		BatchNumber:        1,
		CurrentQuantity:    2,
		CurrentTemperature: 3,
		DueDate:            "2020-01-01",
		InitialQuantity:    4,
		ManufacturingDate:  "2020-01-01",
		ManufacturingHour:  5,
		MinimumTemperature: 6,
		ProductId:          7,
		SectionId:          8,
	}
	sampleRecord = domain.SectionRecords{
		SectionId:     1,
		SectionNumber: 3,
		ProductsCount: 5,
	}
	sampleProduct = products_domain.Product{
		Id:                             7,
		ProductCode:                    "xpto",
		Description:                    "description",
		Width:                          6.3,
		Height:                         2.3,
		Length:                         5.1,
		NetWeight:                      23.5,
		ExpirationRate:                 0.8,
		RecommendedFreezingTemperature: -4.3,
		FreezingRate:                   0.4,
		ProductTypeId:                  3,
		SellerId:                       5,
	}
	sampleSection = sections_domain.Section{
		Id:                 8,
		SectionNumber:      3,
		CurrentTemperature: 15,
		MinimumTemperature: 5,
		CurrentCapacity:    150,
		MinimumCapacity:    15,
		MaximumCapacity:    250,
		WarehouseId:        3,
		ProductTypeId:      3,
	}
)

func callMock(t *testing.T) (*mock_domain.MockProductBatchesRepository, *products_mock.MockProductRepository, *sections_mock.MockRepository, domain.ProductBatchesService) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiMock := mock_domain.NewMockProductBatchesRepository(ctrl)
	productsMock := products_mock.NewMockProductRepository(ctrl)
	sectionsMock := sections_mock.NewMockRepository(ctrl)

	service := NewService(apiMock, productsMock, sectionsMock)
	return apiMock, productsMock, sectionsMock, service
}

func TestService_Create_OK(t *testing.T) {
	api, prMock, scMock, service := callMock(t)

	api.EXPECT().GetAll(context.TODO()).Return([]domain.ProductBatch{}, nil)
	prMock.EXPECT().GetById(context.TODO(), sampleBatch.ProductId).Return(sampleProduct, nil)
	scMock.EXPECT().GetById(sampleBatch.SectionId).Return(&sampleSection, nil)
	api.EXPECT().Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8).Return(&sampleBatch, nil)

	result, err := service.Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8)
	assert.Equal(t, sampleBatch.Id, result.Id)
	assert.Nil(t, err)
}

func TestService_Create_Not_OK(t *testing.T) {
	api, _, _, service := callMock(t)

	api.EXPECT().GetAll(context.TODO()).Return(nil, errors.New("error"))
	api.EXPECT().Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8).Return(nil, errors.New("error"))

	_, err := service.Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8)
	assert.NotNil(t, err)
}

func TestService_Create_Conflict(t *testing.T) {
	api, _, _, service := callMock(t)

	api.EXPECT().GetAll(context.TODO()).Return([]domain.ProductBatch{sampleBatch}, nil)
	api.EXPECT().Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8).Return(nil, errors.New("conflict"))

	_, err := service.Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8)
	assert.NotNil(t, err)
}

func TestService_Create_Product_Not_Found(t *testing.T) {
	api, prMock, _, service := callMock(t)

	api.EXPECT().GetAll(context.TODO()).Return([]domain.ProductBatch{}, nil)
	prMock.EXPECT().GetById(context.TODO(), sampleBatch.ProductId).Return(products_domain.Product{}, errors.New("error"))
	api.EXPECT().Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8).Return(nil, errors.New("product not found"))

	_, err := service.Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8)
	assert.NotNil(t, err)
}

func TestService_Create_Section_Not_Found(t *testing.T) {
	api, prMock, scMock, service := callMock(t)

	api.EXPECT().GetAll(context.TODO()).Return([]domain.ProductBatch{}, nil)
	prMock.EXPECT().GetById(context.TODO(), sampleBatch.ProductId).Return(sampleProduct, nil)
	scMock.EXPECT().GetById(sampleBatch.SectionId).Return(nil, nil)
	api.EXPECT().Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8).Return(nil, errors.New("section not found"))

	_, err := service.Create(context.TODO(), 1, 2, 3, "2020-01-01", 4, "2020-01-01", 5, 6, 7, 8)
	assert.NotNil(t, err)
}
