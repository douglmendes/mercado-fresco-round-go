package controller

import (
	"encoding/json"
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
	relativeBuyerPath = "/api/v1/buyers/"
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

	service.EXPECT().GetAll().Return(buyersList, nil)

	api.ServeHTTP(resp, req)
	respExpect := struct{ Data []domain.Buyer }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, buyersList[1].CardNumberId, respExpect.Data[1].CardNumberId)

}
