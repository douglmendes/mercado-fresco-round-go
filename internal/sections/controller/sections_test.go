package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
	mock_sections "github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain/mock"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	pathSections   = "/api/v1/sections/"
	pathIdSections = "/api/v1/sections/:id"
	idSections     = "1"
)

func mockSections(t *testing.T) (*mock_sections.MockService, *SectionsController, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock_sections.NewMockService(ctrl)
	handler := NewSectionsController(service)
	api := gin.New()
	return service, handler, api
}

func TestSections_Create_OK(t *testing.T) {
	service, handler, api := mockSections(t)
	api.POST(pathSections, handler.Create)

	newSection := domain.Section{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 12,
		MinimumTemperature: 14,
		CurrentCapacity:    25,
		MinimumCapacity:    5,
		MaximumCapacity:    50,
		WarehouseId:        3,
		ProductTypeId:      5,
	}

	service.EXPECT().Create(3, 12, 14, 25, 5, 50, 3, 5).Return(&newSection, nil)

	payload := `{
		"section_number": 3,
		"current_temperature": 12,
		"minimum_temperature": 14,
		"current_capacity": 25,
		"minimum_capacity": 5,
		"maximum_capacity": 50,
		"warehouse_id": 3,
		"product_type_id": 5
	}`

	req := httptest.NewRequest(http.MethodPost, pathSections, strings.NewReader(payload))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	jsonSection, err := json.Marshal(response.NewResponse(newSection))
	assert.Nil(t, err)
	assert.Equal(t, jsonSection, resp.Body.Bytes())
}

func TestSections_Create_Fail(t *testing.T) {
	_, handler, api := mockSections(t)
	api.POST(pathSections, handler.Create)

	payload := `{
		"section_number": 3,
		"current_temperature": 12,
		"minimum_temperature": 14,
		"warehouse_id": 3,
		"product_type_id": 5
	}`

	req := httptest.NewRequest(http.MethodPost, pathSections, strings.NewReader(payload))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
}

func TestSections_Create_Conflict(t *testing.T) {
	service, handler, api := mockSections(t)
	api.POST(pathSections, handler.Create)

	expectedError := domain.ErrorConflict{SectionNumber: 3}

	service.EXPECT().Create(3, 12, 14, 25, 5, 50, 3, 5).Return(nil, &expectedError)

	payload := `{
		"section_number": 3,
		"current_temperature": 12,
		"minimum_temperature": 14,
		"current_capacity": 25,
		"minimum_capacity": 5,
		"maximum_capacity": 50,
		"warehouse_id": 3,
		"product_type_id": 5
	}`

	req := httptest.NewRequest(http.MethodPost, pathSections, strings.NewReader(payload))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusConflict, resp.Code)
}

