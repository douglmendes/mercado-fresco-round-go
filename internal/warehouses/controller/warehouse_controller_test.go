package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	mockwarehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	relativePath       = "/api/v1/warehouses/"
	relativePathWithId = "/api/v1/warehouses/:id"
	idString           = "1"
	idNumber           = 1
	locality           = 101
)

var ctx = context.TODO()

func callWarehousesMock(t *testing.T) (*mockwarehouses.MockWarehouseService, *WarehousesController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mockwarehouses.NewMockWarehouseService(ctrl)
	handler := NewWarehouse(service)
	api := gin.New()
	return service, handler, api
}

func TestWarehousesController_GetAll(t *testing.T) {
	whList := []domain.Warehouse{
		{
			Id:            1,
			Address:       "Monroe 860",
			Telephone:     "47470000",
			WarehouseCode: "TSFK",
		},
		{
			Id:            2,
			Address:       "Rua do Teste 2",
			Telephone:     "555555555",
			WarehouseCode: "JJJ",
		},
	}

	service, handler, api := callWarehousesMock(t)

	api.GET(relativePath, handler.GetAll())

	service.EXPECT().GetAll().Return(whList, nil)

	req := httptest.NewRequest(http.MethodGet, relativePath, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	respExpect := struct{ Data []domain.Warehouse }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, whList[0].WarehouseCode, respExpect.Data[0].WarehouseCode)
}

func TestWarehousesController_GetAll_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)

	api.GET(relativePath, handler.GetAll())

	service.EXPECT().GetAll(context.Background()).Return([]domain.Warehouse{}, errors.New("error 404"))

	req := httptest.NewRequest(http.MethodGet, relativePath, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}

func TestWarehousesController_GetById(t *testing.T) {

	wh := domain.Warehouse{
		Id:            idNumber,
		Address:       "Rua Nova",
		Telephone:     "12121212",
		WarehouseCode: "GHI",
		LocalityId:    locality,
	}

	service, handler, api := callWarehousesMock(t)

	api.GET(relativePathWithId, handler.GetById())

	service.EXPECT().GetById(ctx, gomock.Eq(idNumber)).Return(wh, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/warehouses/%s", idString), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data domain.Warehouse }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, wh.WarehouseCode, respExpect.Data.WarehouseCode)
}

func TestWarehousesController_GetById_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.GET(relativePathWithId, handler.GetById())
	service.EXPECT().GetById(ctx, gomock.Eq(idNumber)).Return(domain.Warehouse{}, errors.New("warehouse not found"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/warehouses/%s", idString), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestWarehousesController_GetById_BadRequest(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.GET(relativePathWithId, handler.GetById())

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/warehouses/%s", "opsHere"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWarehousesController_Create(t *testing.T) {
	wh := domain.Warehouse{
		Id:            idNumber,
		Address:       "Rua 1",
		Telephone:     "555555555",
		WarehouseCode: "ZAQ",
		LocalityId:    locality,
	}

	service, handler, api := callWarehousesMock(t)
	api.POST(relativePath, handler.Create())

	service.EXPECT().Create(
		ctx,
		"Rua 1",
		"555555555",
		"ZAQ",
		locality,
	).Return(&wh, nil)

	payload := `{"address": "Rua 1","telephone": "555555555","warehouse_code": "ZAQ","locality_id": 101}`
	req := httptest.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestWarehousesController_Create_Conflict(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.POST(relativePath, handler.Create())

	service.EXPECT().Create(
		ctx,
		"Rua 1",
		"555555555",
		"ZAQ",
		locality,
	).Return(nil, errors.New("this warehouse already exists"))

	payload := `{"address": "Rua 1","telephone": "555555555","warehouse_code": "ZAQ", "locality_id": 101}`
	req := httptest.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestWarehousesController_Create_Fail(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.POST(relativePath, handler.Create())

	payload := `{"address": "Rua 1","locality_id": 101}`
	req := httptest.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestWarehousesController_Delete_OK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.DELETE(relativePathWithId, handler.Delete())

	service.EXPECT().Delete(ctx, gomock.Eq(idNumber)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/warehouses/%s", idString), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestWarehousesController_Delete_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.DELETE(relativePathWithId, handler.Delete())

	service.EXPECT().Delete(ctx, gomock.Eq(idNumber)).Return(errors.New("erro 404"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/warehouses/%s", idString), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestWarehousesController_Delete_BadRequest(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.DELETE(relativePathWithId, handler.Delete())

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/warehouses/%s", "cuidado-Mando"), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWarehousesController_Update(t *testing.T) {
	wh := domain.Warehouse{
		Id:            1,
		Address:       "Av. Estrela da Morte",
		Telephone:     "987654321",
		WarehouseCode: "LSW",
		LocalityId:    locality,
	}

	service, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	service.EXPECT().Update(
		ctx,
		idNumber,
		"Rua Sem Saida",
		"888888888",
		"LSW",
		locality,
	).Return(wh, nil)

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW", "locality_id": 101}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", idString),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestWarehousesController_Update_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	service.EXPECT().Update(
		ctx,
		idNumber,
		"Rua Sem Saida",
		"888888888",
		"LSW",
		locality,
	).Return(domain.Warehouse{}, errors.New("warehouse not found"))

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW", "locality_id": 101}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", idString),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestWarehousesController_Update_Fail(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	payload := `{"address": "Rua Sem Saida","telephone": 888888888,"warehouse_code": 25,"locality_id": "101"}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", idString),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestWarehousesController_Update_BadRequest(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW","locality_id": 101}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", "Valfenda"),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
