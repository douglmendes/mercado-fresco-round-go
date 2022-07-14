package controller

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	getPath  = "/api/v1/sections/reportProducts"
	postPath = "/api/v1/productBatches/"
)

func callMock(t *testing.T) (*mock_domain.MockProductBatchesService, *ProductBatchesController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mock_domain.NewMockProductBatchesService(ctrl)
	handler := NewController(svc)
	api := gin.New()
	return svc, handler, api
}

func TestController_Create_Ok(t *testing.T) {
	result := domain.ProductBatch{
		Id:                 1,
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 1,
		DueDate:            "2020-01-01",
		InitialQuantity:    1,
		ManufacturingDate:  "2020-01-01",
		ManufacturingHour:  1,
		MinimumTemperature: 1,
		ProductId:          1,
		SectionId:          1,
	}

	service, handler, api := callMock(t)
	api.POST(postPath, handler.Create())

	service.EXPECT().Create(gomock.Any(), 1, 1, 1, "2020-01-01", 1, "2020-01-01", 1, 1, 1, 1).Return(&result, nil)

	payload := `{"batch_number":1,"current_quantity":1,"current_temperature":1,"due_date":"2020-01-01","initial_quantity":1,"manufacturing_date":"2020-01-01","manufacturing_hour":1,"minimum_temperature":1,"product_id":1,"section_id":1}`
	req := httptest.NewRequest(http.MethodPost, postPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestController_Create_Unprocessable(t *testing.T) {
	_, handler, api := callMock(t)
	api.POST(postPath, handler.Create())

	payload := `{"batch_number":1,"current_quantity":1,"current_temperature":1}`
	req := httptest.NewRequest(http.MethodPost, postPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestController_Create_Conflict(t *testing.T) {
	service, handler, api := callMock(t)
	api.POST(postPath, handler.Create())

	service.EXPECT().Create(gomock.Any(), 1, 1, 1, "2020-01-01", 1, "2020-01-01", 1, 1, 1, 1).Return(nil, errors.New("conflict"))

	payload := `{"batch_number":1,"current_quantity":1,"current_temperature":1,"due_date":"2020-01-01","initial_quantity":1,"manufacturing_date":"2020-01-01","manufacturing_hour":1,"minimum_temperature":1,"product_id":1,"section_id":1}`
	req := httptest.NewRequest(http.MethodPost, postPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}
