package service

import (
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
	mock "github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func callMock(t *testing.T) (*mock.MockRepository, domain.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock.NewMockRepository(ctrl)
	service := NewService(apiMock)
	return apiMock, service
}

func TestService_Create_OK(t *testing.T) {
	api, service := callMock(t)
	api.EXPECT().GetAll().Return([]domain.Section{}, nil)

	newSection := domain.Section{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 15,
		MinimumTemperature: 5,
		CurrentCapacity:    150,
		MinimumCapacity:    15,
		MaximumCapacity:    250,
		WarehouseId:        3,
		ProductTypeId:      3,
	}

	api.EXPECT().Create(3, 15, 5, 150, 15, 250, 3, 3).Return(&newSection, nil)

	res, err := service.Create(3, 15, 5, 150, 15, 250, 3, 3)
	assert.Equal(t, res, &newSection)
	assert.Nil(t, err)
}

func TestService_Create_Conflict(t *testing.T) {
	api, service := callMock(t)
	api.EXPECT().GetAll().Return([]domain.Section{{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 15,
		MinimumTemperature: 5,
		CurrentCapacity:    150,
		MinimumCapacity:    15,
		MaximumCapacity:    250,
		WarehouseId:        3,
		ProductTypeId:      3,
	}}, nil)

	expectedError := domain.ErrorConflict{SectionNumber: 1}

	api.EXPECT().Create(1, 15, 5, 150, 15, 250, 1, 1).Return(nil, &expectedError)

	resp, err := service.Create(1, 15, 5, 150, 15, 250, 1, 1)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestService_Create_Data_Error(t *testing.T) {
	api, service := callMock(t)
	api.EXPECT().GetAll().Return(nil, errors.New("error"))

	api.EXPECT().Create(1, 15, 5, 150, 15, 250, 1, 1).Return(nil, errors.New("error"))

	resp, err := service.Create(1, 15, 5, 150, 15, 250, 1, 1)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestService_Find_All(t *testing.T) {
	api, service := callMock(t)

	db := []domain.Section{
		{
			Id:                 1,
			SectionNumber:      1,
			CurrentTemperature: 14,
			MinimumTemperature: 10,
			CurrentCapacity:    50,
			MinimumCapacity:    5,
			MaximumCapacity:    100,
			WarehouseId:        1,
			ProductTypeId:      1,
		},
		{
			Id:                 2,
			SectionNumber:      2,
			CurrentTemperature: 16,
			MinimumTemperature: 5,
			CurrentCapacity:    100,
			MinimumCapacity:    10,
			MaximumCapacity:    200,
			WarehouseId:        2,
			ProductTypeId:      2,
		},
	}

	api.EXPECT().GetAll().Return(db, nil)

	res, err := service.GetAll()
	assert.Equal(t, len(res), len(db))
	assert.Nil(t, err)
}

func TestService_Find_By_Id_Non_Existent(t *testing.T) {
	api, service := callMock(t)
	expectedError := domain.ErrorNotFound{Id: 3}

	api.EXPECT().GetById(3).Return(nil, &expectedError)

	res, err := service.GetById(3)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestService_Find_By_Id_Existent(t *testing.T) {
	api, service := callMock(t)

	foundSection := domain.Section{
		Id:                 2,
		SectionNumber:      2,
		CurrentTemperature: 16,
		MinimumTemperature: 5,
		CurrentCapacity:    100,
		MinimumCapacity:    10,
		MaximumCapacity:    200,
		WarehouseId:        2,
		ProductTypeId:      2,
	}

	api.EXPECT().GetById(2).Return(&foundSection, nil)

	res, err := service.GetById(2)
	assert.Equal(t, res, &foundSection)
	assert.Nil(t, err)
}

func TestService_Update_Existent(t *testing.T) {
	api, service := callMock(t)
	api.EXPECT().Exists(1).Return(nil)

	updatedSection := domain.Section{
		Id:                 1,
		SectionNumber:      1,
		MinimumTemperature: 10,
		CurrentCapacity:    50,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      1,
		CurrentTemperature: 15,
		MinimumCapacity:    15,
	}

	api.EXPECT().Update(1, map[string]int{"current_temperature": 15, "minimum_capacity": 15}).Return(&updatedSection, nil)

	res, err := service.Update(1, map[string]int{"current_temperature": 15, "minimum_capacity": 15})
	assert.Equal(t, res, &updatedSection)
	assert.Nil(t, err)
}

func TestService_Update_Section_Change(t *testing.T) {
	api, service := callMock(t)

	api.EXPECT().GetAll().Return([]domain.Section{
		{
			Id:                 1,
			SectionNumber:      1,
			MinimumTemperature: 10,
			CurrentCapacity:    50,
			MaximumCapacity:    100,
			WarehouseId:        1,
			ProductTypeId:      1,
			CurrentTemperature: 15,
			MinimumCapacity:    15,
		}}, nil)

	api.EXPECT().Exists(1).Return(nil)

	updatedSection := domain.Section{
		Id:                 1,
		SectionNumber:      15,
		MinimumTemperature: 10,
		CurrentCapacity:    50,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      1,
		CurrentTemperature: 15,
		MinimumCapacity:    15,
	}

	api.EXPECT().Update(1, map[string]int{"section_number": 15}).Return(&updatedSection, nil)

	res, err := service.Update(1, map[string]int{"section_number": 15})
	assert.Equal(t, res, &updatedSection)
	assert.Nil(t, err)
}

func TestService_Update_Non_Existent(t *testing.T) {
	api, service := callMock(t)
	expectedError := domain.ErrorNotFound{Id: 3}

	api.EXPECT().Exists(3).Return(&expectedError)

	res, err := service.Update(3, map[string]int{"current_temperature": 8})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestService_Update_Data_Error(t *testing.T) {
	api, service := callMock(t)

	api.EXPECT().Exists(1).Return(nil)
	api.EXPECT().GetAll().Return([]domain.Section{}, errors.New("error"))
	api.EXPECT().Update(1, map[string]int{"current_temperature": 8}).Return(nil, errors.New("error"))

	res, err := service.Update(1, map[string]int{"current_temperature": 8})
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestService_Update_Conflict(t *testing.T) {
	api, service := callMock(t)

	api.EXPECT().Exists(1).Return(nil)
	api.EXPECT().GetAll().Return([]domain.Section{}, nil)
	api.EXPECT().Update(1, map[string]int{"current_temperature": 8}).Return(nil, &domain.ErrorConflict{SectionNumber: 1})

	res, err := service.Update(1, map[string]int{"current_temperature": 8})
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualError(t, err, (&domain.ErrorConflict{SectionNumber: 1}).Error())
}

func TestService_Delete_Non_Existent(t *testing.T) {
	api, service := callMock(t)
	api.EXPECT().GetAll().Return([]domain.Section{}, nil)

	expectedError := domain.ErrorNotFound{Id: 3}

	api.EXPECT().Delete(3).Return(&expectedError)

	err := service.Delete(3)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestService_Delete_OK(t *testing.T) {
	api, service := callMock(t)

	api.EXPECT().Delete(1).Return(nil)

	err := service.Delete(1)
	assert.Nil(t, err)

	api.EXPECT().GetAll().Return([]domain.Section{}, nil)
}
