package service

import (
	"context"
	"errors"

	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain/mock"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func callBuyersMock(t *testing.T) (*mock_domain.MockRepository, domain.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_domain.NewMockRepository(ctrl)
	service := NewService(apiMock)
	return apiMock, service
}

//CREATE create_ok Se contiver os campos necessários, será criado
func TestService_Create_Ok(t *testing.T) {
	buyList := []domain.Buyer{
		{
			Id: 1,
			CardNumberId: "2",
			FirstName: "Fernando",
			LastName: "Souza",
		},
		{
			Id: 2,
			CardNumberId: "3",
			FirstName: "Marcela",
			LastName: "Vieira",
		},
	}

	buy := domain.Buyer{
		Id:           3,
		CardNumberId: "5",
		FirstName:    "Douglas",
		LastName:     "Mendes",
	}

	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().GetAll(context.TODO()).Return(buyList, nil)
	apiMock.EXPECT().Create(
		context.TODO(),
		"5",
		"Douglas",
		"Mendes",
	).Return(&buy, nil)

	result, err := service.Create(context.TODO(), "5", "Douglas", "Mendes")
	assert.Equal(t, result, &buy)
	assert.Nil(t, err)

}

//CREATE create_conflict Se o card_number_id já existir, ele não pode ser criado
func TestService_Create_Nok(t *testing.T) {
	buyList := []domain.Buyer{
		{
			Id: 1,
			CardNumberId: "2",
			FirstName: "Fernando",
			LastName: "Souza",
		},
		{
			Id: 2,
			CardNumberId: "3",
			FirstName: "Marcela",
			LastName: "Vieira",
		},
	}

	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().GetAll(context.TODO()).Return(buyList, nil)
	apiMock.EXPECT().Create(context.TODO(), "3", "Douglas", "Mendes").Return(&domain.Buyer{}, errors.New("this card number id already exists"))

	_, err := service.Create(context.TODO(), "3", "Douglas", "Mendes")
	assert.Equal(t, assert.NotNil(t, err), true)
	assert.EqualError(t, err, "this card number id already exists")
}

//READ find_all Se a lista tiver "n" elementos, retornará uma quantidade do total de elementos
func TestService_GetAll(t *testing.T) {
	buyList := []domain.Buyer{
		{
			Id: 1,
			CardNumberId: "2",
			FirstName: "Fernando",
			LastName: "Souza",
		},
		{
			Id: 2,
			CardNumberId: "3",
			FirstName: "Marcela",
			LastName: "Vieira",
		},
	}

	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().GetAll(context.TODO()).Return(buyList, nil)

	result, err := service.GetAll(context.TODO())
	assert.Equal(t, len(result), len(buyList))
	assert.Nil(t, err)
}

func TestService_GetAll_NOk(t *testing.T) {
	bList := make([]domain.Buyer, 0)

	apiMock, service := callBuyersMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(bList, errors.New("erro"))

	_, err := service.GetAll(context.TODO())
	assert.NotNil(t, err)
}

//READ find_by_id_non_existent Se o elemento procurado por id não existir, retorna null
func TestService_GetById_Nok(t *testing.T) {
	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().GetById(context.TODO(), gomock.Eq(1)).Return(&domain.Buyer{}, errors.New("Buyer 1 not found"))

	_, err := service.GetById(context.TODO(), 1)
	assert.NotNil(t, err)
}

//READ find_by_id_existent Se o elemento procurado por id existir.
func TestService_GetById_ok(t *testing.T) {
	buy := domain.Buyer{
		Id:           1,
		CardNumberId: "5",
		FirstName:    "Douglas",
		LastName:     "Mendes",
	}

	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().GetById(context.TODO(), gomock.Eq(1)).Return(&buy, nil)

	result, err := service.GetById(context.TODO(), 1)
	assert.Equal(t, result.Id, 1)
	assert.Nil(t, err)
}

func TestService_GetOrdersByBuyers_Ok(t *testing.T) {

	by := []domain.OrdersByBuyers{
		{
			Id: 1,
			CardNumberId: "44dm",
			FirstName: "Will",
			LastName: "Spencer",
			PurchaseOrdersCount: 8,
		},
	}

	apiMock, service := callBuyersMock(t)

	apiMock.EXPECT().GetOrdersByBuyers(context.TODO(), 1).Return(by, nil)

	result, err := service.GetOrdersByBuyers(context.TODO(), 1)
	assert.Equal(t, 8, result[len(result)-1].PurchaseOrdersCount)
	assert.Nil(t, err)
}

func TestService_GetOrdersByBuyers_NOk(t *testing.T) {

	apiMock, service := callBuyersMock(t)

	apiMock.EXPECT().GetOrdersByBuyers(context.TODO(), 1).Return([]domain.OrdersByBuyers{}, errors.New("seller not found"))

	_, err := service.GetOrdersByBuyers(context.TODO(), 1)
	assert.NotNil(t, err)
}

//DELETE - delete_non_existent - Quando o funcionário não existir, será retornado null.
func TestService_Delete_Ok(t *testing.T) {
	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().Delete(context.TODO(), 1).Return(nil)
	err := service.Delete(context.TODO(), 1)
	assert.Nil(t, err)
}

func TestService_Delete_Nok(t *testing.T) {
	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().Delete(context.TODO(), 1).Return(errors.New("buyer 1 not found"))
	err := service.Delete(context.TODO(), 1)
	assert.NotNil(t, err)
}

func TestService_Update_Ok(t *testing.T) {
	buyList := []domain.Buyer{
		{
			Id: 1,
			CardNumberId: "2",
			FirstName: "Fernando",
			LastName: "Souza",
		},
		{
			Id: 2,
			CardNumberId: "3",
			FirstName: "Marcela",
			LastName: "Vieira",
		},
	}

	buy := domain.Buyer{
		Id:           3,
		CardNumberId: "5",
		FirstName:    "Douglas",
		LastName:     "Mendes",
	}
	apiMock, service := callBuyersMock(t)

	apiMock.EXPECT().GetAll(context.TODO()).Return(buyList, nil)
	apiMock.EXPECT().Update(context.TODO(), 3, "5", "Douglas", "Mendes").Return(&buy, nil)

	result, err := service.Update(context.TODO(), 3, "5", "Douglas", "Mendes")
	assert.Nil(t, err)
	assert.Equal(t, result, &buy)
}

func TestService_Update_Nok(t *testing.T) {
	buyList := []domain.Buyer{
		{
			Id:           1,
			CardNumberId: "2",
			FirstName:    "Fernando",
			LastName:     "Souza",
		},
		{
			Id:           2,
			CardNumberId: "3",
			FirstName:    "Marcela",
			LastName:     "Vieira",
		},
	}

	apiMock, service := callBuyersMock(t)
	apiMock.EXPECT().GetAll(context.TODO()).Return(buyList, nil)

	_, err := service.Update(context.TODO(), 1, "3", "Joao", "Zinho")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "this Buyer already exists")

}
