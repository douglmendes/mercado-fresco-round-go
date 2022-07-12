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
	t.Skip("Not implemented")
}

func TestRepository_Get_ById(t *testing.T) {
	t.Skip("Not implemented")
}

func TestRepository_Update(t *testing.T) {
	t.Skip("Not implemented")
}

func TestRepository_Delete(t *testing.T) {
	t.Skip("Not implemented")
}
