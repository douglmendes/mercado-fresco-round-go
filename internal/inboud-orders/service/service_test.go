package service

import (
	"context"
	"errors"
	employeeDomain "github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	employeeMock "github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain/mock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func callMock(t *testing.T) (*mock_domain.MockRepository, *employeeMock.MockRepository, domain.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiMockIo := mock_domain.NewMockRepository(ctrl)
	apiMockEmp := employeeMock.NewMockRepository(ctrl)
	serviceIo := NewService(apiMockIo, apiMockEmp)
	return apiMockIo, apiMockEmp, serviceIo
}

func TestService_Create_Ok(t *testing.T) {
	ioList := []domain.InboudOrder{
		{
			1,
			"1900-01-01",
			"order#1",
			2,
			1,
			1,
		},
		{
			2,
			"1900-01-01",
			"order#2",
			3,
			2,
			2,
		},
	}
	io := &domain.InboudOrder{
		Id:             4,
		OrderDate:      "1900-01-01",
		OrderNumber:    "order#3",
		EmployeeId:     4,
		ProductBatchId: 2,
		WarehouseId:    2,
	}
	emp := &employeeDomain.Employee{
		Id:           4,
		CardNumberId: "3030",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}
	apiMockIo, apiMockEmp, service := callMock(t)

	apiMockEmp.EXPECT().GetById(context.TODO(), int64(4)).Return(emp, nil)
	apiMockIo.EXPECT().GetAll(context.TODO()).Return(ioList, nil)
	apiMockIo.EXPECT().Create(context.TODO(), "1900-01-01", "order#3", 4, 2, 2).Return(io, nil)
	result, err := service.Create(context.TODO(), "1900-01-01", "order#3", 4, 2, 2)
	assert.Equal(t, io, result)
	assert.Nil(t, err)
}

func TestService_Create_Nok(t *testing.T) {
	ioList := []domain.InboudOrder{
		{
			1,
			"1900-01-01",
			"order#1",
			2,
			1,
			1,
		},
		{
			2,
			"1900-01-01",
			"order#2",
			3,
			2,
			2,
		},
	}
	emp := &employeeDomain.Employee{
		Id:           4,
		CardNumberId: "3030",
		FirstName:    "Douglas",
		LastName:     "Mendes",
		WarehouseId:  3,
	}
	apiMockIo, apiMockEmp, service := callMock(t)

	apiMockEmp.EXPECT().GetById(context.TODO(), int64(4)).Return(emp, nil)
	apiMockIo.EXPECT().GetAll(context.TODO()).Return(ioList, nil)
	apiMockIo.EXPECT().Create(context.TODO(), "1900-01-01", "order#3", 4, 2, 2).Return(nil, errors.New("employee number not found"))

	_, err := service.Create(context.TODO(), "1900-01-01", "order#3", 4, 2, 2)
	assert.NotNil(t, err)
}

func TestService_GetByEmployee(t *testing.T) {
	ioReport := []domain.EmployeeInboudOrder{{
		Id:               1,
		CardNumberId:     "5555",
		FirstName:        "Douglas",
		LastName:         "Mendes",
		WarehouseId:      3,
		InboudOrderCount: 5,
	},
	}

	apiMockIo, _, service := callMock(t)
	apiMockIo.EXPECT().GetByEmployee(context.TODO(), int64(1)).Return(ioReport, nil)

	result, err := service.GetByEmployee(context.TODO(), int64(1))
	assert.Equal(t, result, ioReport)
	assert.Nil(t, err)

}

func TestService_GetByEmployee_Nok(t *testing.T) {
	apiMock, _, service := callMock(t)
	apiMock.EXPECT().GetByEmployee(context.TODO(), int64(1)).Return(nil, errors.New("Employee 1 not found id"))

	_, err := service.GetByEmployee(context.TODO(), int64(1))
	assert.NotNil(t, err)

}
