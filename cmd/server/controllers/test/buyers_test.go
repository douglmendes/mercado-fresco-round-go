package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	mock_buyers "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	relativePathBuyers       = "/api/v1/buyers/"
	relativePathBuyersWithId = "/api/v1/buyers/:id"
	idBuyers                 = "1"
)

func callMockBuyers(t *testing.T) (*mock_buyers.MockService, *controllers.BuyerController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_buyers.NewMockService(ctrl)
	handler := controllers.NewBuyer(service)
	api := gin.New()
	return service, handler, api
}

func TestBuyersController_GetAll(t *testing.T) {
	buyersList := []buyers.Buyer{
		{
			Id:           1,
			CardNumberId: "0001",
			FirstName:    "Jon",
			LastName:     "Snow",
		},
		{
			Id:           2,
			CardNumberId: "0002",
			FirstName:    "Daenerys",
			LastName:     "Targaryen",
		},
	}

	service, handler, api := callMockBuyers(t)

	api.GET(relativePathBuyers, handler.GetAll())

	service.EXPECT().GetAll().Return(buyersList, nil)

	req := httptest.NewRequest(http.MethodGet, relativePathBuyers, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	respExpect := struct{ Data []buyers.Buyer }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, buyersList[0].CardNumberId, respExpect.Data[0].CardNumberId)
}

func TestBuyersController_GetAll_NOK(t *testing.T) {
	service, handler, api := callMockBuyers(t)

	api.GET(relativePathBuyers, handler.GetAll())

	service.EXPECT().GetAll().Return([]buyers.Buyer{}, errors.New("error 404"))

	req := httptest.NewRequest(http.MethodGet, relativePathBuyers, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_GetById(t *testing.T) {

	buyer := buyers.Buyer{
		Id:           1,
		CardNumberId: "0001",
		FirstName:    "Daenerys",
		LastName:     "Targaryen",
	}

	service, handler, api := callMockBuyers(t)

	api.GET(relativePathBuyersWithId, handler.GetById())

	service.EXPECT().GetById(gomock.Eq(1)).Return(buyer, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/buyers/%s", idBuyers), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data buyers.Buyer }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, buyer.CardNumberId, respExpect.Data.CardNumberId)
}

func TestBuyersController_GetById_NOK(t *testing.T) {
	service, handler, api := callMockBuyers(t)
	api.GET(relativePathBuyersWithId, handler.GetById())
	service.EXPECT().GetById(gomock.Eq(1)).Return(buyers.Buyer{}, errors.New("warehouse not found"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/buyers/%s", idBuyers), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_Create(t *testing.T) {
	buyer := buyers.Buyer{
		Id:           1,
		CardNumberId: "0001",
		FirstName:    "Daenerys",
		LastName:     "Targaryen",
	}

	service, handler, api := callMockBuyers(t)
	api.POST(relativePathBuyers, handler.Create())

	service.EXPECT().Create(
		"0001",
		"Daenerys",
		"Targaryen",
	).Return(buyer, nil)

	payload := `{"card_number_id": "0001","first_name": "Daenerys","last_name": "Targaryen"}`
	req := httptest.NewRequest(http.MethodPost, relativePathBuyers, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestBuyersController_Create_Conflict(t *testing.T) {
	service, handler, api := callMockBuyers(t)
	api.POST(relativePathBuyers, handler.Create())

	service.EXPECT().Create(
		"0001",
		"Daenerys",
		"Targaryen",
	).Return(buyers.Buyer{}, errors.New("this buyer already exists"))

	payload := `{"card_number_id": "0001","first_name": "Daenerys","last_name": "Targaryen"}`
	req := httptest.NewRequest(http.MethodPost, relativePathBuyers, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestBuyersController_Delete_OK(t *testing.T) {
	service, handler, api := callMockBuyers(t)
	api.DELETE(relativePathBuyersWithId, handler.Delete())

	service.EXPECT().Delete(gomock.Eq(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/buyers/%s", idBuyers), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestBuyersController_Delete_NOK(t *testing.T) {
	service, handler, api := callMockBuyers(t)
	api.DELETE(relativePathBuyersWithId, handler.Delete())

	service.EXPECT().Delete(gomock.Eq(1)).Return(errors.New("erro 404"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/buyers/%s", idBuyers), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_Delete_BadRequest(t *testing.T) {
	_, handler, api := callMockBuyers(t)
	api.DELETE(relativePathBuyersWithId, handler.Delete())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/buyers/%s", "cuidado-Mando"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_Update(t *testing.T) {
	buyer := buyers.Buyer{
		Id:           1,
		CardNumberId: "0001",
		FirstName:    "Daenerys",
		LastName:     "Targaryen",
	}

	service, handler, api := callMockBuyers(t)
	api.PATCH(relativePathBuyersWithId, handler.Update())

	service.EXPECT().Update(
		gomock.Eq(1),
		"0001",
		"Mother of Dragon",
		"Breaker of Chains",
	).Return(buyer, nil)

	payload := `{"card_number_id": "0001","first_name": "Mother of Dragon","last_name": "Breaker of Chains"}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/buyers/%s", idBuyers),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuyersController_Update_NOK(t *testing.T) {
	service, handler, api := callMockBuyers(t)
	api.PATCH(relativePathBuyersWithId, handler.Update())

	service.EXPECT().Update(
		gomock.Eq(1),
		"0001",
		"Mother of Dragon",
		"Breaker of Chains",
	).Return(buyers.Buyer{}, errors.New("buyer not found"))

	payload := `{"card_number_id": "0001","first_name": "Mother of Dragon","last_name": "Breaker of Chains"}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/buyers/%s", idBuyers),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestBuyersController_Update_Fail(t *testing.T) {
	_, handler, api := callMockBuyers(t)
	api.PATCH(relativePathBuyersWithId, handler.Update())

	payload := `{"card_number_id": "0001","first_name": 007,"last_name": true}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/buyers/%s", idBuyers),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestBuyersController_Update_BadRequest(t *testing.T) {
	_, handler, api := callMockBuyers(t)
	api.PATCH(relativePathBuyersWithId, handler.Update())

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW","minimun_capacity": 8, "minimun_temperature": 9}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/buyers/%s", "straks00"),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
