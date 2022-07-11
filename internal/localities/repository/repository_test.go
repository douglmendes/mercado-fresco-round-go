package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetAll_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := []domain.Locality{
		{
			Id:           1,
			ZipCode:      "12342123",
			LocalityName: "Gasp",
			ProvinceName: "Maceió",
			CountryName:  "Brasil",
		},
		{
			Id:           2,
			ZipCode:      "12342134",
			LocalityName: "Livreiro",
			ProvinceName: "Olinda",
			CountryName:  "Brasil",
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "zip_code", "locality_name", "province_name", "country_name",
	}).AddRow(
		localityMock[0].Id,
		localityMock[0].ZipCode,
		localityMock[0].LocalityName,
		localityMock[0].ProvinceName,
		localityMock[0].CountryName,
	).AddRow(
		localityMock[1].Id,
		localityMock[1].ZipCode,
		localityMock[1].LocalityName,
		localityMock[1].ProvinceName,
		localityMock[1].CountryName,
	)

	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "Gasp", result[0].LocalityName)
	assert.Equal(t, len(result), 2)
}

func TestRepository_GetAll_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := []domain.Locality{}

	mock.ExpectQuery(queryGetAll).WillReturnError(errors.New("erro"))

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetAll(context.TODO())
	assert.Error(t, err)
	assert.Equal(t, localityMock, result)
}

func TestRepository_GetById_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := domain.Locality{
		Id:           1,
		ZipCode:      "54334212",
		LocalityName: "Lux",
		ProvinceName: "Aracaju",
		CountryName:  "Brasil",
	}

	row := sqlmock.NewRows([]string{
		"id", "zip_code", "locality_name", "province_name", "country_name",
	}).AddRow(
		localityMock.Id,
		localityMock.ZipCode,
		localityMock.LocalityName,
		localityMock.ProvinceName,
		localityMock.CountryName,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(row)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetById(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Lux", result.LocalityName)
}

func TestRepository_GetById_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(errors.New("error"))

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetById(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, domain.Locality{}, result)
}

func TestRepository_GetById_NoId(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(sql.ErrNoRows)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetById(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, domain.Locality{}, result)
}

func TestRepository_GetBySellers_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := domain.SellersByLocality{
		LocalityId:   1,
		LocalityName: "Lux",
		SellersCount: 2,
	}

	row := sqlmock.NewRows([]string{
		"id", "locality_name", "sellers_count",
	}).AddRow(
		localityMock.LocalityId,
		localityMock.LocalityName,
		localityMock.SellersCount,
	)

	mock.ExpectQuery(regexp.QuoteMeta(queryGetBySeller)).WillReturnRows(row)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetBySellers(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Lux", result[len(result)-1].LocalityName)
}

func TestRepository_GetBySellers_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(queryGetBySeller)).WillReturnError(sql.ErrNoRows)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetBySellers(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, []domain.SellersByLocality(nil), result)
}

func TestRepository_GetBySellers_NoLocality(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(queryGetBySeller)).WillReturnError(errors.New("error"))

	lcRepo := NewRepository(db)

	_, err = lcRepo.GetBySellers(context.TODO(), 1)
	assert.Error(t, err)
}

func TestRepository_GetBySellers_NoId_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := domain.SellersByLocality{
		LocalityId:   1,
		LocalityName: "Lux",
		SellersCount: 2,
	}

	row := sqlmock.NewRows([]string{
		"id", "locality_name", "sellers_count",
	}).AddRow(
		localityMock.LocalityId,
		localityMock.LocalityName,
		localityMock.SellersCount,
	)

	mock.ExpectQuery(regexp.QuoteMeta(queryGetBySellers)).WillReturnRows(row)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetBySellers(context.TODO(), 0)
	assert.NoError(t, err)
	assert.Equal(t, "Lux", result[len(result)-1].LocalityName)
}

func TestRepository_GetBySellers_NoId_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(queryGetBySellers)).WillReturnError(errors.New("error"))

	lcRepo := NewRepository(db)

	_, err = lcRepo.GetBySellers(context.TODO(), 0)
	assert.Error(t, err)
}

func TestRepository_Create_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := domain.Locality{
		ZipCode:      "54365212",
		LocalityName: "Lux",
		ProvinceName: "Aracaju",
		CountryName:  "Brasil",
	}

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
		localityMock.ZipCode,
		localityMock.LocalityName,
		localityMock.ProvinceName,
		localityMock.CountryName,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	lcRepo := NewRepository(db)

	result, err := lcRepo.Create(context.TODO(), "54365212", "Lux", "Aracaju", "Brasil")
	assert.NoError(t, err)

	assert.Equal(t, "Lux", result.LocalityName)
}

func TestRepository_Create_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).
		WithArgs(0, 0, 0, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	lcRepo := NewRepository(db)

	_, err = lcRepo.Create(context.TODO(), "54365212", "Lux", "Aracaju", "Brasil")

	assert.Error(t, err)

}

func TestRepository_GetByCarriers_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	localityMock := domain.CarriersByLocality{
		LocalityId:    1,
		LocalityName:  "NEXUS",
		CarriersCount: 10,
	}

	row := sqlmock.NewRows([]string{
		"id", "locality_name", "carriers_count",
	}).AddRow(
		localityMock.LocalityId,
		localityMock.LocalityName,
		localityMock.CarriersCount,
	)

	mock.ExpectQuery(regexp.QuoteMeta(queryGetByCarrier)).WillReturnRows(row)

	lcRepo := NewRepository(db)

	result, err := lcRepo.GetByCarriers(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "NEXUS", result[0].LocalityName)
}
