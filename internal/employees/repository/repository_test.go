package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_GetAll_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	empList := []domain.Employee{
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
	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id",
	}).AddRow(
		empList[0].Id,
		empList[0].CardNumberId,
		empList[0].FirstName,
		empList[0].LastName,
		empList[0].WarehouseId,
	).AddRow(
		empList[1].Id,
		empList[1].CardNumberId,
		empList[1].FirstName,
		empList[1].LastName,
		empList[0].WarehouseId,
	)
	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	empRepo := NewRepository(db)

	result, err := empRepo.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, len(result), 2)
}

func TestRepository_GetAll_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	empList := []domain.Employee(nil)

	mock.ExpectQuery(queryGetAll).WillReturnError(errors.New("erro"))

	empRepo := NewRepository(db)

	result, err := empRepo.GetAll(context.TODO())
	assert.Error(t, err)
	assert.Equal(t, empList, result)

}

func TestRepositoru_GetById_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	emp := domain.Employee{
		44,
		"3030",
		"Douglas",
		"Mendes",
		3,
	}

	row := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id",
	}).AddRow(
		emp.Id,
		emp.CardNumberId,
		emp.FirstName,
		emp.LastName,
		emp.WarehouseId,
	)
	mock.ExpectQuery(queryGetById).WillReturnRows(row)

	empRepo := NewRepository(db)

	result, err := empRepo.GetById(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(44), result.Id)
}

func TestRepository_GetById_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(errors.New("error"))

	empRepo := NewRepository(db)
	var e *domain.Employee
	result, err := empRepo.GetById(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, e, result)
}

func TestRepository_Create_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	emp := domain.Employee{
		1,
		"3030",
		"Douglas",
		"Mendes",
		3,
	}

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
		emp.CardNumberId,
		emp.FirstName,
		emp.LastName,
		emp.WarehouseId,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	empRepo := NewRepository(db)

	result, err := empRepo.Create(context.TODO(), "3030", "Douglas", "Mendes", 3)
	assert.NoError(t, err)
	assert.Equal(t, result.Id, int64(1))

}

func TestRepository_Create_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).
		WithArgs(0, 0, 0, 0).WillReturnResult(sqlmock.NewResult(1, 1))

	empRepo := NewRepository(db)

	_, err = empRepo.Create(context.TODO(), "3030", "Douglas", "Mendes", 3)

	assert.Error(t, err)

}

func TestRepository_Update_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	emp := domain.Employee{
		1,
		"3030",
		"Douglas",
		"Mendes",
		3,
	}

	empUpdate := domain.Employee{
		1,
		"3030",
		"Douglas",
		"Leonardo",
		3,
	}

	row := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id",
	}).AddRow(
		emp.Id,
		emp.CardNumberId,
		emp.FirstName,
		emp.LastName,
		emp.WarehouseId,
	)
	mock.ExpectQuery(queryGetById).WillReturnRows(row)

	mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).WithArgs(
		empUpdate.Id,
		empUpdate.CardNumberId,
		empUpdate.FirstName,
		empUpdate.LastName,
		empUpdate.WarehouseId,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	empRepo := NewRepository(db)

	result, err := empRepo.Update(context.TODO(), 1, "3030", "Douglas", "Leonardo", 3)
	assert.NoError(t, err)
	assert.Equal(t, &empUpdate, result)
}
func TestRepository_Update_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryUpdate).WillReturnError(errors.New("employee not found"))

	empRepo := NewRepository(db)

	_, err = empRepo.Update(context.TODO(), 1, "3030", "Douglas", "Leonardo", 3)
	assert.Error(t, err)
}

func TestRepository_Delete_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryDelete)).WithArgs(
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	empRepo := NewRepository(db)

	err = empRepo.Delete(context.TODO(), 1)
	assert.NoError(t, err)
}

func TestRepositoru_Delete_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryDelete).WillReturnError(errors.New("error"))

	empRepo := NewRepository(db)

	err = empRepo.Delete(context.TODO(), 1)
	assert.Error(t, err)
}
