package repository

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain"
	"github.com/stretchr/testify/assert"
)

var (
	firstPurchaseOrder = domain.PurchaseOrder{
		Id:              1,
		OrderNumber:     "xpto",
		OrderDate:       "2020-02-02",
		TrackingCode:    "ew23143543jn",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
	secondPurchaseOrder = domain.PurchaseOrder{
		Id:              2,
		OrderNumber:     "r2d2",
		OrderDate:       "2020-02-02",
		TrackingCode:    "ks8347587345hd",
		BuyerId:         2,
		ProductRecordId: 2,
		OrderStatusId:   2,
	}
	emptyPurchaseOrder *domain.PurchaseOrder
	allPurchaseOrders  = []domain.PurchaseOrder{
		firstPurchaseOrder,
		secondPurchaseOrder,
	}
	noPurchaseOrder []domain.PurchaseOrder
	someError       = errors.New("some error")
)

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		checkResult func(t *testing.T, result []domain.PurchaseOrder, err error)
	}{
		{
			name: "OK",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"order_number",
					"order_date",
					"tracking_code",
					"buyer_id",
					"product_record_id",
					"order_status_id",
				}).AddRow(
					firstPurchaseOrder.Id,
					firstPurchaseOrder.OrderNumber,
					firstPurchaseOrder.OrderDate,
					firstPurchaseOrder.TrackingCode,
					firstPurchaseOrder.BuyerId,
					firstPurchaseOrder.ProductRecordId,
					firstPurchaseOrder.OrderStatusId,
				).AddRow(
					secondPurchaseOrder.Id,
					secondPurchaseOrder.OrderNumber,
					secondPurchaseOrder.OrderDate,
					secondPurchaseOrder.TrackingCode,
					secondPurchaseOrder.BuyerId,
					secondPurchaseOrder.ProductRecordId,
					secondPurchaseOrder.OrderStatusId,
				)

				mock.ExpectQuery(regexp.QuoteMeta(queryGetAll)).WillReturnRows(rows)
			},
			checkResult: func(t *testing.T, result []domain.PurchaseOrder, err error) {
				assert.NoError(t, err)
				assert.Equal(t, allPurchaseOrders, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func() {
				mock.ExpectQuery(regexp.QuoteMeta(queryGetAll)).WillReturnError(someError)
			},
			checkResult: func(t *testing.T, result []domain.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Equal(t, noPurchaseOrder, result)
			},
		},
		{
			name: "Fail_Scan",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"order_number",
				}).AddRow(
					firstPurchaseOrder.Id,
					firstPurchaseOrder.OrderNumber,
				).AddRow(
					secondPurchaseOrder.Id,
					secondPurchaseOrder.OrderNumber,
				)

				mock.ExpectQuery(regexp.QuoteMeta(queryGetAll)).WillReturnRows(rows)
			},
			checkResult: func(t *testing.T, result []domain.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Equal(t, noPurchaseOrder, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.GetAll(context.Background())

			testCase.checkResult(t, result, err)
		})
	}
}

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name          string
		buildStubs    func()
		purchaseOrder domain.PurchaseOrder
		checkResult   func(t *testing.T, result *domain.PurchaseOrder, err error)
	}{
		{
			name: "OK",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(queryCreate)).
					WithArgs(
						firstPurchaseOrder.OrderNumber,
						firstPurchaseOrder.OrderDate,
						firstPurchaseOrder.TrackingCode,
						firstPurchaseOrder.BuyerId,
						firstPurchaseOrder.ProductRecordId,
						firstPurchaseOrder.OrderStatusId,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			purchaseOrder: firstPurchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.NoError(t, err)
				assert.Equal(t, &firstPurchaseOrder, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(queryCreate)).
					WithArgs(
						firstPurchaseOrder.OrderNumber,
						firstPurchaseOrder.OrderDate,
						firstPurchaseOrder.TrackingCode,
						firstPurchaseOrder.BuyerId,
						firstPurchaseOrder.ProductRecordId,
						firstPurchaseOrder.OrderStatusId,
					).
					WillReturnError(someError)
			},
			purchaseOrder: firstPurchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Equal(t, emptyPurchaseOrder, result)
			},
		},
		{
			name: "Last Id Error",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(queryCreate)).
					WithArgs(
						firstPurchaseOrder.OrderNumber,
						firstPurchaseOrder.OrderDate,
						firstPurchaseOrder.TrackingCode,
						firstPurchaseOrder.BuyerId,
						firstPurchaseOrder.ProductRecordId,
						firstPurchaseOrder.OrderStatusId,
					).
					WillReturnResult(driver.ResultNoRows)
			},
			purchaseOrder: firstPurchaseOrder,
			checkResult: func(t *testing.T, result *domain.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Equal(t, fmt.Errorf("error retrieving id %d", 0), err)
				assert.Equal(t, emptyPurchaseOrder, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.Create(
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
