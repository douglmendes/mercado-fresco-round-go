package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
	"github.com/stretchr/testify/assert"
)

var (
	sampleBatch = domain.ProductBatch{
		Id:                 1,
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
	sampleRecord = domain.SectionRecords{
		SectionId:     1,
		SectionNumber: 2,
		ProductsCount: 5,
	}
)

func TestRepository_Create_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(createQuery)).WithArgs(
		sampleBatch.BatchNumber,
		sampleBatch.CurrentQuantity,
		sampleBatch.CurrentTemperature,
		sampleBatch.DueDate,
		sampleBatch.InitialQuantity,
		sampleBatch.ManufacturingDate,
		sampleBatch.ManufacturingHour,
		sampleBatch.MinimumTemperature,
		sampleBatch.ProductId,
		sampleBatch.SectionId,
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

func TestRepository_Get_All(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	result := sqlmock.NewRows([]string{"id", "batch_number", "current_quantity", "current_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "minimum_temperature", "product_id", "section_id"}).AddRow(
		sampleBatch.Id,
		sampleBatch.BatchNumber,
		sampleBatch.CurrentQuantity,
		sampleBatch.CurrentTemperature,
		sampleBatch.DueDate,
		sampleBatch.InitialQuantity,
		sampleBatch.ManufacturingDate,
		sampleBatch.ManufacturingHour,
		sampleBatch.MinimumTemperature,
		sampleBatch.ProductId,
		sampleBatch.SectionId,
	)

	mock.ExpectQuery(getQuery).WillReturnRows(result)

	repository := NewRepository(db)
	batches, err := repository.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, []domain.ProductBatch{sampleBatch}, batches)
}

func TestRepository_Get_By_Single_Section_Id(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	result := sqlmock.NewRows([]string{"current_quantity", "section_id", "section_number"}).AddRow(
		sampleRecord.ProductsCount,
		sampleRecord.SectionId,
		sampleRecord.SectionNumber,
	)

	mock.ExpectQuery(regexp.QuoteMeta(singleSectionReportQuery)).WillReturnRows(result)

	repository := NewRepository(db)
	records, err := repository.GetBySectionId(context.TODO(), 1)

	assert.NoError(t, err)
	assert.Equal(t, []domain.SectionRecords{sampleRecord}, records)
}

func TestRepository_Get_By_All_Sections(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	result := sqlmock.NewRows([]string{"current_quantity", "section_id", "section_number"}).AddRow(
		sampleRecord.ProductsCount,
		sampleRecord.SectionId,
		sampleRecord.SectionNumber,
	)

	mock.ExpectQuery(regexp.QuoteMeta(allSectionsReportQuery)).WillReturnRows(result)

	repository := NewRepository(db)
	records, err := repository.GetBySectionId(context.TODO(), 0)

	assert.NoError(t, err)
	assert.Equal(t, []domain.SectionRecords{sampleRecord}, records)
}
