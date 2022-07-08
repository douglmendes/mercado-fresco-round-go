package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	sellerRelativePath       = "/api/v1/sellers/"
	sellerRelativePathWithId = "/api/v1/sellers/:id"
	sellerId                 = "1"
)

func callMockSeller(t *testing.T) (*mock_domain.MockService, *SellerController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_domain.NewMockService(ctrl)
	handler := NewSeller(service)
	api := gin.New()
	return service, handler, api
}

func TestSellersController_GetAll(t *testing.T) {
	slList := []domain.Seller{
		{
			ID:          1,
			Cid:         22,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Meli",
			Telephone:   "34235432",
			LocalityId:  "1",
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  "1",
		},
	}

	service, handler, api := callMockSeller(t)

	api.GET(sellerRelativePath, handler.GetAll())

	req := httptest.NewRequest(http.MethodGet, sellerRelativePath, nil)
	resp := httptest.NewRecorder()

	// _, engine := gin.CreateTestContext(resp)

	service.EXPECT().GetAll(gomock.Any()).Return(slList, nil)

	// engine.GET(sellerRelativePath, handler.GetAll())

	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	respExpect := struct{ Data []domain.Seller }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, slList[0].Cid, respExpect.Data[0].Cid)
}

func TestSellersController_GetAll_NOk(t *testing.T) {
	service, handler, api := callMockSeller(t)

	api.GET(sellerRelativePath, handler.GetAll())

	service.EXPECT().GetAll(gomock.Any()).Return([]domain.Seller{}, errors.New("error 404"))

	req := httptest.NewRequest(http.MethodGet, sellerRelativePath, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSellersController_GetById(t *testing.T) {
	sl := domain.Seller{
		ID:          1,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
		LocalityId:  "1",
	}

	service, handler, api := callMockSeller(t)

	api.GET(sellerRelativePathWithId, handler.GetById())

	service.EXPECT().GetById(gomock.Any(), gomock.Eq(1)).Return(sl, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/sellers/%s", sellerId), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data domain.Seller }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, sl.Cid, respExpect.Data.Cid)
}

func TestSellersController_GetById_NOk(t *testing.T) {
	service, handler, api := callMockSeller(t)
	api.GET(sellerRelativePathWithId, handler.GetById())
	service.EXPECT().GetById(gomock.Any(), gomock.Eq(1)).Return(domain.Seller{}, errors.New("seller not found"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/sellers/%s", sellerId), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSellersController_GetById_BadRequest(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.GET(sellerRelativePathWithId, handler.GetById())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/sellers/%s", "test"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestSellersController_Create(t *testing.T) {
	sl := domain.Seller{
		ID:          1,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
		LocalityId:  "1",
	}

	service, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	service.EXPECT().Create(gomock.Any(), 20, "Mercado Livre", "Melicidade", "98787687", "1").Return(sl, nil)

	payload := `{"cid": 20, "company_name": "Mercado Livre", "address": "Melicidade", "telephone": "98787687", "locality_id": "1"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestSellersController_Create_Conflict(t *testing.T) {
	service, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	service.EXPECT().Create(gomock.Any(), 20, "Mercado Livre", "Melicidade", "98787687", "1").Return(domain.Seller{}, errors.New("this seller already exists"))
	payload := `{"cid": 20, "company_name": "Mercado Livre", "address": "Melicidade", "telephone": "98787687", "locality_id": "1"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestSellersController_Create_NoCid(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	payload := `{"company_name": "Mercado Livre", "address": "Melicidade", "telephone": "98787687", "locality_id": "1"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestSellersController_Create_NoCompanyName(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	payload := `{"cid": 20, "address": "Melicidade", "telephone": "98787687", "locality_id": "1"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestSellersController_Create_NoAddress(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	payload := `{"cid": 20, "company_name": "Mercado Livre", "telephone": "98787687", "locality_id": "1"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestSellersController_Create_NoTelephone(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	payload := `{"cid": 20, "company_name": "Mercado Livre", "address": "Melicidade", "locality_id": "1"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestSellersController_Create_NoLocalityId(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.POST(sellerRelativePath, handler.Create())

	payload := `{"cid": 20, "company_name": "Mercado Livre", "address": "Melicidade", "telephone": "98787687"}`
	req := httptest.NewRequest(http.MethodPost, sellerRelativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestSellersController_Update(t *testing.T) {
	sl := domain.Seller{
		ID:          1,
		Cid:         3,
		CompanyName: "Mercado Fresco",
		Address:     "Rua Bananeira",
		Telephone:   "34237123",
		LocalityId:  "1",
	}

	service, handler, api := callMockSeller(t)
	api.PATCH(sellerRelativePathWithId, handler.Update())

	service.EXPECT().Update(gomock.Any(), gomock.Eq(1), 3, "Mercado Pago", "Rua Bananeira, 130", "34237123", "1").Return(sl, nil)

	payload := `{"cid": 3, "company_name": "Mercado Pago", "address": "Rua Bananeira, 130", "telephone": "34237123", "locality_id": "1"}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/sellers/%s", sellerId), bytes.NewBuffer([]byte(payload)))

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestSellersController_Update_NOk(t *testing.T) {
	service, handler, api := callMockSeller(t)
	api.PATCH(sellerRelativePathWithId, handler.Update())

	service.EXPECT().Update(gomock.Any(), gomock.Eq(1), 3, "Mercado Pago", "Rua Bananeira, 130", "34237123", "1").
		Return(domain.Seller{}, errors.New("seller not found"))

	payload := `{"cid": 3, "company_name": "Mercado Pago", "address": "Rua Bananeira, 130", "telephone": "34237123", "locality_id": "1"}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/sellers/%s", sellerId), bytes.NewBuffer([]byte(payload)))

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSellersController_Update_Badrequest(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.PATCH(sellerRelativePathWithId, handler.Update())

	payload := `{"cid": 3, "company_name": "Mercado Pago", "address": "Rua Bananeira, 130", "telephone": "34237123", "locality_id": "1"}`

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/sellers/%s", "larousse"), bytes.NewBuffer([]byte(payload)))

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestSellersController_Delete_Ok(t *testing.T) {
	service, handler, api := callMockSeller(t)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/sellers/%s", sellerId), nil)
	resp := httptest.NewRecorder()

	service.EXPECT().Delete(gomock.Any(), gomock.Eq(1)).Return(nil)

	api.DELETE(sellerRelativePathWithId, handler.Delete())

	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestSellersController_Delete_NOk(t *testing.T) {
	service, handler, api := callMockSeller(t)
	api.DELETE(sellerRelativePathWithId, handler.Delete())

	service.EXPECT().Delete(gomock.Any(), gomock.Eq(1)).Return(errors.New("error 404"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/sellers/%s", sellerId), nil)
	resp := httptest.NewRecorder()

	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSellersController_Delete_BadRequest(t *testing.T) {
	_, handler, api := callMockSeller(t)
	api.DELETE(sellerRelativePathWithId, handler.Delete())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/sellers/%s", "enter-id"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
