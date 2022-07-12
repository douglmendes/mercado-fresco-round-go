package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
	"github.com/stretchr/testify/assert"
)

var (
	sampleSection = domain.Section{
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
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(CreateQuery)).WithArgs(
		sampleSection.SectionNumber,
		sampleSection.CurrentTemperature,
		sampleSection.MinimumTemperature,
		sampleSection.CurrentCapacity,
		sampleSection.MinimumCapacity,
		sampleSection.MaximumCapacity,
		sampleSection.WarehouseId,
		sampleSection.ProductTypeId,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewRepository(db)
	section, err := repository.Create(
		sampleSection.SectionNumber,
		sampleSection.CurrentTemperature,
		sampleSection.MinimumTemperature,
		sampleSection.CurrentCapacity,
		sampleSection.MinimumCapacity,
		sampleSection.MaximumCapacity,
		sampleSection.WarehouseId,
		sampleSection.ProductTypeId,
	)

	assert.NoError(t, err)
	assert.Equal(t, sampleSection, *section)
}

func TestRepository_Get_All(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	result := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}).AddRow(
		sampleSection.Id,
		sampleSection.SectionNumber,
		sampleSection.CurrentTemperature,
		sampleSection.MinimumTemperature,
		sampleSection.CurrentCapacity,
		sampleSection.MinimumCapacity,
		sampleSection.MaximumCapacity,
		sampleSection.WarehouseId,
		sampleSection.ProductTypeId,
	)

	mock.ExpectQuery(GetAllQuery).WillReturnRows(result)

	repository := NewRepository(db)
	sections, err := repository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, []domain.Section{sampleSection}, sections)
}

func TestRepository_Get_ById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	result := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}).AddRow(
		sampleSection.Id,
		sampleSection.SectionNumber,
		sampleSection.CurrentTemperature,
		sampleSection.MinimumTemperature,
		sampleSection.CurrentCapacity,
		sampleSection.MinimumCapacity,
		sampleSection.MaximumCapacity,
		sampleSection.WarehouseId,
		sampleSection.ProductTypeId,
	)

	mock.ExpectQuery(regexp.QuoteMeta(GetByIdQuery)).WithArgs(sampleSection.Id).WillReturnRows(result)

	repository := NewRepository(db)
	section, err := repository.GetById(sampleSection.Id)

	assert.NoError(t, err)
	assert.Equal(t, sampleSection, *section)
}

func TestRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	result := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}).AddRow(
		sampleSection.Id,
		sampleSection.SectionNumber,
		sampleSection.CurrentTemperature,
		sampleSection.MinimumTemperature,
		sampleSection.CurrentCapacity,
		sampleSection.MinimumCapacity,
		sampleSection.MaximumCapacity,
		sampleSection.WarehouseId,
		sampleSection.ProductTypeId,
	)

	mock.ExpectQuery(regexp.QuoteMeta(GetByIdQuery)).WithArgs(sampleSection.Id).WillReturnRows(result)

	mock.ExpectExec(regexp.QuoteMeta(UpdateQuery)).WithArgs(
		sampleSection.SectionNumber,
		sampleSection.CurrentTemperature,
		sampleSection.MinimumTemperature,
		88,
		sampleSection.MinimumCapacity,
		sampleSection.MaximumCapacity,
		sampleSection.WarehouseId,
		sampleSection.ProductTypeId,
		sampleSection.Id,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewRepository(db)
	section, err := repository.Update(
		sampleSection.Id,
		map[string]int{
			"current_capacity": 88,
		},
	)

	assert.NoError(t, err)
	updatedSection := sampleSection
	updatedSection.CurrentCapacity = 88
	assert.Equal(t, updatedSection, *section)
}

func TestRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(DeleteQuery)).WithArgs(sampleSection.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewRepository(db)
	err = repository.Delete(sampleSection.Id)

	assert.NoError(t, err)
}
