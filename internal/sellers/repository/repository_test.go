package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetAll_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sellerMock := []domain.Seller{
		{
			ID:          1,
			Cid:         44,
			CompanyName: "Gasp",
			Address:     "Rua Gaspar, 101",
			Telephone:   "23225422",
			LocalityId:  1,
		},
		{
			ID:          2,
			Cid:         49,
			CompanyName: "Leaf",
			Address:     "Rua Arvoredo, 32",
			Telephone:   "98923425",
			LocalityId:  1,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow(
		sellerMock[0].ID,
		sellerMock[0].Cid,
		sellerMock[0].CompanyName,
		sellerMock[0].Address,
		sellerMock[0].Telephone,
		sellerMock[0].LocalityId,
	).AddRow(
		sellerMock[1].ID,
		sellerMock[1].Cid,
		sellerMock[1].CompanyName,
		sellerMock[1].Address,
		sellerMock[1].Telephone,
		sellerMock[0].LocalityId,
	)

	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	slRepo := NewRepository(db)

	result, err := slRepo.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 44, result[0].Cid)
	assert.Equal(t, len(result), 2)
}

func TestRepository_GetAll_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sellerMock := []domain.Seller{}

	mock.ExpectQuery(queryGetAll).WillReturnError(errors.New("erro"))

	slRepo := NewRepository(db)

	result, err := slRepo.GetAll(context.TODO())
	assert.Error(t, err)
	assert.Equal(t, sellerMock, result)
}

func TestRepository_GetById_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sellerMock := domain.Seller{
		ID:          1,
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 101",
		Telephone:   "23225422",
	}

	row := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow(
		sellerMock.ID,
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
		sellerMock.LocalityId,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(row)

	slRepo := NewRepository(db)

	result, err := slRepo.GetById(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 44, result.Cid)
}

func TestRepository_GetById_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(errors.New("error"))

	slRepo := NewRepository(db)

	result, err := slRepo.GetById(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, domain.Seller{}, result)
}

func TestRepository_GetById_NoId(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(sql.ErrNoRows)

	slRepo := NewRepository(db)

	result, err := slRepo.GetById(context.TODO(), 1)
	assert.Error(t, err)
	assert.Equal(t, domain.Seller{}, result)
}

func TestRepository_Create_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sellerMock := domain.Seller{
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 101",
		Telephone:   "23225422",
		LocalityId:  1,
	}

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
		sellerMock.LocalityId,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	slRepo := NewRepository(db)

	result, err := slRepo.Create(context.TODO(), 44, "Gasp", "Rua Gaspar, 101", "23225422", 1)
	assert.NoError(t, err)

	assert.Equal(t, result.Cid, 44)
}

func TestRepository_Create_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).
		WithArgs(0, 0, 0, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	slRepo := NewRepository(db)

	_, err = slRepo.Create(context.TODO(), 44, "Gasp", "Rua Gaspar, 101", "23225422", 1)

	assert.Error(t, err)
}

func TestRepository_Update_Ok(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sellerMock := domain.Seller{
		ID:          1,
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 101",
		Telephone:   "23225422",
		LocalityId:  1,
	}

	sellerMockUpdated := domain.Seller{
		ID:          1,
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 111",
		Telephone:   "23222222",
		LocalityId:  1,
	}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow(
		sellerMock.ID,
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
		sellerMock.LocalityId,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).WithArgs(
		sellerMockUpdated.Cid,
		sellerMockUpdated.CompanyName,
		sellerMockUpdated.Address,
		sellerMockUpdated.Telephone,
		sellerMockUpdated.LocalityId,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	slRepo := NewRepository(db)

	result, err := slRepo.Update(context.TODO(), 1, 44, "Gasp", "Rua Gaspar, 111", "23222222", 1)
	assert.NoError(t, err)
	assert.Equal(t, sellerMockUpdated, result)
}

func TestRepository_Update_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryUpdate).WillReturnError(errors.New("seller not found"))

	slRepo := NewRepository(db)

	_, err = slRepo.Update(context.TODO(), 1, 44, "Gasp", "Rua Gaspar, 111", "23222222", 1)
	assert.Error(t, err)
}

func TestRepository_Update_ExecContextError(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sellerMockUpdated := domain.Seller{
		ID:          1,
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 111",
		Telephone:   "23222222",
		LocalityId:  1,
	}

	sellerMock := domain.Seller{
		ID:          1,
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 101",
		Telephone:   "23225422",
		LocalityId:  1,
	}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow(
		sellerMock.ID,
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
		sellerMock.LocalityId,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).WithArgs(
		sellerMockUpdated.CompanyName,
		sellerMockUpdated.Address,
		sellerMockUpdated.Telephone,
		sellerMockUpdated.LocalityId,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	slRepo := NewRepository(db)

	_, err = slRepo.Update(context.TODO(), 1, 44, "Gasp", "Rua Gaspar, 111", "23222222", 1)
	assert.Error(t, err)
}

func TestRepository_Delete_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryDelete)).WithArgs(
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	slRepo := NewRepository(db)

	err = slRepo.Delete(context.TODO(), 1)
	assert.NoError(t, err)
}

func TestRepository_Delete_NOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryDelete).WillReturnError(errors.New("error"))

	slRepo := NewRepository(db)

	err = slRepo.Delete(context.TODO(), 1)
	assert.Error(t, err)
}

func TestRepository_Delete_NoId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(queryDelete)).WithArgs(
		1,
	).WillReturnResult(sqlmock.NewResult(0, 0))

	slRepo := NewRepository(db)

	err = slRepo.Delete(context.TODO(), 1)
	assert.Error(t, err)
}
