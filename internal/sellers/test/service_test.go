package test

import (
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers"
	mock_sellers "github.com/douglmendes/mercado-fresco-round-go/internal/sellers/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = 1

func callMock(t *testing.T) (*mock_sellers.MockRepository, sellers.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_sellers.NewMockRepository(ctrl)
	service := sellers.NewService(apiMock)
	return apiMock, service
}

func TestService_GetAll_Ok(t *testing.T) {
	sl := []sellers.Seller{
		{
			ID:          1,
			Cid:         9,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Fantasma",
			Telephone:   "54327683",
		},
		{
			ID:          2,
			Cid:         10,
			CompanyName: "Mercado Pago",
			Address:     "Rua Cachoeira",
			Telephone:   "23098712",
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(sl, nil)

	result, err := service.GetAll()
	assert.Equal(t, len(result), len(sl))
	assert.Nil(t, err)

}

func TestService_GetAll_NOk(t *testing.T) {
	sList := make([]sellers.Seller, 0)

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(sList, errors.New("erro"))

	s, err := service.GetAll()
	assert.Equal(t, sList, s, "empty list")
	assert.NotNil(t, err)
}

func TestService_GetById_Ok(t *testing.T) {
	sl := sellers.Seller{
		ID:          1,
		Cid:         23,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(sl, nil)

	result, err := service.GetById(id)
	assert.Equal(t, result.ID, id)
	assert.Nil(t, err)

}

func TestService_GetById_NOk(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(sellers.Seller{}, errors.New("seller not found"))

	_, err := service.GetById(id)
	assert.NotNil(t, err)

}

func TestService_Delete_Ok(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().Delete(gomock.Eq(id)).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestService_Delete_NOk(t *testing.T) {
	apiMock, service := callMock(t)

	apiMock.EXPECT().Delete(gomock.Eq(id)).Return(errors.New("id is not valid"))

	err := service.Delete(id)
	assert.NotNil(t, err)
}

func TestService_Update_Ok(t *testing.T) {
	sl := sellers.Seller{
		ID:          1,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
	}

	slList := []sellers.Seller{
		{
			ID:          1,
			Cid:         22,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Meli",
			Telephone:   "34235432",
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(slList, nil)
	apiMock.EXPECT().Update(1, 20, "Mercado Livre", "Melicidade", "98787687").Return(sl, nil)

	result, err := service.Update(1, 20, "Mercado Livre", "Melicidade", "98787687")
	assert.Nil(t, err)
	assert.Equal(t, result, sl)
}

func TestService_Update_NOk(t *testing.T) {

	sl := sellers.Seller{}

	slList := []sellers.Seller{
		{
			ID:          1,
			Cid:         22,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Meli",
			Telephone:   "34235432",
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
		},
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(slList, nil)
	apiMock.EXPECT().Update(10, 20, "Mercado Livre", "Melicidade", "98787687").Return(sl, errors.New("seller 10 not found"))

	result, err := service.Update(10, 20, "Mercado Livre", "Melicidade", "98787687")
	assert.NotNil(t, err)
	assert.Equal(t, result, sl)
}
