package controller

import (
	"bytes"
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	mockcarriers "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	carriersRelativePath = "/api/v1/carriers/"
)

var ctxMock = gomock.Any()

func callCarriersMock(t *testing.T) (*mockcarriers.MockCarrierService, *CarrierController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mockcarriers.NewMockCarrierService(ctrl)
	handler := NewCarries(service)
	api := gin.New()
	return service, handler, api
}

func TestCarrierController_Create_OK(t *testing.T) {
	carrier := domain.Carrier{
		Id:          1,
		Cid:         "DUDE",
		CompanyName: "Friends Group",
		Address:     "Rua dos Amigos",
		Telephone:   "24313243",
		LocalityId:  1,
	}

	service, handler, api := callCarriersMock(t)
	api.POST(carriersRelativePath, handler.Create())

	service.EXPECT().CreateCarrier(
		ctxMock,
		"DUDE",
		"Friends Group",
		"Rua dos Amigos",
		"24313243",
		1,
	).Return(carrier, nil)

	payload := `{"cid": "DUDE", "company_name": "Friends Group", "address": "Rua dos Amigos", "telephone": "24313243", "locality_id": 1}`
	req := httptest.NewRequest(http.MethodPost, carriersRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestCarrierController_Create_StatusUnprocessableEntity(t *testing.T) {
	_, handler, api := callCarriersMock(t)
	api.POST(carriersRelativePath, handler.Create())

	payload := `{"cid": "DUDE", "telephone": "24313243", "locality_id": 1}`
	req := httptest.NewRequest(http.MethodPost, carriersRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestCarrierController_Create_Conflict(t *testing.T) {
	service, handler, api := callCarriersMock(t)
	api.POST(carriersRelativePath, handler.Create())

	service.EXPECT().CreateCarrier(
		ctxMock,
		"DUDE",
		"Other Company",
		"Rua dos outros amigos",
		"1133111111",
		1,
	).Return(domain.Carrier{}, errors.New("already exists a carrier with this cid: DUDE"))

	payload := `{"cid": "DUDE", "company_name": "Other Company", "address": "Rua dos outros amigos", "telephone": "1133111111", "locality_id": 1}`
	req := httptest.NewRequest(http.MethodPost, carriersRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)

}
