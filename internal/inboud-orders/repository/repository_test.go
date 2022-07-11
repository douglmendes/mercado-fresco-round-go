package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_GetAll_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

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
	rows := sqlmock.NewRows([]string{
		"id", "order_date", "order_number", "employee_id", "product_batch_id", "warehouse_id",
	}).AddRow(
		ioList[0].Id,
		ioList[0].OrderDate,
		ioList[0].OrderNumber,
		ioList[0].EmployeeId,
		ioList[0].ProductBatchId,
		ioList[0].WarehouseId,
	).AddRow(
		ioList[1].Id,
		ioList[1].OrderDate,
		ioList[1].OrderNumber,
		ioList[1].EmployeeId,
		ioList[1].ProductBatchId,
		ioList[1].WarehouseId,
	)
	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	ioRepo := NewRepository(db)
	result, err := ioRepo.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, ioList, result)

}
func TestRepository_GetAll_Nok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ioList := []domain.InboudOrder(nil)

	mock.ExpectQuery(queryGetAll).WillReturnError(errors.New("erro"))

	ioRepo := NewRepository(db)

	result, err := ioRepo.GetAll(context.TODO())
	assert.Error(t, err)
	assert.Equal(t, ioList, result)

}

func TestRepository_Create_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	io := domain.InboudOrder{
		1,
		"1900-01-01",
		"order#1",
		2,
		1,
		1,
	}

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
		io.OrderDate,
		io.OrderNumber,
		io.EmployeeId,
		io.ProductBatchId,
		io.WarehouseId,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	ioRepo := NewRepository(db)

	result, err := ioRepo.Create(context.TODO(), "1900-01-01", "order#1", 2, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, result.Id, 1)
}

func TestRepository_Create_Nok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).
		WithArgs(0, 0, 0, 0, 0).WillReturnResult(sqlmock.NewResult(1, 1))

	ioRepo := NewRepository(db)

	_, err = ioRepo.Create(context.TODO(), "1900-01-01", "order#1", 2, 1, 1)

	assert.Error(t, err)
}

func TestRepository_GetByEmployee_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ioReport := []domain.EmployeeInboudOrder{
		{
			Id:               1,
			CardNumberId:     "5555",
			FirstName:        "Douglas",
			LastName:         "Mendes",
			WarehouseId:      3,
			InboudOrderCount: 5,
		},
	}
	ioList := []domain.EmployeeInboudOrder{
		{
			Id:               1,
			CardNumberId:     "5555",
			FirstName:        "Douglas",
			LastName:         "Mendes",
			WarehouseId:      3,
			InboudOrderCount: 5,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inboud_order_count",
	}).AddRow(
		ioList[0].Id,
		ioList[0].CardNumberId,
		ioList[0].FirstName,
		ioList[0].LastName,
		ioList[0].WarehouseId,
		ioList[0].InboudOrderCount,
	)
	mock.ExpectQuery(regexp.QuoteMeta(queryGetByEmplyee)).WillReturnRows(rows)

	empIo := NewRepository(db)

	result, err := empIo.GetByEmployee(context.TODO(), 1)

	assert.NoError(t, err)
	assert.Equal(t, ioReport[0], result[0])

}

func TestRepository_GetByEmployee_Nok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetByEmplyee).WillReturnError(errors.New("error"))

	ioRepo := NewRepository(db)
	var i []domain.EmployeeInboudOrder
	result, err := ioRepo.GetByEmployee(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, i, result)
}

func TestRepository_GetByEmployee_equal_zero(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ioReport := []domain.EmployeeInboudOrder{
		{
			Id:               1,
			CardNumberId:     "5555",
			FirstName:        "Douglas",
			LastName:         "Mendes",
			WarehouseId:      3,
			InboudOrderCount: 5,
		},
	}
	ioList := []domain.EmployeeInboudOrder{
		{
			Id:               1,
			CardNumberId:     "5555",
			FirstName:        "Douglas",
			LastName:         "Mendes",
			WarehouseId:      3,
			InboudOrderCount: 5,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inboud_order_count",
	}).AddRow(
		ioList[0].Id,
		ioList[0].CardNumberId,
		ioList[0].FirstName,
		ioList[0].LastName,
		ioList[0].WarehouseId,
		ioList[0].InboudOrderCount,
	)
	mock.ExpectQuery(regexp.QuoteMeta(queryGetByEmplyee)).WillReturnRows(rows)

	empIo := NewRepository(db)

	result, err := empIo.GetByEmployee(context.TODO(), 0)

	assert.NoError(t, err)
	assert.Equal(t, ioReport[0], result[0])

}
