package test

import (
	"encoding/json"
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

const relativePath = "/api/v1/warehouses/"

func TestWarehousesController_GetAll(t *testing.T) {
	whList := []warehouses.Warehouse{
		{
			1,
			"Monroe 860",
			"47470000",
			"TSFK",
			10,
			10,
		},
		{
			2,
			"Rua do Teste 2",
			"555555555",
			"JJJ",
			10,
			2,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_warehouses.NewMockService(ctrl)
	handler := controllers.NewWarehouse(service)
	api := gin.New()

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
