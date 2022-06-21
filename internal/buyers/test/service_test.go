package test

import (
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	mock_buyers "github.com/douglmendes/mercado-fresco-round-go/internal/buyers/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func callMock(t *testing.T) (*mock_buyers.MockRepository, buyers.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_buyers.NewMockRepository(ctrl)
	service := buyers.NewService(apiMock)
	return apiMock, service
}

//CREATE create_ok Se contiver os campos necessários, será criado
func TestService_Create_Ok(t *testing.T) {
	buyList := []buyers.Buyer{
		{
			1,
			"2",
			"Fernando",
			"Souza",
		},
		{
			2,
			"3",
			"Marcela",
			"Vieira",
		},
	}

	buy := buyers.Buyer{
		Id:           3,
		CardNumberId: "5",
		FirstName:    "Douglas",
		LastName:     "Mendes",
	}
	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(buyList, nil)
	apiMock.EXPECT().Create(3, "5", "Douglas", "Mendes").Return(buy, nil)
	//service
	result, err := service.Create("5", "Douglas", "Mendes")
	assert.Equal(t, result, buy)
	assert.Nil(t, err)

}

//CREATE create_conflict Se o card_number_id já existir, ele não pode ser criado
func TestService_Create_Nok(t *testing.T) {
	buy := buyers.Buyer{}
	buyList := []buyers.Buyer{
		{
			1,
			"2",
			"Fernando",
			"Souza",
		},
		{
			2,
			"3",
			"Marcela",
			"Vieira",
		},
	}

	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(buyList, nil)
	apiMock.EXPECT().Create(3, "3", "Douglas", "Mendes").Return(buyers.Buyer{}, errors.New("this Buyer already exists"))
	//service
	result, err := service.Create("3", "Douglas", "Mendes")
	assert.Equal(t, result, buy)
	assert.NotNil(t, err)

}

//READ find_all Se a lista tiver "n" elementos, retornará uma quantidade do total de elementos
func TestService_GetAll(t *testing.T) {
	buyList := []buyers.Buyer{
		{
			1,
			"2",
			"Fernando",
			"Souza",
		},
		{
			2,
			"3",
			"Marcela",
			"Vieira",
		},
	}

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetAll().Return(buyList, nil)

	result, err := service.GetAll()
	assert.Equal(t, len(result), len(buyList))
	assert.Nil(t, err)
}

//READ find_by_id_non_existent Se o elemento procurado por id não existir, retorna null
func TestService_GetById_Nok(t *testing.T) {
	apiMock, service := callMock(t)
	apiMock.EXPECT().GetById(gomock.Eq(1)).Return(buyers.Buyer{}, errors.New("Buyer 1 not found"))

	_, err := service.GetById(1)
	assert.NotNil(t, err)
}

//READ find_by_id_existent Se o elemento procurado por id existir.
func TestService_GetById_ok(t *testing.T) {
	buy := buyers.Buyer{
		Id:           1,
		CardNumberId: "5",
		FirstName:    "Douglas",
		LastName:     "Mendes",
	}

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetById(gomock.Eq(1)).Return(buy, nil)

	result, err := service.GetById(1)
	assert.Equal(t, result.Id, 1)
	assert.Nil(t, err)
}

//DELETE - delete_non_existent - Quando o funcionário não existir, será retornado null.
func TestService_Delete_Ok(t *testing.T) {
	apiMock, service := callMock(t)
	apiMock.EXPECT().Delete(1).Return(nil)
	err := service.Delete(1)
	assert.Nil(t, err)
}

func TestService_Delete_Nok(t *testing.T) {
	apiMock, service := callMock(t)
	apiMock.EXPECT().Delete(1).Return(errors.New("buyer 1 not found"))
	err := service.Delete(1)
	assert.NotNil(t, err)
}

func TestService_Update_Ok(t *testing.T) {
	buyList := []buyers.Buyer{
		{
			1,
			"2",
			"Fernando",
			"Souza",
		},
		{
			2,
			"3",
			"Marcela",
			"Vieira",
		},
	}

	buy := buyers.Buyer{
		Id:           3,
		CardNumberId: "5",
		FirstName:    "Douglas",
		LastName:     "Mendes",
	}
	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(buyList, nil)
	apiMock.EXPECT().Update(3, "5", "Douglas", "Mendes").Return(buy, nil)

	result, err := service.Update(3, "5", "Douglas", "Mendes")
	assert.Nil(t, err)
	assert.Equal(t, result, buy)
}

func TestService_Update_Nok(t *testing.T) {
	buy := buyers.Buyer{}
	buyList := []buyers.Buyer{
		{
			1,
			"2",
			"Fernando",
			"Souza",
		},
		{
			2,
			"3",
			"Marcela",
			"Vieira",
		},
	}
	apiMock, service := callMock(t)

	apiMock.EXPECT().GetAll().Return(buyList, nil)
	apiMock.EXPECT().Update(3, "50", "Joao", "Zinho").Return(buy, errors.New("this Buyer already exists"))

	result, err := service.Update(3, "50", "Joao", "Zinho")
	assert.NotNil(t, err)
	assert.Equal(t, result, buy)

}