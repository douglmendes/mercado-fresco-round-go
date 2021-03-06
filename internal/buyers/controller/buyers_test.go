package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain"
	mockbuyers "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	relativeBuyerPath       = "/api/v1/buyers/"
	relativePathBuyersId    = "/api/v1/buyers/:id"
	relativeOrdersBuyerPath = "/api/v1/buyers/reportPurchaseOrders"
)

func callBuyersMock(t *testing.T) (*mockbuyers.MockService, *BuyerController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mockbuyers.NewMockService(ctrl)
	handler := NewBuyer(service)
	api := gin.New()
	return service, handler, api
}

func TestBuyerController_GetAll_OK(t *testing.T) {
	buyersList := []domain.Buyer{
		{
			Id:           1,
			CardNumberId: "1234",
			FirstName:    "Mickey",
			LastName:     "Mouse",
		},
		{
			Id:           2,
			CardNumberId: "2222",
			FirstName:    "Pato",
			LastName:     "Donald",
		},
	}

	service, handler, api := callBuyersMock(t)
	req := httptest.NewRequest(http.MethodGet, relativeBuyerPath, nil)
	resp := httptest.NewRecorder()

	api.GET(relativeBuyerPath, handler.GetAll())

	service.EXPECT().GetAll(gomock.Any()).Return(buyersList, nil)

	api.ServeHTTP(resp, req)
	respExpect := struct{ Data []domain.Buyer }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, buyersList[1].CardNumberId, respExpect.Data[1].CardNumberId)

}

