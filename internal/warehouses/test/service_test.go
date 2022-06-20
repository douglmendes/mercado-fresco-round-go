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
			Id:                 1,
			Address:            "Rua 25 de Março",
			Telephone:          "9911100011",
			WarehouseCode:      "XYZ",
			MinimunCapacity:    2,
			MinimunTemperature: 1,
		},
		{
			Id:                 2,
			Address:            "Av. Santo Agostinho",
			Telephone:          "7777777777",
			WarehouseCode:      "DEF",
			MinimunCapacity:    23,
			MinimunTemperature: 12,
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

func TestService_Delete_OK(t *testing.T) {
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

func TestServiceCreate_OK(t *testing.T) {
	apiMock, service := callMock(t)

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

	whExpec := warehouses.Warehouse{
		Id:                 3,
		Address:            "Rua Nova",
		Telephone:          "12121212",
		WarehouseCode:      "GHI",
		MinimunCapacity:    2,
		MinimunTemperature: 2,
	}

	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(wh, nil)
	apiMock.EXPECT().Create(
		3,
		"Rua Nova",
		"12121212",
		"GHI",
		2,
		2,
	).Return(whExpec, nil)

	result, err := service.Create(
		"Rua Nova",
		"12121212",
		"GHI",
		2,
		2,
	)

	assert.Equal(t, result, &whExpec)
	assert.Nil(t, err)
}

func TestServiceCreate_NOK(t *testing.T) {
	apiMock, service := callMock(t)

	wh := []warehouses.Warehouse{
		{
			Id:                 1,
			Address:            "Rua Café Torrado",
			Telephone:          "918288888",
			WarehouseCode:      "ABC",
			MinimunCapacity:    100,
			MinimunTemperature: 2,
		},
		{
			Id:                 2,
			Address:            "Av. Santo Agostinho",
			Telephone:          "7777777777",
			WarehouseCode:      "DEF",
			MinimunCapacity:    23,
			MinimunTemperature: 12,
		},
	}

	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(wh, nil)
	apiMock.EXPECT().Create(
		3,
		"Rua Nova",
		"12121212",
		"ABC",
		2,
		2,
	).Return(warehouses.Warehouse{}, errors.New("this warehouse already exists"))

	_, err := service.Create(
		"Rua Nova",
		"12121212",
		"ABC",
		2,
		2,
	)

	assert.Equal(t, assert.NotNil(t, err), true)
	assert.EqualError(t, err, "this warehouse already exists")
}

func TestService_Update_WithTheSameWarehouseCode(t *testing.T) {
	apiMock, service := callMock(t)

	oldWh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua 25 de Março",
		Telephone:          "9911100011",
		WarehouseCode:      "XYZ",
		MinimunCapacity:    2,
		MinimunTemperature: 1,
	}

	newWh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua 23 de Março",
		Telephone:          "8899900099",
		WarehouseCode:      "XYZ",
		MinimunCapacity:    2,
		MinimunTemperature: 1,
	}

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(oldWh, nil)
	apiMock.EXPECT().Update(
		1,
		"Rua 23 de Março",
		"8899900099",
		"XYZ",
		2,
		1,
	).Return(newWh, nil)

	result, err := service.Update(
		1,
		"Rua 23 de Março",
		"8899900099",
		"XYZ",
		2,
		1,
	)

	assert.Equal(t, result, newWh)
	assert.Nil(t, err)
}

func TestService_Update_WithOtherWarehouseCode(t *testing.T) {
	apiMock, service := callMock(t)

	oldWh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua 25 de Março",
		Telephone:          "9911100011",
		WarehouseCode:      "XYZ",
		MinimunCapacity:    2,
		MinimunTemperature: 1,
	}

	newWh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Av. Paulista",
		Telephone:          "000000000",
		WarehouseCode:      "LLL",
		MinimunCapacity:    20,
		MinimunTemperature: 2,
	}

	whList := []warehouses.Warehouse{
		{
			Id:                 1,
			Address:            "Rua 25 de Março",
			Telephone:          "9911100011",
			WarehouseCode:      "XYZ",
			MinimunCapacity:    2,
			MinimunTemperature: 1,
		},
		{
			Id:                 2,
			Address:            "Av. Santo Agostinho",
			Telephone:          "7777777777",
			WarehouseCode:      "GNK",
			MinimunCapacity:    23,
			MinimunTemperature: 12,
		},
	}

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(oldWh, nil)
	apiMock.EXPECT().GetAll().Return(whList, nil)
	apiMock.EXPECT().Update(
		1,
		"Av. Paulista",
		"000000000",
		"LLL",
		20,
		2,
	).Return(newWh, nil)

	result, err := service.Update(
		1,
		"Av. Paulista",
		"000000000",
		"LLL",
		20,
		2,
	)

	assert.Equal(t, result, newWh)
	assert.Nil(t, err)
}

func TestService_Update_WithOtherWarehouseCode_NOK(t *testing.T) {
	apiMock, service := callMock(t)

	oldWh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua 25 de Março",
		Telephone:          "9911100011",
		WarehouseCode:      "XYZ",
		MinimunCapacity:    2,
		MinimunTemperature: 1,
	}

	//newWh := warehouses.Warehouse{
	//	Id:                 1,
	//	Address:            "Av. Paulista",
	//	Telephone:          "000000000",
	//	WarehouseCode:      "GNK",
	//	MinimunCapacity:    20,
	//	MinimunTemperature: 2,
	//}

	whList := []warehouses.Warehouse{
		{
			Id:                 1,
			Address:            "Rua 25 de Março",
			Telephone:          "9911100011",
			WarehouseCode:      "XYZ",
			MinimunCapacity:    2,
			MinimunTemperature: 1,
		},
		{
			Id:                 2,
			Address:            "Av. Santo Agostinho",
			Telephone:          "7777777777",
			WarehouseCode:      "GNK",
			MinimunCapacity:    23,
			MinimunTemperature: 12,
		},
	}

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(oldWh, nil)
	apiMock.EXPECT().GetAll().Return(whList, nil)

	result, err := service.Update(
		1,
		"Av. Paulista",
		"000000000",
		"GNK",
		20,
		2,
	)

	assert.Equal(t, result, warehouses.Warehouse{})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "this warehouse already exists")
}
