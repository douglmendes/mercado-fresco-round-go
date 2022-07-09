package service

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	productRecordMockDomain "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain/mock"
	productDomain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	productMockDomain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const ONCE = 1

var (
	emptyProductRecord = domain.ProductRecord{}
	productRecord      = domain.ProductRecord{
		Id:             1,
		LastUpdateDate: "2022-07-09",
		PurchasePrice:  23.89,
		SalePrice:      43.99,
		ProductId:      1,
	}
	emptyProduct = productDomain.Product{}
)

func callMock(t *testing.T) (
	*productRecordMockDomain.MockProductRecordRepository,
	*productMockDomain.MockProductRepository,
	domain.ProductRecordService,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRecordRepository := productRecordMockDomain.NewMockProductRecordRepository(ctrl)
	productRepository := productMockDomain.NewMockProductRepository(ctrl)
	service := NewProductRecordService(productRecordRepository, productRepository)

	return productRecordRepository, productRepository, service
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name       string
		buildStubs func(
			productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
			productRepository *productMockDomain.MockProductRepository,
		)
		productRecord domain.ProductRecord
		checkResult   func(t *testing.T, result domain.ProductRecord, err error)
	}{
		{
			name: "OK",
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
			) {
				productRepository.
					EXPECT().
					GetById(productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, nil)

				productRecordRepository.
					EXPECT().
					Create(productRecord).
					Times(ONCE).
					Return(productRecord, nil)
			},
			productRecord: productRecord,
			checkResult: func(t *testing.T, result domain.ProductRecord, err error) {
				assert.NoError(t, err)

				assert.Equal(t, productRecord, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
			) {
				productRepository.
					EXPECT().
					GetById(productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, nil)

				productRecordRepository.
					EXPECT().
					Create(productRecord).
					Times(ONCE).
					Return(emptyProductRecord, sql.ErrConnDone)
			},
			productRecord: productRecord,
			checkResult: func(t *testing.T, result domain.ProductRecord, err error) {
				assert.Equal(t, fmt.Errorf("failed to create product record"), err)

				assert.Equal(t, emptyProductRecord, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			productRecordRepository, productRepository, service := callMock(t)

			testCase.buildStubs(productRecordRepository, productRepository)

			result, err := service.Create(testCase.productRecord)
			testCase.checkResult(t, result, err)
		})
	}
}
