package service

import (
	"context"
	"errors"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain/mock"
	locality "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	localityMock "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = 1

func callMock(t *testing.T) (*mock_domain.MockRepository, *localityMock.MockLocalityRepository, domain.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_domain.NewMockRepository(ctrl)
	apiLocalityMock := localityMock.NewMockLocalityRepository(ctrl)
	service := NewService(apiMock, apiLocalityMock)
	return apiMock, apiLocalityMock, service
}

func TestService_GetAll_Ok(t *testing.T) {
	sl := []domain.Seller{
		{
			ID:          1,
			Cid:         9,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Fantasma",
			Telephone:   "54327683",
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         10,
			CompanyName: "Mercado Pago",
			Address:     "Rua Cachoeira",
			Telephone:   "23098712",
			LocalityId:  1,
		},
	}

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(sl, nil)

	result, err := service.GetAll(context.TODO())
	assert.Equal(t, len(result), len(sl))
	assert.Nil(t, err)

}

func TestService_GetAll_NOk(t *testing.T) {
	sList := make([]domain.Seller, 0)

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(sList, errors.New("erro"))

	s, err := service.GetAll(context.TODO())
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
		LocalityId:  1,
	}

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetById(context.TODO(), gomock.Eq(id)).Return(sl, nil)

	result, err := service.GetById(context.TODO(), id)
	assert.Equal(t, result.ID, id)
	assert.Nil(t, err)

}

func TestService_GetById_NOk(t *testing.T) {
	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetById(context.TODO(), gomock.Eq(id)).Return(domain.Seller{}, errors.New("seller not found"))

	_, err := service.GetById(context.TODO(), id)
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
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  1,
		},
	}

	sl := domain.Seller{
		ID:          3,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
		LocalityId:  1,
	}

	apiMock, apiLocalityMock, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(slList, nil)
	apiLocalityMock.EXPECT().GetById(context.TODO(), gomock.Eq(id)).Return(locality.Locality{}, nil)
	apiMock.EXPECT().Create(context.TODO(), 20, "Mercado Livre", "Melicidade", "98787687", 1).Return(sl, nil)

	result, err := service.Create(context.TODO(), 20, "Mercado Livre", "Melicidade", "98787687", 1)
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
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  1,
		},
	}

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(slList, nil)
	apiMock.EXPECT().Create(3, 22, "Mercado Livre", "Melicidade", "98787687", 1).Return(domain.Seller{}, errors.New("this seller already exists"))

	_, err := service.Create(context.TODO(), 22, "Mercado Livre", "Melicidade", "98787687", 1)
	assert.NotNil(t, err)
}

func TestCreate_Locality_NOk(t *testing.T) {

	slList := []domain.Seller{
		{
			ID:          1,
			Cid:         22,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Meli",
			Telephone:   "34235432",
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  1,
		},
	}

	apiMock, apiLocalityMock, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(slList, nil)
	apiLocalityMock.EXPECT().GetById(context.TODO(), gomock.Eq(id)).Return(locality.Locality{}, errors.New("locality not found"))

	_, err := service.Create(context.TODO(), 20, "Mercado Livre", "Melicidade", "98787687", 1)
	assert.NotNil(t, err)
}

func TestService_Update_Ok(t *testing.T) {
	sl := domain.Seller{
		ID:          1,
		Cid:         20,
		CompanyName: "Mercado Livre",
		Address:     "Melicidade",
		Telephone:   "98787687",
		LocalityId:  1,
	}

	slList := []domain.Seller{
		{
			ID:          1,
			Cid:         22,
			CompanyName: "Mercado Fresco",
			Address:     "Rua Meli",
			Telephone:   "34235432",
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  1,
		},
	}

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(slList, nil)
	apiMock.EXPECT().Update(context.TODO(), 1, 20, "Mercado Livre", "Melicidade", "98787687", 1).Return(sl, nil)

	result, err := service.Update(context.TODO(), 1, 20, "Mercado Livre", "Melicidade", "98787687", 1)
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
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  1,
		},
	}

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(slList, nil)
	apiMock.EXPECT().Update(context.TODO(), 10, 20, "Mercado Livre", "Melicidade", "98787687", 1).Return(sl, errors.New("seller 10 not found"))

	result, err := service.Update(context.TODO(), 10, 20, "Mercado Livre", "Melicidade", "98787687", 1)
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
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         23,
			CompanyName: "Mercado Pago",
			Address:     "Rua Parque",
			Telephone:   "12349870",
			LocalityId:  1,
		},
	}

	apiMock, _, service := callMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(slList, nil)

	result, err := service.Update(context.TODO(), 10, 22, "Mercado Livre", "Melicidade", "98787687", 1)
	assert.NotNil(t, err)
	assert.Equal(t, result, sl)
}

func TestService_Delete_Ok(t *testing.T) {
	apiMock, _, service := callMock(t)

	apiMock.EXPECT().Delete(context.TODO(), gomock.Eq(id)).Return(nil)

	err := service.Delete(context.TODO(), id)
	assert.Nil(t, err)
}

func TestService_Delete_NOk(t *testing.T) {
	apiMock, _, service := callMock(t)

	apiMock.EXPECT().Delete(context.TODO(), gomock.Eq(id)).Return(errors.New("id is not valid"))

	err := service.Delete(context.TODO(), id)
	assert.NotNil(t, err)
}
