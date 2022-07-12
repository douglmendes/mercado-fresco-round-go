package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productBatchMock := domain.ProductBatch{
		BatchNumber:        1,
		CurrentQuantity:    2,
		CurrentTemperature: 3,
		DueDate:            "2020-01-01",
		InitialQuantity:    4,
		ManufacturingDate:  "2020-01-01",
		ManufacturingHour:  5,
		MinimumTemperature: 6,
		ProductId:          7,
		SectionId:          8,
	}

	mock.ExpectExec(regexp.QuoteMeta(createQuery)).WithArgs(
		productBatchMock.BatchNumber,
		productBatchMock.CurrentQuantity,
		productBatchMock.CurrentTemperature,
		productBatchMock.DueDate,
		productBatchMock.InitialQuantity,
		productBatchMock.ManufacturingDate,
		productBatchMock.ManufacturingHour,
		productBatchMock.MinimumTemperature,
		productBatchMock.ProductId,
		productBatchMock.SectionId,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	pbRepo := NewRepository(db)
	result, err := pbRepo.Create(
		context.TODO(),
		1,
		2,
		3,
		"2020-01-01",
		4,
		"2020-01-01",
		5,
		6,
		7,
		8,
	)

	assert.NoError(t, err)
	assert.Equal(t, result.BatchNumber, 1)
}

func TestRepository_Create_Conflict(t *testing.T) {
	db, mock, err := sqlmock.New()
	pbRepo := NewRepository(db)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(createQuery)).
		WithArgs(0, 0, 0, "", 0, "", 0, 0, 0, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = pbRepo.Create(
		context.TODO(),
		1,
		2,
		3,
		"2020-01-01",
		4,
		"2020-01-01",
		5,
		6,
		7,
		8,
	)

	assert.Error(t, err)
}
