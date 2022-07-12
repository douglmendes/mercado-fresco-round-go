package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	productRecordMockDomain "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain/mock"
	productDomain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	productMockDomain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	ONCE       = 1
	GET_ALL_ID = 0
)

var (
	emptyProductRecord = domain.ProductRecord{}
	productRecord      = domain.ProductRecord{
		Id:             1,
		LastUpdateDate: getCurrentDate(),
		PurchasePrice:  23.89,
		SalePrice:      43.99,
		ProductId:      1,
	}
	productRecordWithInvalidUpdateDate = domain.ProductRecord{
		Id:             1,
		LastUpdateDate: "2022/07/09",
		PurchasePrice:  23.89,
		SalePrice:      43.99,
		ProductId:      1,
	}
	firstProductRecordsCount = domain.ProductRecordCount{
		ProductId:    2,
		Description:  "Chocolate",
		RecordsCount: 3,
	}
	secondProductRecordsCount = domain.ProductRecordCount{
		ProductId:    4,
		Description:  "Ice Cream",
		RecordsCount: 1,
	}
	allProductRecordsCount = []domain.ProductRecordCount{
		firstProductRecordsCount,
		secondProductRecordsCount,
	}
	someProductRecordsCount = []domain.ProductRecordCount{
		firstProductRecordsCount,
	}
	emptyProductRecordsCount = []domain.ProductRecordCount{}
	noProductRecordsCount    = []domain.ProductRecordCount{}
	emptyProduct             = productDomain.Product{}
	someError                = errors.New("some error")
)

func getCurrentDate() string {
	currentDate := time.Now()

	return currentDate.String()[:10]
}

func callMock(t *testing.T) (
	*productRecordMockDomain.MockProductRecordRepository,
	*productMockDomain.MockProductRepository,
	domain.ProductRecordService,
	context.Context,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRecordRepository := productRecordMockDomain.NewMockProductRecordRepository(ctrl)
	productRepository := productMockDomain.NewMockProductRepository(ctrl)
	service := NewProductRecordService(productRecordRepository, productRepository)

	return productRecordRepository, productRepository, service, context.Background()
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name       string
		buildStubs func(
			productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
			productRepository *productMockDomain.MockProductRepository,
			ctx context.Context,
		)
		productRecord domain.ProductRecord
		checkResult   func(t *testing.T, result domain.ProductRecord, err error)
	}{
		{
			name: "OK",
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
				productRepository.
					EXPECT().
					GetById(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, nil)

				productRecordRepository.
					EXPECT().
					Create(ctx, productRecord).
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
				ctx context.Context,
			) {
				productRepository.
					EXPECT().
					GetById(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, nil)

				productRecordRepository.
					EXPECT().
					Create(ctx, productRecord).
					Times(ONCE).
					Return(emptyProductRecord, sql.ErrConnDone)
			},
			productRecord: productRecord,
			checkResult: func(t *testing.T, result domain.ProductRecord, err error) {
				assert.Equal(t, fmt.Errorf("failed to create product record"), err)

				assert.Equal(t, emptyProductRecord, result)
			},
		},
		{
			name: "Fail_Product_GetByID",
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
				productRepository.
					EXPECT().
					GetById(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, someError)
			},
			productRecord: productRecord,
			checkResult: func(t *testing.T, result domain.ProductRecord, err error) {
				assert.Equal(
					t,
					fmt.Errorf("product with id (%v) not found", productRecord.ProductId),
					err,
				)

				assert.Equal(t, emptyProductRecord, result)
			},
		},
		{
			name: "Last Update Date Error",
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
			},
			productRecord: productRecordWithInvalidUpdateDate,
			checkResult: func(t *testing.T, result domain.ProductRecord, err error) {
				assert.Equal(
					t,
					errors.New(
						"last update date must be valid date (ex.: 2020-02-20) and greater than or equal current date",
					),
					err,
				)

				assert.Equal(t, emptyProductRecord, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			productRecordRepository, productRepository, service, ctx := callMock(t)

			testCase.buildStubs(productRecordRepository, productRepository, ctx)

			result, err := service.Create(ctx, testCase.productRecord)
			testCase.checkResult(t, result, err)
		})
	}
}

func TestGetByProductId(t *testing.T) {
	testCases := []struct {
		name       string
		productId  int
		buildStubs func(
			productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
			productRepository *productMockDomain.MockProductRepository,
			ctx context.Context,
		)
		checkResult func(t *testing.T, result []domain.ProductRecordCount, err error)
	}{
		{
			name:      "OK_GetAll",
			productId: GET_ALL_ID,
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
				productRecordRepository.
					EXPECT().
					GetByProductId(ctx, GET_ALL_ID).
					Times(ONCE).
					Return(allProductRecordsCount, nil)
			},
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.NoError(t, err)

				assert.Equal(t, allProductRecordsCount, result)
			},
		},
		{
			name:      "OK_GetByProductId",
			productId: productRecord.ProductId,
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
				productRepository.
					EXPECT().
					GetById(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, nil)

				productRecordRepository.
					EXPECT().
					GetByProductId(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(someProductRecordsCount, nil)
			},
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.NoError(t, err)

				assert.Equal(t, someProductRecordsCount, result)
			},
		},
		{
			name:      "Fail_GetByProductId",
			productId: productRecord.ProductId,
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
				productRepository.
					EXPECT().
					GetById(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, someError)
			},
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.Error(t, err)

				assert.Equal(t, emptyProductRecordsCount, result)
			},
		},
		{
			name:      "Fail",
			productId: productRecord.ProductId,
			buildStubs: func(
				productRecordRepository *productRecordMockDomain.MockProductRecordRepository,
				productRepository *productMockDomain.MockProductRepository,
				ctx context.Context,
			) {
				productRepository.
					EXPECT().
					GetById(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProduct, nil)

				productRecordRepository.
					EXPECT().
					GetByProductId(ctx, productRecord.ProductId).
					Times(ONCE).
					Return(emptyProductRecordsCount, someError)
			},
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.Error(t, err)

				assert.Equal(t, emptyProductRecordsCount, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			productRecordRepository, productRepository, service, ctx := callMock(t)

			testCase.buildStubs(productRecordRepository, productRepository, ctx)

			result, err := service.GetByProductId(ctx, testCase.productId)
			testCase.checkResult(t, result, err)
		})
	}
}
