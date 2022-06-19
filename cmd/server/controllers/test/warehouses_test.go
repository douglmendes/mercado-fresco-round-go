package test

import (
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
	relativePath        = "/api/v1/warehouses/"
	relativePathGetById = "/api/v1/warehouses/:id"
	id                  = "1"
)

func callMock(t *testing.T) (*mock_warehouses.MockService, *controllers.WarehousesController, *gin.Engine) {
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

	service, handler, api := callMock(t)

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
	service, handler, api := callMock(t)

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

	service, handler, api := callMock(t)

	api.GET(relativePathGetById, handler.GetById())

	service.EXPECT().GetById(gomock.Eq(1)).Return(wh, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/warehouses/%s", id), nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	respExpect := struct{ Data warehouses.Warehouse }{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respExpect)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, wh.WarehouseCode, respExpect.Data.WarehouseCode)
}
