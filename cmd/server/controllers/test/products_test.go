package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	mock_products "github.com/douglmendes/mercado-fresco-round-go/internal/products/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	RELATIVE_PATH         = "/api/v1/products/"
	RELATIVE_PATH_WITH_ID = RELATIVE_PATH + ":id"
)

var (
	firstProduct = products.Product{
		Id:                             1,
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
	secondProduct = products.Product{
		Id:                             2,
		ProductCode:                    "xablau",
		Description:                    "description",
		Width:                          3.6,
		Height:                         3.2,
		Length:                         1.5,
		NetWeight:                      5.23,
		ExpirationRate:                 0.08,
		RecommendedFreezingTemperature: -3.4,
		FreezingRate:                   0.8,
		ProductTypeId:                  2,
		SellerId:                       3,
	}
	allProducts = []products.Product{
		firstProduct,
		secondProduct,
	}
)

type sliceResponseBody struct {
	Data  []products.Product `json:"data"`
	Error string             `json:"error"`
}

type productResponseBody struct {
	Data  products.Product `json:"data"`
	Error string           `json:"error"`
}

func callMock(t *testing.T) (
	*mock_products.MockService,
	*controllers.ProductController,
	*gin.Engine,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_products.NewMockService(ctrl)
	handler := controllers.NewProductController(service)
	api := gin.New()

	return service, handler, api
}

func TestProductController_GetAll(t *testing.T) {
	testCases := []struct {
		name        string
		buildStubs  func(service *mock_products.MockService)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(service *mock_products.MockService) {
				service.
					EXPECT().
					GetAll().
					Times(1).
					Return(allProducts, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
			},
		},
		{
			name: "Fail",
			buildStubs: func(service *mock_products.MockService) {
				service.
					EXPECT().
					GetAll().
					Times(1).
					Return([]products.Product{}, os.ErrClosed)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, res.Code)

				body := sliceResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name: "Empty",
			buildStubs: func(service *mock_products.MockService) {
				service.
					EXPECT().
					GetAll().
					Times(1).
					Return([]products.Product{}, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, res.Code)

				body := sliceResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Empty(t, body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callMock(t)

			api.GET(RELATIVE_PATH, handler.GetAll())

			testCase.buildStubs(service)

			req := httptest.NewRequest(http.MethodGet, RELATIVE_PATH, nil)
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}

func TestProductController_GetById(t *testing.T) {
	testCases := []struct {
		name        string
		productId   int
		buildStubs  func(service *mock_products.MockService)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productId: firstProduct.Id,
			buildStubs: func(service *mock_products.MockService) {
				service.
					EXPECT().
					GetById(firstProduct.Id).
					Times(1).
					Return(firstProduct, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, firstProduct, body.Data)
				assert.Empty(t, body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callMock(t)

			api.GET(RELATIVE_PATH_WITH_ID, handler.GetById())

			testCase.buildStubs(service)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s%d", RELATIVE_PATH, testCase.productId), nil)
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}
