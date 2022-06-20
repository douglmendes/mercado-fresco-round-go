package test

import (
	"bytes"
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/cmd/server/controllers"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees"
	mock_employees "github.com/douglmendes/mercado-fresco-round-go/internal/employees/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	relativePathEmployees = "/api/v1/employees/"
	target                = "/api/v1/employees/:id"
)

func callMockEmployees(t *testing.T) (*mock_employees.MockService, *controllers.EmployeesController) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_employees.NewMockService(ctrl)
	handler := controllers.NewEmployees(service)

	return service, handler
}

//READ find_all Se a lista tiver "n" elementos, retornará uma quantidade do total de elementos
func TestController_GetAll(t *testing.T) {
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

	service, handler := callMockEmployees(t)
	api := gin.New()
	api.GET(relativePathEmployees, handler.GetAll())

	service.EXPECT().GetAll().Return(empList, nil)
	req := httptest.NewRequest(http.MethodGet, relativePathEmployees, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestController_ById_BadRequest(t *testing.T) {

	_, handler := callMockEmployees(t)

	api := gin.New()
	api.GET(relativePathEmployees, handler.GetById())
	req := httptest.NewRequest(http.MethodGet, relativePathEmployees, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

}

//READ find_by_id_non_existent Se o elemento procurado por id não existir, retorna null
func TestController_ById_Nok(t *testing.T) {

	service, handler := callMockEmployees(t)
	api := gin.New()
	api.GET(target, handler.GetById())
	service.EXPECT().GetById(gomock.Eq(1)).Return(employees.Employee{}, errors.New("employee 1 not found id"))
	req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}

//READ find_by_id_existent Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado
func TestController_ById_Ok(t *testing.T) {

	emp := employees.Employee{
		Id:           1,
		CardNumberId: "3030",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}
	service, handler := callMockEmployees(t)

	//id := "1"
	api := gin.New()
	api.GET(target, handler.GetById())
	service.EXPECT().GetById(gomock.Eq(1)).Return(emp, nil)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	//respExpect := struct{ Data employees.Employee }{}
	//_ = json.Unmarshal(resp.Body.Bytes(), &respExpect.Data.Id)

	assert.Equal(t, http.StatusOK, resp.Code)
	//assert.Equal(t, emp.Id, respExpect.Data.Id)

}

//CREATE create_ok Se contiver os campos necessários, será criado
func TestController_Create_Ok(t *testing.T) {
	emp := employees.Employee{
		Id:           1,
		CardNumberId: "3030",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}
	service, handler := callMockEmployees(t)
	api := gin.New()
	api.POST(relativePathEmployees, handler.Create())
	service.EXPECT().Create("3030", "Douglas", "Mendes", 3).Return(emp, nil)
	body := `{"card_number_id": "3030","first_name": "Douglas","last_name": "Mendes","warehouse_id": 3}`
	req := httptest.NewRequest(http.MethodPost, relativePathEmployees, bytes.NewBuffer([]byte(body)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

}

//CREATE create_conflict Se o card_number_id já existir, ele não pode ser criado
func TestController_Create_Nok(t *testing.T) {

	service, handler := callMockEmployees(t)
	api := gin.New()
	api.POST(relativePathEmployees, handler.Create())

	service.EXPECT().Create(
		"3030",
		"Douglas",
		"Mendes",
		3,
	).Return(employees.Employee{}, errors.New("this card number id already exists"))

	body := `{"card_number_id": "3030","first_name": "Douglas","last_name": "Mendes","warehouse_id": 3}`
	req := httptest.NewRequest(http.MethodPost, relativePathEmployees, bytes.NewBuffer([]byte(body)))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)

}

//UPDATE update_ok Quando a atualização dos dados for bem
//sucedida, o funcionário será devolvido
//com as informações atualizadas
//juntamente com um código 200
func TestController_Update_Ok(t *testing.T) {
	emp := employees.Employee{
		Id:           1,
		CardNumberId: "3030",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}

	service, handler := callMockEmployees(t)
	api := gin.New()
	api.PATCH(target, handler.Update())
	service.EXPECT().Update(
		gomock.Eq(1),
		"3030",
		"Douglas",
		"Mendes",
		3,
	).Return(emp, nil)

	body := `{"card_number_id": "3030","first_name": "Douglas","last_name": "Mendes","warehouse_id": 3}`
	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/v1/employees/1",
		bytes.NewBuffer([]byte(body)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

//UPDATE update_non_existent Se o funcionário a ser atualizado não
//existir, um código 404 será retornado.
func TestController_Update_Nok(t *testing.T) {

	service, handler := callMockEmployees(t)
	api := gin.New()
	api.PATCH(target, handler.Update())
	service.EXPECT().Update(
		gomock.Eq(1),
		"3030",
		"Douglas",
		"Mendes",
		3,
	).Return(employees.Employee{}, errors.New("this employee already exists"))

	body := `{"card_number_id": "3030","first_name": "Douglas","last_name": "Mendes","warehouse_id": 3}`
	req := httptest.NewRequest(
		http.MethodPatch,
		"/api/v1/employees/1",
		bytes.NewBuffer([]byte(body)),
	)

	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}
func TestController_Delete_Ok(t *testing.T) {
	service, handler := callMockEmployees(t)
	api := gin.New()
	api.DELETE(target, handler.Delete())

	service.EXPECT().Delete(1).Return(nil)
	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/employees/1",
		nil,
	)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestController_Delete_Nok(t *testing.T) {
	service, handler := callMockEmployees(t)
	api := gin.New()
	api.DELETE(target, handler.Delete())

	service.EXPECT().Delete(1).Return(errors.New("this employee already exists"))
	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/employees/1",
		nil,
	)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
}
func TestController_Delete_BadRequest(t *testing.T) {
	_, handler := callMockEmployees(t)
	api := gin.New()
	api.DELETE(target, handler.Delete())

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/v1/employees/ixi",
		nil,
	)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
