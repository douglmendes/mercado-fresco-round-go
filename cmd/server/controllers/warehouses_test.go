package controllers

import (
	"encoding/json"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	mock_warehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWarehousesController_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//mockResponse := []warehouses.Warehouse{
	//	{
	//		1,
	//		"Monroe 860",
	//		"47470000",
	//		"TSFK",
	//		10,
	//		10,
	//	},
	//	{
	//		2,
	//		"Rua do Teste 2",
	//		"555555555",
	//		"JJJ",
	//		10,
	//		2,
	//	},
	//}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_warehouses.NewMockService(ctrl)
	handler := NewWarehouse(service).GetAll()

	appRouter := gin.New()
	appRouter.GET("/api/v1/warehouses/", handler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/", nil)
	rec := httptest.NewRecorder()

	appRouter.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	objRes := struct {
		Code int
		Data []warehouses.Warehouse
	}{}

	err := json.Unmarshal(rec.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)

	////ctrl.Call(http.MethodGet, "/warehouses/")
	//appRoute.GET("/warehouses", )

	//service.EXPECT().GetAll().Return(&warehouses.Warehouse{
	//	Id:                 1,
	//	Address:            "Rua teste",
	//	Telephone:          "9299292922",
	//	WarehouseCode:      "AAA",
	//	MinimunCapacity:    0,
	//	MinimunTemperature: 0,
	//}, nil)

	//r := appRoute.GET("/warehouses", handler)
	//req, _ := http.NewRequest("GET", "/warehouses/", nil)
	//w := httptest.NewRecorder()
	//appRoute.ServeHTTP(w, req)
	//
	//responseData, _ := ioutil.ReadAll(w.Body)
	//
	//assert.Equal(t, mockResponse, string(responseData))
	//assert.Equal(t, http.StatusOK, w.Code)

}
