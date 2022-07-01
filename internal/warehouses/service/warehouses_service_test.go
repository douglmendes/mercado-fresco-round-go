package service

import (
	"context"
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	mockWarehouses "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const id = 1

func callMock(t *testing.T) (*mockWarehouses.MockWarehouseRepository, domain.WarehouseService, context.Context) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mockWarehouses.NewMockWarehouseRepository(ctrl)
	service := NewService(apiMock)
	ctxTest := context.Background()
	return apiMock, service, ctxTest
}

func TestService_GetAll(t *testing.T) {
	wh := []domain.Warehouse{
		{
			Id:            1,
			Address:       "Rua 25 de Março",
			Telephone:     "9911100011",
			WarehouseCode: "XYZ",
			LocalityId:    101,
		},
		{
			Id:            2,
			Address:       "Av. Santo Agostinho",
			Telephone:     "7777777777",
			WarehouseCode: "DEF",
			LocalityId:    202,
		},
	}

	apiMock, service, ctxTest := callMock(t)

	apiMock.EXPECT().GetAll(ctxTest).Return(wh, nil)

	result, err := service.GetAll(context.Background())
	assert.Equal(t, len(result), len(wh))
	assert.Nil(t, err)
}

func TestService_GetAll_NOK(t *testing.T) {
	wList := make([]domain.Warehouse, 0)

	apiMock, service, ctxTest := callMock(t)

	apiMock.EXPECT().GetAll(ctxTest).Return(wList, errors.New("erro"))

	w, err := service.GetAll(ctxTest)
	assert.Equal(t, wList, w, "empty list")
	assert.NotNil(t, err)

}

func TestService_GetById(t *testing.T) {
	wh := domain.Warehouse{
		Id:            1,
		Address:       "Rua 25 de Março",
		Telephone:     "9911100011",
		WarehouseCode: "XYZ",
		LocalityId:    101,
	}

	apiMock, service, ctxTest := callMock(t)

	apiMock.EXPECT().GetById(context.Background(), gomock.Eq(id)).Return(wh, nil)

	result, err := service.GetById(ctxTest, id)
	assert.Equal(t, result.Id, id)
	assert.Nil(t, err)

}

func TestService_GetById_NOK(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	apiMock.EXPECT().GetById(ctxTest, gomock.Eq(id)).Return(domain.Warehouse{}, errors.New("warehouse not found"))

	_, err := service.GetById(ctxTest, id)
	assert.NotNil(t, err)
}

func TestService_Delete_OK(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	apiMock.EXPECT().Delete(ctxTest, gomock.Eq(id)).Return(nil)

	err := service.Delete(ctxTest, id)
	assert.Nil(t, err)
}

func TestService_Delete_NOK(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	apiMock.EXPECT().Delete(ctxTest, gomock.Eq(id)).Return(errors.New("error id is not valid"))

	err := service.Delete(ctxTest, id)
	assert.NotNil(t, err)
}

func TestServiceCreate_OK(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	wh := []domain.Warehouse{
		{
			1,
			"Rua Café Torrado",
			"918288888",
			"ABC",
			101,
		},
		{
			2,
			"Av. Santo Agostinho",
			"7777777777",
			"DEF",
			102,
		},
	}

	whExpec := domain.Warehouse{
		Id:            3,
		Address:       "Rua Nova",
		Telephone:     "12121212",
		WarehouseCode: "GHI",
		LocalityId:    101,
	}

	apiMock.EXPECT().GetAll(ctxTest).Return(wh, nil)
	apiMock.EXPECT().Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"GHI",
		1,
	).Return(whExpec, nil)

	result, err := service.Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"GHI",
		1,
	)

	assert.Equal(t, result, &whExpec)
	assert.Nil(t, err)
}

func TestServiceCreate_Conflict(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	wh := []domain.Warehouse{
		{
			Id:            1,
			Address:       "Rua Café Torrado",
			Telephone:     "918288888",
			WarehouseCode: "ABC",
			LocalityId:    10,
		},
		{
			Id:            2,
			Address:       "Av. Santo Agostinho",
			Telephone:     "7777777777",
			WarehouseCode: "DEF",
			LocalityId:    11,
		},
	}

	apiMock.EXPECT().GetAll(ctxTest).Return(wh, nil)
	apiMock.EXPECT().Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"ABC",
		12,
	).Return(domain.Warehouse{}, errors.New("this warehouse already exists"))

	_, err := service.Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"ABC",
		12,
	)

	assert.Equal(t, assert.NotNil(t, err), true)
	assert.EqualError(t, err, "this warehouse already exists")
}

