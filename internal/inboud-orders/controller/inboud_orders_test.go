package controller

import (
	"bytes"
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	relativePathInboudOrders = "/api/v1/inboud-orders"
	target                   = "/api/v1/inboud-orders"
)

func callMock(t *testing.T) (*mock_domain.MockService, *InboudOrdersController) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_domain.NewMockService(ctrl)
	handler := NewInboudOrders(service)

	return service, handler
}

func TestController_GetByEmployee_Ok(t *testing.T) {
	ioReport := []domain.EmployeeInboudOrder{{
		Id:               1,
		CardNumberId:     "5555",
		FirstName:        "Douglas",
		LastName:         "Mendes",
		WarehouseId:      3,
		InboudOrderCount: 5,
	},
	}
	service, handler := callMock(t)
	api := gin.New()
	api.GET(target, handler.GetById())
	service.EXPECT().GetByEmployee(gomock.Any(), int64(1)).Return(ioReport, nil)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/inboud-orders?employee_id=1", nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestController_GetByEmployee_Nok(t *testing.T) {
	service, handler := callMock(t)
	api := gin.New()
	api.GET(relativePathInboudOrders, handler.GetById())
	service.EXPECT().GetByEmployee(gomock.Any(), int64(0)).Return(nil, errors.New("error 404"))

	req := httptest.NewRequest(http.MethodGet, relativePathInboudOrders, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestController_Create_Ok(t *testing.T) {
	io := domain.InboudOrder{
		Id:             4,
		OrderDate:      "1900-01-01",
		OrderNumber:    "order#3",
		EmployeeId:     4,
		ProductBatchId: 2,
		WarehouseId:    2,
	}
	service, handler := callMock(t)
	api := gin.New()
	api.POST(relativePathInboudOrders, handler.Create())

	service.EXPECT().Create(gomock.Any(), "1900-01-01", "order#3", 4, 2, 2).Return(&io, nil)

	body := `{"order_date": "1900-01-01","order_number": "order#3","employee_id":4,"product_batch_id": 2,"warehouse_id": 2}`
	req := httptest.NewRequest(http.MethodPost, relativePathInboudOrders, bytes.NewBuffer([]byte(body)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}
