package controller

import (
	"bytes"
	"encoding/json"
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
	PRODUCT_RECORDS_PATH = "/api/v1/productRecords/"
)

var productRecord = domain.ProductRecord{
	Id:             1,
	LastUpdateDate: "2022-07-09",
	PurchasePrice:  22.2,
	SalePrice:      33.3,
	ProductId:      1,
}

type productRecordResponseBody struct {
	Data  domain.ProductRecord `json:"data"`
	Error string               `json:"error"`
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

func TestProductController_Create(t *testing.T) {
	newProductRecord := productRecord
	newProductRecord.Id = 0

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
