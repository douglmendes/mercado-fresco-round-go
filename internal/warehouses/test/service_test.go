package test

import (
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses"
	mock_warehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const id = 1

func callMock(t *testing.T) (*mock_warehouses.MockRepository, warehouses.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(apiMock)
	return apiMock, service
}

func TestService_GetAll(t *testing.T) {
	wh := []warehouses.Warehouse{
		{
			1,
			"Rua Café Torrado",
			"918288888",
			"ABC",
			100,
			2,
		},
		{
			2,
			"Av. Santo Agostinho",
			"7777777777",
			"DEF",
			23,
			12,
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(wh, nil)

	result, err := service.GetAll()
	assert.Equal(t, len(result), len(wh))
	assert.Nil(t, err)
}

func TestService_GetAll_NOK(t *testing.T) {
	wList := make([]warehouses.Warehouse, 0)

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(wList, errors.New("erro"))

	w, err := service.GetAll()
	assert.Equal(t, wList, w, "empty list")
	assert.NotNil(t, err)

}

func TestService_GetById(t *testing.T) {
	wh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua 25 de Março",
		Telephone:          "9911100011",
		WarehouseCode:      "XYZ",
		MinimunCapacity:    2,
		MinimunTemperature: 1,
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(wh, nil)

	result, err := service.GetById(id)
	assert.Equal(t, result.Id, id)
	assert.Nil(t, err)

}

func TestService_GetById_NOK(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(warehouses.Warehouse{}, errors.New("warehouse not found"))

	_, err := service.GetById(id)
	assert.NotNil(t, err)
}

func TestDelete_OK(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().Delete(gomock.Eq(id)).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestService_Delete_NOK(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().Delete(gomock.Eq(id)).Return(errors.New("error id is not valid"))

	err := service.Delete(id)
	assert.NotNil(t, err)
}
