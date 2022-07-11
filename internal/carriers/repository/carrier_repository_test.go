package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_Create_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	carrierMock := domain.Carrier{
		Cid:         "1010",
		CompanyName: "Kato Canecas LTDA",
		Address:     "Wall Street, 20",
		Telephone:   "55111444111",
		LocalityId:  25,
	}

	mock.ExpectExec(regexp.QuoteMeta(sqlCreateCarrier)).WithArgs(
		carrierMock.Cid,
		carrierMock.CompanyName,
		carrierMock.Address,
		carrierMock.Telephone,
		carrierMock.LocalityId,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	carrierRepo := NewRepository(db)

	result, err := carrierRepo.Create(
		context.TODO(),
		"1010",
		"Kato Canecas LTDA",
		"Wall Street, 20",
		"55111444111",
		25,
	)
	assert.NoError(t, err)

	assert.Equal(t, result.Cid, "1010")
}

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	carriersMock := []domain.Carrier{
		{
			Id:          1,
			Cid:         "ABC",
			CompanyName: "Company 1",
			Address:     "Rua teste 1",
			Telephone:   "0000000001",
			LocalityId:  1,
		},
		{
			Id:          2,
			Cid:         "DEF",
			CompanyName: "Company2",
			Address:     "Rua Teste 2",
			Telephone:   "0000000002",
			LocalityId:  1,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow(
		carriersMock[0].Id,
		carriersMock[0].Cid,
		carriersMock[0].CompanyName,
		carriersMock[0].Address,
		carriersMock[0].Telephone,
		carriersMock[0].LocalityId,
	).AddRow(
		carriersMock[1].Id,
		carriersMock[1].Cid,
		carriersMock[1].CompanyName,
		carriersMock[1].Address,
		carriersMock[1].Telephone,
		carriersMock[1].LocalityId,
	)

	mock.ExpectQuery(sqlGetAllCarriers).WillReturnRows(rows)

	whRepo := NewRepository(db)

	result, err := whRepo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, result[0].Cid, "ABC")
	assert.Equal(t, len(result), 2)

}
