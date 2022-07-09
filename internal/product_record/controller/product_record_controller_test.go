package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	productRecordMockDomain "github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	PRODUCT_RECORDS_PATH        = "/api/v1/productRecords/"
	PRODUCT_REPORT_RECORDS_PATH = "/api/v1/products/reportRecords"
	GET_ALL_ID                  = 0
)

var (
	productRecord = domain.ProductRecord{
		Id:             1,
		LastUpdateDate: "2022-07-09",
		PurchasePrice:  22.2,
		SalePrice:      33.3,
		ProductId:      1,
	}
	emptyProductRecord       = domain.ProductRecord{}
	someError                = errors.New("some error")
	firstProductRecordsCount = domain.ProductRecordCount{
		ProductId:    1,
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
)

type productRecordResponseBody struct {
	Data  domain.ProductRecord `json:"data"`
	Error string               `json:"error"`
}

type sliceResponseBody struct {
	Data  []domain.ProductRecordCount `json:"data"`
	Error string                      `json:"error"`
}

func callProductsMock(t *testing.T) (
	*productRecordMockDomain.MockProductRecordService,
	*ProductRecordController,
	*gin.Engine,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := productRecordMockDomain.NewMockProductRecordService(ctrl)
	handler := NewProductRecordController(service)
	api := gin.New()

	return service, handler, api
}

func TestProductRecordController_Create(t *testing.T) {
	newProductRecord := productRecord
	newProductRecord.Id = 0

	productRecordWithMissingFields := struct{}{}

	testCases := []struct {
		name        string
		payload     interface{}
		buildStubs  func(service *productRecordMockDomain.MockProductRecordService)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			payload: newProductRecord,
			buildStubs: func(service *productRecordMockDomain.MockProductRecordService) {
				service.
					EXPECT().
					Create(newProductRecord).
					Times(1).
					Return(productRecord, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, res.Code)

				body := productRecordResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, productRecord, body.Data)
				assert.Empty(t, body.Error)
			},
		},
		{
			name:    "Fail",
			payload: productRecordWithMissingFields,
			buildStubs: func(service *productRecordMockDomain.MockProductRecordService) {
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnprocessableEntity, res.Code)

				body := productRecordResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, emptyProductRecord, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name:    "Conflict",
			payload: newProductRecord,
			buildStubs: func(service *productRecordMockDomain.MockProductRecordService) {
				service.
					EXPECT().
					Create(newProductRecord).
					Times(1).
					Return(emptyProductRecord, someError)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusConflict, res.Code)

				body := productRecordResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, emptyProductRecord, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callProductsMock(t)

			api.POST(PRODUCT_RECORDS_PATH, handler.Create())

			testCase.buildStubs(service)

			payload, _ := json.Marshal(testCase.payload)
			req := httptest.NewRequest(
				http.MethodPost,
				PRODUCT_RECORDS_PATH,
				bytes.NewBuffer(payload),
			)
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}

func TestProductRecordController_GetByProductId(t *testing.T) {
	testCases := []struct {
		name        string
		productId   int
		buildStubs  func(service *productRecordMockDomain.MockProductRecordService)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:      "OK_GetAll",
			productId: GET_ALL_ID,
			buildStubs: func(service *productRecordMockDomain.MockProductRecordService) {
				service.
					EXPECT().
					GetByProductId(GET_ALL_ID).
					Times(1).
					Return(allProductRecordsCount, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)

				body := sliceResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, allProductRecordsCount, body.Data)
				assert.Empty(t, body.Error)
			},
		},
		{
			name:      "OK_GetByProductId",
			productId: productRecord.ProductId,
			buildStubs: func(service *productRecordMockDomain.MockProductRecordService) {
				service.
					EXPECT().
					GetByProductId(productRecord.ProductId).
					Times(1).
					Return(someProductRecordsCount, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)

				body := sliceResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, someProductRecordsCount, body.Data)
				assert.Empty(t, body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callProductsMock(t)

			api.GET(PRODUCT_REPORT_RECORDS_PATH, handler.GetByProductId())

			testCase.buildStubs(service)

			var req *http.Request
			if testCase.productId != GET_ALL_ID {
				req = httptest.
					NewRequest(
						http.MethodGet,
						fmt.Sprintf("%s?id=%d", PRODUCT_REPORT_RECORDS_PATH, testCase.productId),
						nil,
					)
			} else {
				req = httptest.
					NewRequest(
						http.MethodGet,
						fmt.Sprintf("%s", PRODUCT_REPORT_RECORDS_PATH),
						nil,
					)
			}
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}
