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

//CREATE create_ok Se contiver os campos necessários, será criado
func TestService_Create_Ok(t *testing.T) {
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
		Id:           3,
		CardNumberId: "5050",
		FirstName:    "Renata",
		LastName:     "Leal",
		WarehouseId:  3,
	}
	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(empList, nil)
	apiMock.EXPECT().Create(3, "5050", "Renata", "Leal", 3).Return(emp, nil)
	//service
	result, err := service.Create("5050", "Renata", "Leal", 3)
	assert.Equal(t, result, emp)
	assert.Nil(t, err)

}

//CREATE create_conflict Se o card_number_id já existir, ele não pode ser criado
func TestService_Create_Nok(t *testing.T) {
	emp := employees.Employee{}
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

	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().LastID().Return(2, nil)
	apiMock.EXPECT().GetAll().Return(empList, nil)
	apiMock.EXPECT().Create(3, "3030", "Renata", "Leal", 3).Return(employees.Employee{}, errors.New("this card number id already exists"))
	//service
	result, err := service.Create("3030", "Renata", "Leal", 3)
	assert.NotNil(t, err)
	assert.Equal(t, result, emp)

}

//READ find_all Se a lista tiver "n" elementos, retornará uma quantidade do total de elementos
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
	//repository
	apiMock, service := callMock(t)
	apiMock.EXPECT().GetAll().Return(emp, nil)
	//service
	result, err := service.GetAll()
	assert.Equal(t, len(result), len(emp))
	assert.Nil(t, err)
}

//READ find_by_id_non_existent Se o elemento procurado por id não existir, retorna null
func TestService_GetById_Nok(t *testing.T) {

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetById(gomock.Eq(1)).Return(employees.Employee{}, errors.New("Employee 1 not found id"))

	_, err := service.GetById(1)
	assert.NotNil(t, err)
}

//READ find_by_id_existent Se o elemento procurado por id existir, ele
func TestService_GetById_Ok(t *testing.T) {
	emp := employees.Employee{
		Id:           1,
		CardNumberId: "3030",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}

	apiMock, service := callMock(t)
	apiMock.EXPECT().GetById(gomock.Eq(1)).Return(emp, nil)

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

//DELETE delete_ok Se a exclusão for bem-sucedida, o item não aparecerá na lista.
func TestService_Delete_Nok(t *testing.T) {
	apiMock, service := callMock(t)
	apiMock.EXPECT().Delete(1).Return(errors.New("employee 1 not found"))
	err := service.Delete(1)
	assert.NotNil(t, err)
}

//UPDATE update_existent Quando a atualização dos dados for bem-sucedida, o
//funcionário será devolvido com as informações atualizadas
func TestService_Update_Ok(t *testing.T) {

	emp := employees.Employee{
		Id:           1,
		CardNumberId: "5050",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}
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
	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().GetAll().Return(empList, nil)
	apiMock.EXPECT().Update(1, "5050", "Douglas", "Mendes", 3).Return(emp, nil)
	//service
	result, err := service.Update(1, "5050", "Douglas", "Mendes", 3)
	assert.Nil(t, err)
	assert.Equal(t, result, emp)

}

//UPDATE update_non_existent Se o funcionário a ser atualizado não existir, será retornado null.
func TestService_Update_Nok(t *testing.T) {

	emp := employees.Employee{}
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
	apiMock, service := callMock(t)
	//repository
	apiMock.EXPECT().GetAll().Return(empList, nil)
	apiMock.EXPECT().Update(50, "5050", "Douglas", "Mendes", 3).Return(emp, errors.New("employee 60 not found"))
	//service
	result, err := service.Update(50, "5050", "Douglas", "Mendes", 3)
	assert.NotNil(t, err)
	assert.Equal(t, result, emp)

}