func TestSections_Find_All(t *testing.T) {
	service, handler, api := mockSections(t)

	db := []domain.Section{
		{
			Id:                 1,
			SectionNumber:      3,
			CurrentTemperature: 12,
			MinimumTemperature: 14,
			CurrentCapacity:    25,
			MinimumCapacity:    5,
			MaximumCapacity:    50,
			WarehouseId:        3,
			ProductTypeId:      5,
		},
		{
			Id:                 2,
			SectionNumber:      4,
			CurrentTemperature: 13,
			MinimumTemperature: 15,
			CurrentCapacity:    26,
			MinimumCapacity:    6,
			MaximumCapacity:    51,
			WarehouseId:        3,
			ProductTypeId:      5,
		},
	}

	api.GET(pathSections, handler.GetAll)
	service.EXPECT().GetAll().Return(db, nil)

	req := httptest.NewRequest(http.MethodGet, pathSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	expecBody := struct{ Data []domain.Section }{}
	err := json.Unmarshal(resp.Body.Bytes(), &expecBody)
	assert.Nil(t, err)

	assert.Equal(t, db, expecBody.Data)
}

func TestSections_Find_All_Error(t *testing.T) {
	service, handler, api := mockSections(t)
	api.GET(pathSections, handler.GetAll)

	service.EXPECT().GetAll().Return([]domain.Section{}, errors.New("internal server error"))

	req := httptest.NewRequest(http.MethodGet, pathSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestSections_Find_All_Empty(t *testing.T) {
	service, handler, api := mockSections(t)
	api.GET(pathSections, handler.GetAll)

	service.EXPECT().GetAll().Return([]domain.Section{}, nil)

	req := httptest.NewRequest(http.MethodGet, pathSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestSections_Find_By_Id_Non_Existent(t *testing.T) {
	service, handler, api := mockSections(t)
	api.GET(pathIdSections, handler.GetById)

	service.EXPECT().GetById(1).Return(nil, &domain.ErrorNotFound{Id: 1})

	req := httptest.NewRequest(http.MethodGet, pathSections+idSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSections_Find_By_Id_Existent(t *testing.T) {
	service, handler, api := mockSections(t)
	api.GET(pathIdSections, handler.GetById)

	db := domain.Section{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 12,
		MinimumTemperature: 14,
		CurrentCapacity:    25,
		MinimumCapacity:    5,
		MaximumCapacity:    50,
		WarehouseId:        3,
		ProductTypeId:      5,
	}

	service.EXPECT().GetById(1).Return(&db, nil)

	req := httptest.NewRequest(http.MethodGet, pathSections+idSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	expectBody := struct{ Data domain.Section }{}
	err := json.Unmarshal(resp.Body.Bytes(), &expectBody)
	assert.Nil(t, err)

	assert.Equal(t, db, expectBody.Data)
}

func TestSections_Find_By_Id_Bad_Request(t *testing.T) {
	_, handler, api := mockSections(t)
	api.GET(pathIdSections, handler.GetById)

	//service.EXPECT().GetById(1).Return(nil, &sections.ErrorNotFound{Id: 1})

	req := httptest.NewRequest(http.MethodGet, pathSections+"a", nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestSections_Update_OK(t *testing.T) {
	service, handler, api := mockSections(t)
	api.PATCH(pathIdSections, handler.Update)

	db := domain.Section{
		Id:                 1,
		SectionNumber:      3,
		CurrentTemperature: 15,
		MinimumTemperature: 14,
		CurrentCapacity:    25,
		MinimumCapacity:    15,
		MaximumCapacity:    50,
		WarehouseId:        3,
		ProductTypeId:      5,
	}

	service.EXPECT().Update(1, map[string]int{"current_temperature": 15, "minimum_capacity": 15}).Return(&db, nil)

	payload := `{
		"current_temperature": 15,
		"minimum_capacity": 15
	}`

	req := httptest.NewRequest(http.MethodPatch, pathSections+idSections, strings.NewReader(payload))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	expecBody := struct{ Data domain.Section }{}
	err := json.Unmarshal(resp.Body.Bytes(), &expecBody)
	assert.Nil(t, err)

	assert.Equal(t, db, expecBody.Data)
}

func TestSections_Update_Non_Existent(t *testing.T) {
	service, handler, api := mockSections(t)
	api.PATCH(pathIdSections, handler.Update)

	service.EXPECT().Update(1, map[string]int{"current_temperature": 15, "minimum_capacity": 15}).Return(nil, &domain.ErrorNotFound{Id: 1})

	payload := `{
		"current_temperature": 15,
		"minimum_capacity": 15
	}`

	req := httptest.NewRequest(http.MethodPatch, pathSections+idSections, strings.NewReader(payload))
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSections_Delete_Non_Existent(t *testing.T) {
	service, handler, api := mockSections(t)
	api.DELETE(pathIdSections, handler.Delete)

	service.EXPECT().Delete(1).Return(&domain.ErrorNotFound{Id: 1})

	req := httptest.NewRequest(http.MethodDelete, pathSections+idSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestSections_Delete_OK(t *testing.T) {
	service, handler, api := mockSections(t)
	api.DELETE(pathIdSections, handler.Delete)

	service.EXPECT().Delete(1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, pathSections+idSections, nil)
	resp := httptest.NewRecorder()
	api.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}
