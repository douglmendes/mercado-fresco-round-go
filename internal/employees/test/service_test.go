package test

import (
	"errors"
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

//CREATE
func TestService_Create(t *testing.T) {
	empList := []employees.Employee{
		{
			1,
			"3030",
			"Douglas",
			"Mendes",
			3,
		},
		{
			2,
			"40",
			"Gustavo",
			"Naganuma",
			33,
		},
	}

	emp := employees.Employee{
		3,
		"5050",
		"Renata",
		"Leal",
		3,
	}
	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(empList, nil)
	apiMock.EXPECT().Create(3, "5050", "Renata", "Leal", 3).Return(emp, nil)
	//service
	result, err := service.Create("5050", "Renata", "Leal", 3)
	assert.Equal(t, result, emp)
	assert.NotNil(t, err)

}

//READ - GETALL
func TestService_GetAll(t *testing.T) {
	emp := []employees.Employee{
		{
			1,
			"3030",
			"Douglas",
			"Mendes",
			3,
		},
		{
			2,
			"40",
			"Gustavo",
			"Naganuma",
			33,
		},
	}

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetAll().Return(emp, nil)

	result, err := service.GetAll()
	assert.Equal(t, len(result), len(emp))
	assert.Nil(t, err)
}

//READ - GETBYID - OK
func TestService_GetById_Nok(t *testing.T) {

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetById(gomock.Eq(1)).Return(employees.Employee{}, errors.New("Employee 1 not found id"))

	_, err := service.GetById(1)
	assert.NotNil(t, err)
}

//READ - GETBYID - NÃO OK
func TestService_GetById_Ok(t *testing.T) {
	emp := employees.Employee{
		1,
		"3030",
		"Douglas",
		"Mendes",
		3,
	}

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetById(gomock.Eq(1)).Return(emp, nil)

	result, err := service.GetById(1)
	assert.Equal(t, result.Id, 1)
	assert.Nil(t, err)
}
