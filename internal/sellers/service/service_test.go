package service

import (
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = 1

func callMock(t *testing.T) (*mock_domain.MockRepository, domain.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_domain.NewMockRepository(ctrl)
	service := NewService(apiMock)
	return apiMock, service
}

func TestService_GetAll_Ok(t *testing.T) {
	sl := []domain.Seller{
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
	sList := make([]domain.Seller, 0)

	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(sList, errors.New("erro"))

	s, err := service.GetAll()
	assert.Equal(t, sList, s, "empty list")
	assert.NotNil(t, err)
}

func TestService_GetById_Ok(t *testing.T) {
	sl := domain.Seller{
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

	apiMock.EXPECT().GetById(gomock.Eq(id)).Return(domain.Seller{}, errors.New("seller not found"))

	_, err := service.GetById(id)
	assert.NotNil(t, err)

}

func TestCreate_Ok(t *testing.T) {

	slList := []domain.Seller{
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

	sl := domain.Seller{
		ID:          3,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
	}

	apiMock, service := callMock(t)

	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(slList, nil)
	apiMock.EXPECT().Create(3, 20, "Mercado Livre", "Melicidade", "98787687").Return(sl, nil)

	result, err := service.Create(20, "Mercado Livre", "Melicidade", "98787687")
	assert.Equal(t, result, sl)
	assert.Nil(t, err)
}

func TestCreate_NOk(t *testing.T) {

	slList := []domain.Seller{
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

	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(slList, nil)
	// apiMock.EXPECT().Create(3, 24, "Mercado Livre", "Melicidade", "98787687").Return(sellers.Seller{}, errors.New("this seller already exists"))

	_, err := service.Create(22, "Mercado Livre", "Melicidade", "98787687")
	assert.NotNil(t, err)
}

func TestService_Update_Ok(t *testing.T) {
	sl := domain.Seller{
		ID:          1,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
	}

	slList := []domain.Seller{
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

	sl := domain.Seller{}

	slList := []domain.Seller{
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

func TestService_Update_ExistentCid_NOk(t *testing.T) {

	sl := domain.Seller{}

	slList := []domain.Seller{
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

	result, err := service.Update(10, 22, "Mercado Livre", "Melicidade", "98787687")
	assert.NotNil(t, err)
	assert.Equal(t, result, sl)
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