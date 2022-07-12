package service

import (
	"context"
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	ONCE = 1
)

var (
	purchaseOrder = domain.PurchaseOrder{
		Id:              1,
		OrderNumber:     "xpto",
		OrderDate:       "2020-02-02",
		TrackingCode:    "ew23143543jn",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	emptyPurchaseOrder *domain.PurchaseOrder
	purchaseOrders     = []domain.PurchaseOrder{purchaseOrder}
	noPurchaseOrders   = []domain.PurchaseOrder{}
	someError          = errors.New("some error")
)

func callMock(t *testing.T) (
	*mock_domain.MockRepository,
	domain.Service,
	context.Context,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_domain.NewMockRepository(ctrl)
	service := NewService(repository)

	return repository, service, context.Background()
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(repository *mock_domain.MockRepository, ctx context.Context)
		purchaseOrder domain.PurchaseOrder
		checkResult   func(t *testing.T, result *domain.PurchaseOrder, err error)
	}{
		{
			name: "OK",
			buildStubs: func(repository *mock_domain.MockRepository, ctx context.Context) {
				repository.
					EXPECT().
					GetAll(ctx).
					Times(ONCE).
					Return(noPurchaseOrders, nil)

				repository.
					EXPECT().
					Create(
						ctx,
						purchaseOrder.OrderNumber,
						purchaseOrder.OrderDate,
						purchaseOrder.TrackingCode,
						purchaseOrder.BuyerId,
						purchaseOrder.ProductRecordId,
						purchaseOrder.OrderStatusId,
					).
					Times(ONCE).
					Return(&purchaseOrder, nil)
			},
			purchaseOrder: purchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.NoError(t, err)

				assert.Equal(t, &purchaseOrder, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func(repository *mock_domain.MockRepository, ctx context.Context) {
				repository.
					EXPECT().
					GetAll(ctx).
					Times(ONCE).
					Return(noPurchaseOrders, nil)

				repository.
					EXPECT().
					Create(
						ctx,
						purchaseOrder.OrderNumber,
						purchaseOrder.OrderDate,
						purchaseOrder.TrackingCode,
						purchaseOrder.BuyerId,
						purchaseOrder.ProductRecordId,
						purchaseOrder.OrderStatusId,
					).
					Times(ONCE).
					Return(emptyPurchaseOrder, someError)
			},
			purchaseOrder: purchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.Error(t, err)

				assert.Equal(t, emptyPurchaseOrder, result)
			},
		},
		{
			name: "Fail_GetAll",
			buildStubs: func(repository *mock_domain.MockRepository, ctx context.Context) {
				repository.
					EXPECT().
					GetAll(ctx).
					Times(ONCE).
					Return(noPurchaseOrders, someError)
			},
			purchaseOrder: purchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.Error(t, err)

				assert.Equal(t, emptyPurchaseOrder, result)
			},
		},
		{
			name: "Conflict",
			buildStubs: func(repository *mock_domain.MockRepository, ctx context.Context) {
				repository.
					EXPECT().
					GetAll(ctx).
					Times(ONCE).
					Return(purchaseOrders, nil)
			},
			purchaseOrder: purchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Equal(t, errors.New("order number already exists"), err)

				assert.Equal(t, emptyPurchaseOrder, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service, ctx := callMock(t)

			testCase.buildStubs(repository, ctx)

			result, err := service.Create(
				context.Background(),
				testCase.purchaseOrder.OrderNumber,
				testCase.purchaseOrder.OrderDate,
				testCase.purchaseOrder.TrackingCode,
				testCase.purchaseOrder.BuyerId,
				testCase.purchaseOrder.ProductRecordId,
				testCase.purchaseOrder.OrderStatusId,
			)

			testCase.checkResult(t, result, err)
		})
	}
}
