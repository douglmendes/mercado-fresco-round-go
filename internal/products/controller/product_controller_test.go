package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	RELATIVE_PATH         = "/api/v1/products/"
	RELATIVE_PATH_WITH_ID = RELATIVE_PATH + ":id"
	INVALID_ID            = 0
	WRONG_TYPE_ID         = "xpto"
)

var (
	emptyProduct = domain.Product{}
	firstProduct = domain.Product{
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
	secondProduct = domain.Product{
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
	allProducts = []domain.Product{
		firstProduct,
		secondProduct,
	}
)

type sliceResponseBody struct {
	Data  []domain.Product `json:"data"`
	Error string             `json:"error"`
}

type productResponseBody struct {
	Data  domain.Product `json:"data"`
	Error string           `json:"error"`
}

func callProductsMock(t *testing.T) (
	*mock_domain.MockService,
	*ProductController,
	*gin.Engine,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_domain.NewMockService(ctrl)
	handler := NewProductController(service)
	api := gin.New()

	return service, handler, api
}

func TestProductController_GetAll(t *testing.T) {
	testCases := []struct {
		name        string
		buildStubs  func(service *mock_domain.MockService)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(service *mock_domain.MockService) {
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
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, os.ErrClosed)
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
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, nil)
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
			service, handler, api := callProductsMock(t)

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
		name               string
		productId          int
		wrongTypeProductId string
		buildStubs         func(service *mock_domain.MockService)
		checkResult        func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productId: firstProduct.Id,
			buildStubs: func(service *mock_domain.MockService) {
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
		{
			name:      "NotFound",
			productId: INVALID_ID,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					GetById(INVALID_ID).
					Times(1).
					Return(domain.Product{}, fmt.Errorf("product (%d) not found", INVALID_ID))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Equal(t, fmt.Sprintf("product (%d) not found", INVALID_ID), body.Error)
			},
		},
		{
			name:               "InvalidId",
			wrongTypeProductId: WRONG_TYPE_ID,
			buildStubs:         func(service *mock_domain.MockService) {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Contains(t, body.Error, strconv.ErrSyntax.Error())
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callProductsMock(t)

			api.GET(RELATIVE_PATH_WITH_ID, handler.GetById())

			testCase.buildStubs(service)

			var req *http.Request
			if testCase.wrongTypeProductId != "" {
				req = httptest.
					NewRequest(
						http.MethodGet,
						fmt.Sprintf("%s%s", RELATIVE_PATH, testCase.wrongTypeProductId),
						nil,
					)
			} else {
				req = httptest.
					NewRequest(
						http.MethodGet,
						fmt.Sprintf("%s%d", RELATIVE_PATH, testCase.productId),
						nil,
					)
			}
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}

func TestProductController_Create(t *testing.T) {
	newProduct := firstProduct
	newProduct.Id = 0
	productWithMissingFields := struct {
		Id          int
		ProductCode string
		Description string
	}{
		Id:          1,
		ProductCode: "xpto",
		Description: "description",
	}

	testCases := []struct {
		name        string
		payload     interface{}
		buildStubs  func(service *mock_domain.MockService)
		checkResult func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			payload: newProduct,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Create(newProduct).
					Times(1).
					Return(firstProduct, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, firstProduct, body.Data)
				assert.Empty(t, body.Error)
			},
		},
		{
			name:       "Fail",
			payload:    productWithMissingFields,
			buildStubs: func(service *mock_domain.MockService) {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnprocessableEntity, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name:    "Conflict",
			payload: newProduct,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Create(newProduct).
					Times(1).
					Return(
						emptyProduct,
						fmt.Errorf(
							"the product with code \"%s\" already exists",
							firstProduct.ProductCode,
						),
					)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusConflict, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Equal(
					t,
					fmt.Sprintf(
						"the product with code \"%s\" already exists",
						firstProduct.ProductCode,
					),
					body.Error,
				)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callProductsMock(t)

			api.POST(RELATIVE_PATH, handler.Create())

			testCase.buildStubs(service)

			payload, _ := json.Marshal(testCase.payload)
			req := httptest.NewRequest(http.MethodPost, RELATIVE_PATH, bytes.NewBuffer(payload))
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}

func TestProductController_Update(t *testing.T) {
	updatedProduct := secondProduct
	updatedProduct.Id = firstProduct.Id
	updatedProductWithInvalidId := secondProduct
	updatedProductWithInvalidId.Id = INVALID_ID
	invalidProduct := struct {
		Width string
	}{
		Width: "invalid",
	}

	testCases := []struct {
		name               string
		payload            interface{}
		wrongTypeProductId string
		productId          int
		buildStubs         func(service *mock_domain.MockService)
		checkResult        func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			payload:   updatedProduct,
			productId: updatedProduct.Id,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Update(updatedProduct).
					Times(1).
					Return(updatedProduct, nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Equal(t, updatedProduct, body.Data)
				assert.Empty(t, body.Error)
			},
		},
		{
			name:       "Fail",
			payload:    invalidProduct,
			productId:  firstProduct.Id,
			buildStubs: func(service *mock_domain.MockService) {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnprocessableEntity, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name:               "BadRequest",
			payload:            updatedProduct,
			wrongTypeProductId: WRONG_TYPE_ID,
			buildStubs:         func(service *mock_domain.MockService) {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name:      "NotFound",
			payload:   updatedProduct,
			productId: INVALID_ID,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Update(updatedProductWithInvalidId).
					Times(1).
					Return(emptyProduct, fmt.Errorf("product (%d) not found", INVALID_ID))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Equal(t, fmt.Sprintf("product (%d) not found", INVALID_ID), body.Error)
			},
		},
		{
			name:      "InternalServerError",
			payload:   updatedProduct,
			productId: updatedProduct.Id,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Update(updatedProduct).
					Times(1).
					Return(emptyProduct, os.ErrClosed)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callProductsMock(t)

			api.PATCH(RELATIVE_PATH_WITH_ID, handler.Update())

			testCase.buildStubs(service)

			payload, _ := json.Marshal(testCase.payload)
			var req *http.Request
			if testCase.wrongTypeProductId != "" {
				req = httptest.NewRequest(
					http.MethodPatch,
					fmt.Sprintf("%s%s", RELATIVE_PATH, testCase.wrongTypeProductId),
					bytes.NewBuffer(payload),
				)
			} else {
				req = httptest.NewRequest(
					http.MethodPatch,
					fmt.Sprintf("%s%d", RELATIVE_PATH, testCase.productId),
					bytes.NewBuffer(payload),
				)
			}
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}

func TestProductController_Delete(t *testing.T) {
	testCases := []struct {
		name               string
		wrongTypeProductId string
		productId          int
		buildStubs         func(service *mock_domain.MockService)
		checkResult        func(t *testing.T, res *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productId: firstProduct.Id,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Delete(firstProduct.Id).
					Times(1).
					Return(nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Empty(t, body.Error)
			},
		},
		{
			name:               "Fail",
			wrongTypeProductId: WRONG_TYPE_ID,
			buildStubs:         func(service *mock_domain.MockService) {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.NotEmpty(t, body.Error)
			},
		},
		{
			name:      "NotFound",
			productId: INVALID_ID,
			buildStubs: func(service *mock_domain.MockService) {
				service.
					EXPECT().
					Delete(INVALID_ID).
					Times(1).
					Return(fmt.Errorf("product (%d) not found", INVALID_ID))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, res.Code)

				body := productResponseBody{}
				json.Unmarshal(res.Body.Bytes(), &body)

				assert.Empty(t, body.Data)
				assert.Equal(t, fmt.Sprintf("product (%d) not found", INVALID_ID), body.Error)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service, handler, api := callProductsMock(t)

			api.DELETE(RELATIVE_PATH_WITH_ID, handler.Delete())

			testCase.buildStubs(service)

			var req *http.Request
			if testCase.wrongTypeProductId != "" {
				req = httptest.NewRequest(
					http.MethodDelete,
					fmt.Sprintf("%s%s", RELATIVE_PATH, testCase.wrongTypeProductId),
					nil,
				)
			} else {
				req = httptest.NewRequest(
					http.MethodDelete,
					fmt.Sprintf("%s%d", RELATIVE_PATH, testCase.productId),
					nil,
				)
			}
			res := httptest.NewRecorder()
			api.ServeHTTP(res, req)

			testCase.checkResult(t, res)
		})
	}
}
