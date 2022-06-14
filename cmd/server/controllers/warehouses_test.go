package controllers

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	mock_warehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestWarehousesController_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_warehouses.NewMockService(ctrl)
	NewWarehouse(service)

	service.EXPECT().GetAll().Return(&warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua teste",
		Telephone:          "9299292922",
		WarehouseCode:      "AAA",
		MinimunCapacity:    0,
		MinimunTemperature: 0,
	}, nil)

}
