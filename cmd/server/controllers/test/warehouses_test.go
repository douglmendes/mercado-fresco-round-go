package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	mock_warehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/mock"
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
	id                 = "1"
)

func callWarehousesMock(t *testing.T) (*mock_warehouses.MockService, *controllers.WarehousesController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_warehouses.NewMockService(ctrl)
	handler := controllers.NewWarehouse(service)
	api := gin.New()
	return service, handler, api
}

func TestWarehousesController_GetAll(t *testing.T) {
	whList := []warehouses.Warehouse{
		{
			Id:                 1,
			Address:            "Monroe 860",
			Telephone:          "47470000",
			WarehouseCode:      "TSFK",
			MinimunCapacity:    10,
			MinimunTemperature: 10,
		},
		{
			Id:                 2,
			Address:            "Rua do Teste 2",
			Telephone:          "555555555",
			WarehouseCode:      "JJJ",
			MinimunCapacity:    10,
			MinimunTemperature: 2,
		},
	}

	service, handler, api := callWarehousesMock(t)

	api.GET(relativePath, handler.GetAll())

	service.EXPECT().GetAll().Return(whList, nil)

	req := httptest.NewRequest(http.MethodGet, relativePath, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	respExpect := struct{ Data []warehouses.Warehouse }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, whList[0].WarehouseCode, respExpect.Data[0].WarehouseCode)
}

func TestWarehousesController_GetAll_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)

	api.GET(relativePath, handler.GetAll())

	service.EXPECT().GetAll().Return([]warehouses.Warehouse{}, errors.New("error 404"))

	req := httptest.NewRequest(http.MethodGet, relativePath, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}

func TestWarehousesController_GetById(t *testing.T) {

	wh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua Nova",
		Telephone:          "12121212",
		WarehouseCode:      "GHI",
		MinimunCapacity:    2,
		MinimunTemperature: 2,
	}

	service, handler, api := callWarehousesMock(t)

	api.GET(relativePathWithId, handler.GetById())

	service.EXPECT().GetById(gomock.Eq(1)).Return(wh, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/warehouses/%s", id), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data warehouses.Warehouse }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, wh.WarehouseCode, respExpect.Data.WarehouseCode)
}

func TestWarehousesController_GetById_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.GET(relativePathWithId, handler.GetById())
	service.EXPECT().GetById(gomock.Eq(1)).Return(warehouses.Warehouse{}, errors.New("warehouse not found"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/warehouses/%s", id), nil)
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
	wh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua 1",
		Telephone:          "555555555",
		WarehouseCode:      "ZAQ",
		MinimunCapacity:    8,
		MinimunTemperature: 9,
	}

	service, handler, api := callWarehousesMock(t)
	api.POST(relativePath, handler.Create())

	service.EXPECT().Create(
		"Rua 1",
		"555555555",
		"ZAQ",
		8,
		9,
	).Return(&wh, nil)

	payload := `{"address": "Rua 1","telephone": "555555555","warehouse_code": "ZAQ","minimun_capacity": 8, "minimun_temperature": 9}`
	req := httptest.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestWarehousesController_Create_Conflict(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.POST(relativePath, handler.Create())

	service.EXPECT().Create(
		"Rua 1",
		"555555555",
		"ZAQ",
		8,
		9,
	).Return(nil, errors.New("this warehouse already exists"))

	payload := `{"address": "Rua 1","telephone": "555555555","warehouse_code": "ZAQ","minimun_capacity": 8, "minimun_temperature": 9}`
	req := httptest.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestWarehousesController_Create_Fail(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.POST(relativePath, handler.Create())

	payload := `{"address": "Rua 1","minimun_capacity": 8, "minimun_temperature": 9}`
	req := httptest.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(payload)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestWarehousesController_Delete_OK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.DELETE(relativePathWithId, handler.Delete())

	service.EXPECT().Delete(gomock.Eq(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/warehouses/%s", id), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestWarehousesController_Delete_NOK(t *testing.T) {
	service, handler, api := callWarehousesMock(t)
	api.DELETE(relativePathWithId, handler.Delete())

	service.EXPECT().Delete(gomock.Eq(1)).Return(errors.New("erro 404"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/warehouses/%s", id), nil)
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
	wh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Av. Estrela da Morte",
		Telephone:          "987654321",
		WarehouseCode:      "LSW",
		MinimunCapacity:    10,
		MinimunTemperature: 10,
	}

	service, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	service.EXPECT().Update(
		gomock.Eq(1),
		"Rua Sem Saida",
		"888888888",
		"LSW",
		8,
		9,
	).Return(wh, nil)

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW","minimun_capacity": 8, "minimun_temperature": 9}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", id),
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
		gomock.Eq(1),
		"Rua Sem Saida",
		"888888888",
		"LSW",
		8,
		9,
	).Return(warehouses.Warehouse{}, errors.New("warehouse not found"))

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW","minimun_capacity": 8, "minimun_temperature": 9}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", id),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestWarehousesController_Update_Fail(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	payload := `{"address": "Rua Sem Saida","telephone": 888888888,"warehouse_code": 25,"minimun_capacity": "Nice!!", "minimun_temperature": 9}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", id),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestWarehousesController_Update_BadRequest(t *testing.T) {
	_, handler, api := callWarehousesMock(t)
	api.PATCH(relativePathWithId, handler.Update())

	payload := `{"address": "Rua Sem Saida","telephone": "888888888","warehouse_code": "LSW","minimun_capacity": 8, "minimun_temperature": 9}`

	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/warehouses/%s", "Valfenda"),
		bytes.NewBuffer([]byte(payload)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
