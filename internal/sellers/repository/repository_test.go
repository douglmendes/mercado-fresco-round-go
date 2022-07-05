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

const (
	queryGetAll  = "SELECT id, cid, company_name, address, telephone FROM sellers"
	queryGetById = "SELECT id, cid, company_name, address, telephone FROM sellers where id = ?"
	queryCreate = "INSERT INTO sellers (cid, company_name, address, telephone) VALUES (?, ?, ?, ?)"
	queryUpdate = "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ? WHERE id = ?"
	queryDelete = "DELETE FROM sellers WHERE id = ?"
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
		},
		{
			ID:          2,
			Cid:         49,
			CompanyName: "Leaf",
			Address:     "Rua Arvoredo, 32",
			Telephone:   "98923425",
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone",
	}).AddRow(
		sellerMock[0].ID,
		sellerMock[0].Cid,
		sellerMock[0].CompanyName,
		sellerMock[0].Address,
		sellerMock[0].Telephone,
	).AddRow(
		sellerMock[1].ID,
		sellerMock[1].Cid,
		sellerMock[1].CompanyName,
		sellerMock[1].Address,
		sellerMock[1].Telephone,
	)

	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	slRepo := NewRepository(db)

	result, err := slRepo.GetAll(context.Background())
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

	result, err := slRepo.GetAll(context.Background())
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
		"id", "cid", "company_name", "address", "telephone",
	}).AddRow(
		sellerMock.ID,
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(row)

	slRepo := NewRepository(db)

	result, err := slRepo.GetById(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 44, result.Cid)
}

func TestRepository_GetById_NOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(errors.New("error"))

	slRepo := NewRepository(db)

	result, err := slRepo.GetById(context.Background(), 1)
	assert.Error(t, err)
	assert.Equal(t, domain.Seller{}, result)
}

func TestRepository_GetById_NoId(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryGetById).WillReturnError(sql.ErrNoRows)

	slRepo := NewRepository(db)

	result, err := slRepo.GetById(context.Background(), 1)
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
	}

	mock.ExpectExec(regexp.QuoteMeta(queryCreate)).WithArgs(
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	slRepo := NewRepository(db)

	result, err := slRepo.Create(context.TODO(), 44, "Gasp", "Rua Gaspar, 101", "23225422")
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

	_, err = slRepo.Create(context.TODO(), 44, "Gasp", "Rua Gaspar, 101", "23225422")

	assert.Error(t, err)
}

func TestRepository_Update_Ok(t *testing.T)  {

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
	
	sellerMockUpdated := domain.Seller{
			ID:          1,
			Cid:         44,
			CompanyName: "Gasp",
			Address:     "Rua Gaspar, 111",
			Telephone:   "23222222",
	}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone",
	}).AddRow(
		sellerMock.ID,
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).WithArgs(
		sellerMockUpdated.Cid,
		sellerMockUpdated.CompanyName,
		sellerMockUpdated.Address,
		sellerMockUpdated.Telephone,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	slRepo := NewRepository(db)

	result, err := slRepo.Update(context.TODO(), 1, 44, "Gasp", "Rua Gaspar, 111", "23222222")
	assert.NoError(t, err)
	assert.Equal(t, sellerMockUpdated, result)
}

func TestRepository_Update_NOk(t *testing.T)  {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryUpdate).WillReturnError(errors.New("seller not found"))

	slRepo := NewRepository(db)

	_, err = slRepo.Update(context.TODO(), 1, 44, "Gasp", "Rua Gaspar, 111", "23222222")
	assert.Error(t, err)
}

func TestRepository_Update_ExecContextError(t *testing.T)  {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	
	sellerMockUpdated := domain.Seller{
			ID:          1,
			Cid:         44,
			CompanyName: "Gasp",
			Address:     "Rua Gaspar, 111",
			Telephone:   "23222222",
	}

	sellerMock := domain.Seller{
		ID:          1,
		Cid:         44,
		CompanyName: "Gasp",
		Address:     "Rua Gaspar, 101",
		Telephone:   "23225422",
}

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone",
	}).AddRow(
		sellerMock.ID,
		sellerMock.Cid,
		sellerMock.CompanyName,
		sellerMock.Address,
		sellerMock.Telephone,
	)

	mock.ExpectQuery(queryGetById).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(queryUpdate)).WithArgs(
		sellerMockUpdated.CompanyName,
		sellerMockUpdated.Address,
		sellerMockUpdated.Telephone,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	slRepo := NewRepository(db)

	_, err = slRepo.Update(context.TODO(), 1, 44, "Gasp", "Rua Gaspar, 111", "23222222")
	assert.Error(t, err)
}

func TestRepository_Delete_Ok(t *testing.T)  {
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

func TestRepository_Delete_NOk(t *testing.T)  {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryDelete).WillReturnError(errors.New("error"))

	slRepo := NewRepository(db)

	err = slRepo.Delete(context.TODO(), 1)
	assert.Error(t, err)
}

func TestRepository_Delete_NoId(t *testing.T)  {
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