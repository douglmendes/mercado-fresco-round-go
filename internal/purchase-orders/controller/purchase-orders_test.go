package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/purchase-orders/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	PATH = "/api/v1/purchase-orders/"
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
	noPurchaseOrder domain.PurchaseOrder
	someError       = errors.New("some error")
)

type purchaseOrderResponseBody struct {
	Data  domain.PurchaseOrder `json:"data"`
	Error string               `json:"error"`
}

func callMock(t *testing.T) (
	*mock_domain.MockService,
	*PurchaseOrder,
	*gin.Engine,
	gomock.Matcher,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_domain.NewMockService(ctrl)
	handler := NewPurchaseOrders(service)
	api := gin.New()

	return service, handler, api, gomock.Any()
}

func TestPurchaseOrderController_Create(t *testing.T) {
	newPurchaseOrder := purchaseOrder
	newPurchaseOrder.Id = 0

	unprocessablePurchaseOrder := struct{}{}

	testCases := []struct {
		name        string
		payload     interface{}
		buildStubs  func(service *mock_domain.MockService, ctx gomock.Matcher)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			payload: newPurchaseOrder,
			buildStubs: func(service *mock_domain.MockService, ctx gomock.Matcher) {
				service.
					EXPECT().
					Create(
						ctx,
						newPurchaseOrder.OrderNumber,
						newPurchaseOrder.OrderDate,
						newPurchaseOrder.TrackingCode,
						newPurchaseOrder.BuyerId,
						newPurchaseOrder.ProductRecordId,
						newPurchaseOrder.OrderStatusId,
					).
					Times(ONCE).
					Return(&purchaseOrder, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, res.Code)

				body := purchaseOrderResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, purchaseOrder, body.Data)
				assert.Empty(t, body.Error)
			},
		},
		{
			name:       "Fail",
			payload:    unprocessablePurchaseOrder,
			buildStubs: func(service *mock_domain.MockService, ctx gomock.Matcher) {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnprocessableEntity, res.Code)

				body := purchaseOrderResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, noPurchaseOrder, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name:    "Conflict",
			payload: newPurchaseOrder,
			buildStubs: func(service *mock_domain.MockService, ctx gomock.Matcher) {
				service.
					EXPECT().
					Create(
						ctx,
						newPurchaseOrder.OrderNumber,
						newPurchaseOrder.OrderDate,
						newPurchaseOrder.TrackingCode,
						newPurchaseOrder.BuyerId,
						newPurchaseOrder.ProductRecordId,
						newPurchaseOrder.OrderStatusId,
					).
					Times(ONCE).
					Return(&noPurchaseOrder, someError)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusConflict, res.Code)

				body := purchaseOrderResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, noPurchaseOrder, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api, ctx := callMock(t)

			api.POST(PATH, handler.Create())

			testCase.buildStubs(service, ctx)

			payload, _ := json.Marshal(testCase.payload)
			req := httptest.NewRequest(
				http.MethodPost,
				PATH,
				bytes.NewBuffer(payload),
			)
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}
