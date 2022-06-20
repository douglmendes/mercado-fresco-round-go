package test

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees"
	mock_employees "github.com/douglmendes/mercado-fresco-round-go/internal/employees/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func callMock(t *testing.T) (*mock_employees.MockRepository, employees.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMock := mock_employees.NewMockRepository(ctrl)
	service := employees.NewService(apiMock)
	return apiMock, service
}

//CREATE create_ok Se contiver os campos necessários, será criado
func TestService_Create_Ok(t *testing.T) {
	buyList := []buyers.Buyer{
		{
			2,
			"2",
			"Fernando",
			"Souza",
		},
		{
			3,
			"3",
			"Marcela",
			"Vieira",
		},
	}

	buy := buyers.Buyer{
		Id:           2,
		CardNumberId: "2",
		FirstName:    "Fernando",
		LastName:     "Souza",
	}
	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(buyList, nil)
	apiMock.EXPECT().Create("2", "Fernando", "Souza").Return(buy, nil)
	//service
	result, err := service.Create("5050", "Renata", "Leal", 3)
	assert.Equal(t, result, buy)
	assert.Nil(t, err)

}
