package service

import (
	"context"
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = 1

func callMock(t *testing.T) (*mock_domain.MockLocalityRepository, domain.LocalityService) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_domain.NewMockLocalityRepository(ctrl)
	service := NewService(apiMock)
	return apiMock, service
}

func TestService_GetBySellers_Ok(t *testing.T) {

	lc := []domain.SellersByLocality{
		{
			LocalityId:   1,
			LocalityName: "Lux",
			SellersCount: 2,
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetBySellers(context.TODO(), 1).Return(lc, nil)

	result, err := service.GetBySellers(context.TODO(), 1)
	assert.Equal(t, 2, result[len(result)-1].SellersCount)
	assert.Nil(t, err)
}

func TestService_GetBySellers_NOk(t *testing.T) {

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetBySellers(context.TODO(), 1).Return([]domain.SellersByLocality{}, errors.New("seller not found"))

	_, err := service.GetBySellers(context.TODO(), id)
	assert.NotNil(t, err)
}

func TestCreate_Ok(t *testing.T) {

	lcList := []domain.Locality{
		{
			Id:           3,
			LocalityName: "Oliva",
			ProvinceName: "Curitiba",
			CountryName:  "Brasil",
		},
	}

	lc := domain.Locality{
		Id:           3,
		LocalityName: "Oliva",
		ProvinceName: "Curitiba",
		CountryName:  "Brasil",
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(lcList, nil)
	apiMock.EXPECT().Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil").Return(lc, nil)

	result, err := service.Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil")
	assert.Equal(t, result, lc)
	assert.Nil(t, err)
}

func TestCreate_NOk(t *testing.T) {

	lcList := []domain.Locality{
		{
			Id:           3,
			ZipCode:      "54365211",
			LocalityName: "Oliva",
			ProvinceName: "Curitiba",
			CountryName:  "Brasil",
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(lcList, nil)
	apiMock.EXPECT().Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil").Return(domain.Locality{}, errors.New("error"))

	_, err := service.Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil")
	assert.NotNil(t, err)
}

func TestCreate_Conflict(t *testing.T) {

	lcList := []domain.Locality{
		{
			Id:           3,
			ZipCode:      "54365212",
			LocalityName: "Oliva",
			ProvinceName: "Curitiba",
			CountryName:  "Brasil",
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(lcList, nil)
	apiMock.EXPECT().Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil").Return(domain.Locality{}, errors.New("error"))

	_, err := service.Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil")
	assert.NotNil(t, err)
}

func TestCreate_GetAll_NOk(t *testing.T) {

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return([]domain.Locality{}, errors.New("error"))

	_, err := service.Create(context.TODO(), "54365212", "Oliva", "Curitiba", "Brasil")
	assert.NotNil(t, err)
}

func TestService_GetByCarriers_OK(t *testing.T) {
	localCarrier := []domain.CarriersByLocality{
		{
			LocalityId:    1,
			LocalityName:  "Nexus",
			CarriersCount: 12,
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetByCarriers(context.TODO(), 1).Return(localCarrier, nil)

	result, err := service.GetByCarriers(context.TODO(), 1)
	assert.Equal(t, localCarrier[0].CarriersCount, result[0].CarriersCount)
	assert.Nil(t, err)

}

func TestService_GetByCarriers_NOK(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().GetByCarriers(context.TODO(), 1).Return([]domain.CarriersByLocality{}, errors.New("locality 1 not found"))

	_, err := service.GetByCarriers(context.TODO(), 1)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "locality 1 not found")
}