func TestBuyerController_GetAll_NotFound(t *testing.T) {
	service, handler, api := callBuyersMock(t)
	api.GET(relativeBuyerPath, handler.GetAll())
	service.EXPECT().GetAll(gomock.Any()).Return([]domain.Buyer{}, errors.New("error"))
	req := httptest.NewRequest(http.MethodGet, relativeBuyerPath, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_GetById(t *testing.T) {

	buyer := domain.Buyer{
		Id:           1,
		CardNumberId: "1234",
		FirstName:    "Mickey",
		LastName:     "Mouse",
	}

	service, handler, api := callBuyersMock(t)

	api.GET(relativePathBuyersId, handler.GetById())

	service.EXPECT().GetById(gomock.Any(), 1).Return(&buyer, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/buyers/%s", "1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data domain.Buyer }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, buyer.FirstName, respExpect.Data.FirstName)
}

func TestBuyersController_GetById_NOK(t *testing.T) {
	service, handler, api := callBuyersMock(t)
	api.GET(relativePathBuyersId, handler.GetById())
	service.EXPECT().GetById(gomock.Any(), 1).Return(&domain.Buyer{}, errors.New("buyer 1 not found"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/buyers/%s", "1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_GetById_BadRequest(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.GET(relativePathBuyersId, handler.GetById())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/buyers/%s", "hadouken"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyerController_Create_OK(t *testing.T) {
	buyer := domain.Buyer{
		Id:           1,
		CardNumberId: "1234",
		FirstName:    "Mickey",
		LastName:     "Mouse",
	}

	service, handler, api := callBuyersMock(t)
	api.POST(relativeBuyerPath, handler.Create())

	service.EXPECT().Create(gomock.Any(), "1234", "Mickey", "Mouse").Return(&buyer, nil)
	payload := `{"card_number_id": "1234", "first_name": "Mickey", "last_name": "Mouse"}`

	req := httptest.NewRequest(http.MethodPost, relativeBuyerPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestBuyerController_Create_Conflict(t *testing.T) {
	service, handler, api := callBuyersMock(t)
	api.POST(relativeBuyerPath, handler.Create())

	service.EXPECT().Create(gomock.Any(), "1234", "Mickey", "Mouse").Return(&domain.Buyer{}, errors.New("this card number id already exists"))
	payload := `{"card_number_id": "1234", "first_name": "Mickey", "last_name": "Mouse"}`

	req := httptest.NewRequest(http.MethodPost, relativeBuyerPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestBuyerController_Create_WithoutCardNumberID(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.POST(relativeBuyerPath, handler.Create())

	payload := `{"first_name": "Mickey", "last_name": "Mouse"}`

	req := httptest.NewRequest(http.MethodPost, relativeBuyerPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestBuyerController_Create_WithoutFirstName(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.POST(relativeBuyerPath, handler.Create())

	payload := `{"card_number_id": "1234", "last_name": "Mouse"}`

	req := httptest.NewRequest(http.MethodPost, relativeBuyerPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestBuyerController_Create_WithoutLastName(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.POST(relativeBuyerPath, handler.Create())

	payload := `{"card_number_id": "1234", "first_name": "Mickey"}`

	req := httptest.NewRequest(http.MethodPost, relativeBuyerPath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestBuyerController_Update_OK(t *testing.T) {
	buyer := domain.Buyer{
		Id:           1,
		CardNumberId: "1234",
		FirstName:    "Mickey",
		LastName:     "Mouse",
	}

	service, handler, api := callBuyersMock(t)
	api.PATCH(relativePathBuyersId, handler.Update())

	service.EXPECT().Update(
		gomock.Any(),
		1,
		"1234",
		"Silvio",
		"Santos",
	).Return(&buyer, nil)

	payload := `{"card_number_id": "1234", "first_name": "Silvio", "last_name": "Santos"}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/buyers/%s", "1"), bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuyerController_Update_BadRequest(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.PATCH(relativePathBuyersId, handler.Update())

	payload := `{"card_number_id": "1234", "first_name": "Silvio", "last_name": "Santos"}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/buyers/%s", "m??oi!"), bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestBuyerController_Update_UnprocessableEntity(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.PATCH(relativePathBuyersId, handler.Update())

	payload := `{"card_number_id": "1234", "first_name": 25, "last_name": 35}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/buyers/%s", "1"), bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestBuyerController_Update_NotFound(t *testing.T) {
	service, handler, api := callBuyersMock(t)
	api.PATCH(relativePathBuyersId, handler.Update())

	service.EXPECT().Update(
		gomock.Any(),
		1,
		"1234",
		"Silvio",
		"Santos",
	).Return(&domain.Buyer{}, errors.New("employee 1 not found"))

	payload := `{"card_number_id": "1234", "first_name": "Silvio", "last_name": "Santos"}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/buyers/%s", "1"), bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_Delete_OK(t *testing.T) {
	service, handler, api := callBuyersMock(t)
	api.DELETE(relativePathBuyersId, handler.Delete())

	service.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/buyers/%s", "1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestBuyersController_Delete_NOK(t *testing.T) {
	service, handler, api := callBuyersMock(t)
	api.DELETE(relativePathBuyersId, handler.Delete())

	service.EXPECT().Delete(gomock.Any(), 1).Return(errors.New("erro 404"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/buyers/%s", "1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_Delete_BadRequest(t *testing.T) {
	_, handler, api := callBuyersMock(t)
	api.DELETE(relativePathBuyersId, handler.Delete())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/buyers/%s", "hello"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyerController_GetOrdersByBuyers_OK(t *testing.T) {
	ordersBuyer := []domain.OrdersByBuyers{
		{
			Id:                  1,
			CardNumberId:        "1234",
			FirstName:           "Silvio",
			LastName:            "Santos",
			PurchaseOrdersCount: 2,
		},
	}

	service, handler, api := callBuyersMock(t)

	api.GET(relativeOrdersBuyerPath, handler.GetOrdersByBuyers())

	service.EXPECT().GetOrdersByBuyers(gomock.Any(), 1).Return(ordersBuyer, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/buyers/reportPurchaseOrders?id=1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data []domain.OrdersByBuyers }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, respExpect.Data[0].PurchaseOrdersCount, ordersBuyer[0].PurchaseOrdersCount)
}

func TestBuyersController_GetOrdersByBuyers_BadRequest(t *testing.T) {
	_, handler, api := callBuyersMock(t)

	api.GET(relativeOrdersBuyerPath, handler.GetOrdersByBuyers())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/buyers/reportPurchaseOrders?id=bom_dia"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestBuyersController_GetOrdersByBuyers_StatusNotFound(t *testing.T) {
	service, handler, api := callBuyersMock(t)

	api.GET(relativeOrdersBuyerPath, handler.GetOrdersByBuyers())
	service.EXPECT().GetOrdersByBuyers(gomock.Any(), gomock.Eq(1)).Return([]domain.OrdersByBuyers{}, errors.New("error"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/buyers/reportPurchaseOrders?id=1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