func TestService_Create_GetAll_Fail(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)
	apiMock.EXPECT().GetAll(ctxTest).Return([]domain.Warehouse{}, errors.New("error"))

	_, err := service.Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"ABC",
		10,
	)
	assert.NotNil(t, err)
}

func TestService_Create_NOK(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)
	wh := []domain.Warehouse{
		{
			1,
			"Rua Café Torrado",
			"918288888",
			"ABC",
			1,
		},
	}

	apiMock.EXPECT().GetAll(ctxTest).Return(wh, nil)
	apiMock.EXPECT().Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"GHI",
		1,
	).Return(domain.Warehouse{}, errors.New("error"))

	result, err := service.Create(
		ctxTest,
		"Rua Nova",
		"12121212",
		"GHI",
		1,
	)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestService_Update_WithTheSameWarehouseCode(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	oldWh := domain.Warehouse{
		Id:            id,
		Address:       "Rua 25 de Março",
		Telephone:     "9911100011",
		WarehouseCode: "XYZ",
		LocalityId:    1,
	}

	newWh := domain.Warehouse{
		Id:            id,
		Address:       "Rua 23 de Março",
		Telephone:     "8899900099",
		WarehouseCode: "XYZ",
		LocalityId:    1,
	}

	apiMock.EXPECT().GetById(ctxTest, gomock.Eq(id)).Return(oldWh, nil)
	apiMock.EXPECT().Update(
		ctxTest,
		id,
		"Rua 23 de Março",
		"8899900099",
		"XYZ",
		2,
	).Return(newWh, nil)

	result, err := service.Update(
		ctxTest,
		id,
		"Rua 23 de Março",
		"8899900099",
		"XYZ",
		2,
	)

	assert.Equal(t, result, newWh)
	assert.Nil(t, err)
}

func TestService_Update_WithOtherWarehouseCode(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	oldWh := domain.Warehouse{
		Id:            id,
		Address:       "Rua 25 de Março",
		Telephone:     "9911100011",
		WarehouseCode: "XYZ",
		LocalityId:    10,
	}

	newWh := domain.Warehouse{
		Id:            id,
		Address:       "Av. Paulista",
		Telephone:     "000000000",
		WarehouseCode: "LLL",
		LocalityId:    10,
	}

	whList := []domain.Warehouse{
		{
			Id:            1,
			Address:       "Rua 25 de Março",
			Telephone:     "9911100011",
			WarehouseCode: "XYZ",
			LocalityId:    11,
		},
		{
			Id:            2,
			Address:       "Av. Santo Agostinho",
			Telephone:     "7777777777",
			WarehouseCode: "GNK",
			LocalityId:    12,
		},
	}

	apiMock.EXPECT().GetById(ctxTest, gomock.Eq(id)).Return(oldWh, nil)
	apiMock.EXPECT().GetAll(ctxTest).Return(whList, nil)
	apiMock.EXPECT().Update(
		ctxTest,
		1,
		"Av. Paulista",
		"000000000",
		"LLL",
		10,
	).Return(newWh, nil)

	result, err := service.Update(
		ctxTest,
		1,
		"Av. Paulista",
		"000000000",
		"LLL",
		10,
	)

	assert.Equal(t, result, newWh)
	assert.Nil(t, err)
}

func TestService_Update_WithOtherWarehouseCode_NOK(t *testing.T) {
	apiMock, service, ctxTest := callMock(t)

	oldWh := domain.Warehouse{
		Id:            1,
		Address:       "Rua 25 de Março",
		Telephone:     "9911100011",
		WarehouseCode: "XYZ",
		LocalityId:    10,
	}

	whList := []domain.Warehouse{
		{
			Id:            1,
			Address:       "Rua 25 de Março",
			Telephone:     "9911100011",
			WarehouseCode: "XYZ",
			LocalityId:    10,
		},
		{
			Id:            2,
			Address:       "Av. Santo Agostinho",
			Telephone:     "7777777777",
			WarehouseCode: "GNK",
			LocalityId:    15,
		},
	}

	apiMock.EXPECT().GetById(ctxTest, gomock.Eq(id)).Return(oldWh, nil)
	apiMock.EXPECT().GetAll(ctxTest).Return(whList, nil)

	result, err := service.Update(
		ctxTest,
		1,
		"Av. Paulista",
		"000000000",
		"GNK",
		10,
	)

	assert.Equal(t, result, domain.Warehouse{})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "this warehouse already exists")
}
