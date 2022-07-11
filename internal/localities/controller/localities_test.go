package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	localityRelativePath       = "/api/v1/localities/"
	localityRelativePathReport = "/api/v1/localities/reportSellers"
	localityPathCarrierReport  = "/api/v1/localities/reportCarriers"
)

func callMockLocality(t *testing.T) (*mock_domain.MockLocalityService, *LocalityController, *gin.Engine) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_domain.NewMockLocalityService(ctrl)
	handler := NewLocality(service)
	api := gin.New()
	return service, handler, api
}

func TestLocalitiesController_GetBySellers_Ok(t *testing.T) {

	lc := []domain.SellersByLocality{
		{
			LocalityId:   1,
			LocalityName: "Lux",
			SellersCount: 2,
		},
	}

	service, handler, api := callMockLocality(t)

	api.GET(localityRelativePathReport, handler.GetBySellers())

	service.EXPECT().GetBySellers(gomock.Any(), gomock.Eq(1)).Return(lc, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/localities/reportSellers?id=1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data []domain.SellersByLocality }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, lc[len(lc)-1].LocalityName, respExpect.Data[len(respExpect.Data)-1].LocalityName)
}

func TestLocalitiesController_GetBySellers_NOk(t *testing.T) {

	service, handler, api := callMockLocality(t)

	api.GET(localityRelativePathReport, handler.GetBySellers())

	service.EXPECT().GetBySellers(gomock.Any(), gomock.Eq(1)).Return([]domain.SellersByLocality{}, errors.New("error"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/localities/reportSellers?id=1"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data []domain.SellersByLocality }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestLocalitiesController_Create_Ok(t *testing.T) {

	lc := domain.Locality{
		Id:           1,
		ZipCode:      "54365212",
		LocalityName: "Lux",
		ProvinceName: "Aracaju",
		CountryName:  "Brasil",
	}

	service, handler, api := callMockLocality(t)
	api.POST(localityRelativePath, handler.Create())

	service.EXPECT().Create(gomock.Any(), "54365212", "Lux", "Aracaju", "Brasil").Return(lc, nil)

	payload := `{"zip_code": "54365212", "locality_name": "Lux", "province_name": "Aracaju", "country_name": "Brasil"}`
	req := httptest.NewRequest(http.MethodPost, localityRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestLocalityController_Create_NoLocalityName(t *testing.T) {

	_, handler, api := callMockLocality(t)
	api.POST(localityRelativePath, handler.Create())

	payload := `{"zip_code": "54365212", "province_name": "Aracaju", "country_name": "Brasil"}`
	req := httptest.NewRequest(http.MethodPost, localityRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestLocalityController_Create_Conflit(t *testing.T) {

	service, handler, api := callMockLocality(t)
	api.POST(localityRelativePath, handler.Create())

	service.EXPECT().Create(gomock.Any(), "54365212", "Lux", "Aracaju", "Brasil").Return(domain.Locality{}, errors.New("this seller already exists"))
	payload := `{"zip_code": "54365212", "locality_name": "Lux", "province_name": "Aracaju", "country_name": "Brasil"}`
	req := httptest.NewRequest(http.MethodPost, localityRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestLocalitiesController_GetBySellers_InvalidId(t *testing.T) {

	lc := []domain.SellersByLocality{
		{
			LocalityId:   1,
			LocalityName: "Lux",
			SellersCount: 2,
		},
	}

	service, handler, api := callMockLocality(t)

	api.GET(localityRelativePathReport, handler.GetBySellers())

	service.EXPECT().GetBySellers(gomock.Any(), gomock.Eq(1)).Return(lc, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/localities/reportSellers?id=r"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data []domain.SellersByLocality }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestLocalityController_GetByCarriers_OK(t *testing.T) {
	carriersLocal := []domain.CarriersByLocality{
		{
			LocalityId:    2,
			LocalityName:  "Rohan",
			CarriersCount: 10,
		},
	}

	service, handler, api := callMockLocality(t)

	api.GET(localityPathCarrierReport, handler.GetByCarriers())

	service.EXPECT().GetByCarriers(gomock.Any(), gomock.Eq(2)).Return(carriersLocal, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/api/v1/localities/reportCarriers?id=2"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data []domain.CarriersByLocality }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, respExpect.Data[0].CarriersCount, carriersLocal[0].CarriersCount)
}
