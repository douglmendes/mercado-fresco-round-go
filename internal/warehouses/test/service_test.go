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

func TestService_GetAll(t *testing.T) {

	input := []warehouses.Warehouse{
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
			"Rua do doido 5",
			"555555555",
			"JJJ",
			10,
			2,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(apiMock)

	apiMock.EXPECT().GetAll().Return(input, nil)

	result, err := service.GetAll()

	assert.Equal(t, result, input)
	assert.Nil(t, err)
}

func TestService_GetAll_NOK(t *testing.T) {

	wList := make([]warehouses.Warehouse, 0)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(apiMock)

	apiMock.EXPECT().GetAll().Return(wList, errors.New("erro"))

	w, err := service.GetAll()
	assert.Equal(t, wList, w, "vazia")
	assert.NotNil(t, err)

}

func TestService_GetById(t *testing.T) {

	wh := warehouses.Warehouse{
		Id:                 1,
		Address:            "Rua teste 2",
		Telephone:          "9911100011",
		WarehouseCode:      "ASD",
		MinimunCapacity:    2,
		MinimunTemperature: 1,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(apiMock)

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(wh, nil)

	result, err := service.GetById(id)

	assert.Equal(t, result.Id, id)
	assert.Nil(t, err)

}

func TestService_GetById_NOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(apiMock)

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(warehouses.Warehouse{}, errors.New("meu erro"))

	_, err := service.GetById(id)
	assert.NotNil(t, err)
}
