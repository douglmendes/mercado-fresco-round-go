package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetAll_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	buyerMock := []domain.Buyer{
		{
			Id:           1,
			CardNumberId: "44dm",
			FirstName:    "Will",
			LastName:     "Spencer",
		},
		{
			Id:           2,
			CardNumberId: "234rt",
			FirstName:    "Letícia",
			LastName:     "Lumack",
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name",
	}).AddRow(
		buyerMock[0].Id,
		buyerMock[0].CardNumberId,
		buyerMock[0].FirstName,
		buyerMock[0].LastName,
	).AddRow(
		buyerMock[1].Id,
		buyerMock[1].CardNumberId,
		buyerMock[1].FirstName,
		buyerMock[1].LastName,
	)

	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	byRepo := NewRepository(db)

	result, err := byRepo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, "44dm", result[0].CardNumberId)
	assert.Equal(t, len(result), 2)
}

func TestRepository_GetAll_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetAll).WillReturnError(errors.New("erro"))

	byRepo := NewRepository(db)

	_, err = byRepo.GetAll()
	assert.Error(t, err)
}

func TestRepository_GetById_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	buyerMock := domain.Buyer{
		Id:           1,
		CardNumberId: "44dm",
		FirstName:    "Will",
		LastName:     "Spencer",
	}

	row := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name",
	}).AddRow(
		buyerMock.Id,
		buyerMock.CardNumberId,
		buyerMock.FirstName,
		buyerMock.LastName,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(row)

	byRepo := NewRepository(db)

	result, err := byRepo.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, "44dm", result.CardNumberId)
}

func TestRepository_GetById_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(errors.New("error"))

	byRepo := NewRepository(db)

	_, err = byRepo.GetById(1)
	assert.Error(t, err)
}

func TestRepository_GetById_NoId(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(sql.ErrNoRows)

	byRepo := NewRepository(db)

	_, err = byRepo.GetById(1)
	assert.Error(t, err)
}

func TestRepository_GetOrdersByBuyers_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderMock := domain.OrdersByBuyers{
		Id: 1,
		CardNumberId: "44dm",
		FirstName: "Will",
		LastName: "Spencer",
		PurchaseOrdersCount: 8,
	}

	row := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "purchase_orders_count",
	}).AddRow(
		orderMock.Id,
		orderMock.CardNumberId,
		orderMock.FirstName,
		orderMock.LastName,
		orderMock.PurchaseOrdersCount,
	)

	mock.ExpectQuery(regexp.QuoteMeta(queryGetOrdersByBuyer)).WillReturnRows(row)

	byRepo := NewRepository(db)

	result, err := byRepo.GetOrdersByBuyers(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 8, result[len(result)-1].PurchaseOrdersCount)	
}

func TestRepository_GetOrdersByBuyers_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(queryGetOrdersByBuyer)).WillReturnError(sql.ErrNoRows)

	byRepo := NewRepository(db)

	result, err := byRepo.GetOrdersByBuyers(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, []domain.OrdersByBuyers(nil), result)
}

func TestRepository_GetOrdersByBuyers_NoLocality(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(queryGetOrdersByBuyer)).WillReturnError(errors.New("error"))

	byRepo := NewRepository(db)

	_, err = byRepo.GetOrdersByBuyers(context.TODO(), 1)
	assert.Error(t, err)
}

func TestRepository_GetOrdersByBuyers_NoId_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderMock := domain.OrdersByBuyers{
		Id: 1,
		CardNumberId: "44dm",
		FirstName: "Will",
		LastName: "Spencer",
		PurchaseOrdersCount: 8,
	}

	row := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "purchase_orders_count",
	}).AddRow(
		orderMock.Id,
		orderMock.CardNumberId,
		orderMock.FirstName,
		orderMock.LastName,
		orderMock.PurchaseOrdersCount,
	)

	mock.ExpectQuery(regexp.QuoteMeta(queryGetOrdersByBuyers)).WillReturnRows(row)

	byRepo := NewRepository(db)

	result, err := byRepo.GetOrdersByBuyers(context.TODO(), 0)
	assert.NoError(t, err)
	assert.Equal(t, "Will", result[len(result)-1].FirstName)	
}

func TestRepository_GetOrdersByBuyers_NoId_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(queryGetOrdersByBuyers)).WillReturnError(errors.New("error"))

	byRepo := NewRepository(db)

	_, err = byRepo.GetOrdersByBuyers(context.TODO(), 0)
	assert.Error(t, err)
}

func TestRepository_Create_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	buyerMock := domain.Buyer{
		CardNumberId: "44dm",
		FirstName:    "Will",
		LastName:     "Spencer",
	}

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
		buyerMock.CardNumberId,
		buyerMock.FirstName,
		buyerMock.LastName,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	byRepo := NewRepository(db)

	result, err := byRepo.Create("44dm", "Will", "Spencer")
	assert.NoError(t, err)

	assert.Equal(t, "44dm", result.CardNumberId)
}

func TestRepository_Create_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).
		WithArgs(0, 0, 0, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	byRepo := NewRepository(db)

	_, err = byRepo.Create("44dm", "Will", "Spencer")

	assert.Error(t, err)
}

func TestRepository_Update_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	buyerMock := domain.Buyer{
		Id:           1,
		CardNumberId: "44dm",
		FirstName:    "Will",
		LastName:     "Spencer",
	}

	buyerMockUpdated := domain.Buyer{
		Id:           1,
		CardNumberId: "44dm",
		FirstName:    "Will",
		LastName:     "M. Spencer",
	}

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name",
	}).AddRow(
		buyerMock.Id,
		buyerMock.CardNumberId,
		buyerMock.FirstName,
		buyerMock.LastName,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).WithArgs(
		buyerMockUpdated.CardNumberId,
		buyerMockUpdated.FirstName,
		buyerMockUpdated.LastName,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	byRepo := NewRepository(db)

	result, err := byRepo.Update(1, "44dm", "Will", "M. Spencer")
	assert.NoError(t, err)
	assert.Equal(t, &buyerMockUpdated, result)
}

func TestRepository_Update_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryUpdate).WillReturnError(errors.New("buyer not found"))

	byRepo := NewRepository(db)

	_, err = byRepo.Update(1, "44dm", "Will", "M. Spencer")
	assert.Error(t, err)
}

func TestRepository_Delete_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryDelete)).WithArgs(
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	byRepo := NewRepository(db)

	err = byRepo.Delete(1)
	assert.NoError(t, err)
}

func TestRepository_Delete_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryDelete).WillReturnError(errors.New("error"))

	byRepo := NewRepository(db)

	err = byRepo.Delete(1)
	assert.Error(t, err)
}