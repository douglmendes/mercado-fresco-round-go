package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const (
	queryGetAll = "SELECT id, address, telephone, warehouse_code, locality_id FROM warehouse"
)

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	warehosesMock := []domain.Warehouse{
		{
			Id:            1,
			Address:       "Rua teste 1",
			Telephone:     "0000000001",
			WarehouseCode: "AAA",
			LocalityId:    10,
		},
		{
			Id:            2,
			Address:       "Rua Teste 2",
			Telephone:     "0000000002",
			WarehouseCode: "BBB",
			LocalityId:    20,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "address", "telephone", "warehouse_code", "locality_id",
	}).AddRow(
		warehosesMock[0].Id,
		warehosesMock[0].Address,
		warehosesMock[0].Telephone,
		warehosesMock[0].WarehouseCode,
		warehosesMock[0].LocalityId,
	).AddRow(
		warehosesMock[1].Id,
		warehosesMock[1].Address,
		warehosesMock[1].Telephone,
		warehosesMock[1].WarehouseCode,
		warehosesMock[1].LocalityId,
	)

	mock.ExpectQuery(queryGetAll).WillReturnRows(rows)

	whRepo := NewRepository(db)

	result, err := whRepo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, result[0].WarehouseCode, "AAA")
	assert.Equal(t, len(result), 2)
}
func TestRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	wh := domain.Warehouse{
		Id:            1,
		Address:       "Rua 25 de Março",
		Telephone:     "9911100011",
		WarehouseCode: "XYZ",
		LocalityId:    101,
	}

	row := sqlmock.NewRows([]string{
		"id", "address", "telephone", "warehouse_code", "locality_id",
	}).AddRow(wh.Id, wh.Address, wh.Telephone, wh.WarehouseCode, wh.LocalityId)

	mock.ExpectQuery(sqlGetById).WillReturnRows(row)

	whRepo := NewRepository(db)

	result, err := whRepo.GetById(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, wh.WarehouseCode, result.WarehouseCode)
}

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	whMock := domain.Warehouse{
		Address:       "Rua do Teste",
		Telephone:     "9888100000",
		WarehouseCode: "TEST",
		LocalityId:    1,
	}

	mock.ExpectExec(regexp.QuoteMeta(sqlCreate)).WithArgs(
		whMock.Address,
		whMock.Telephone,
		whMock.WarehouseCode,
		whMock.LocalityId,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	whRepo := NewRepository(db)

	result, err := whRepo.Create(
		context.Background(),
		"Rua do Teste",
		"9888100000",
		"TEST",
		1,
	)

	assert.NoError(t, err)
	assert.Equal(t, result.WarehouseCode, whMock.WarehouseCode)
}

func TestRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	whMock := domain.Warehouse{
		Id:            1,
		Address:       "Rua do Teste",
		Telephone:     "9888100000",
		WarehouseCode: "TEST",
		LocalityId:    1,
	}

	newWhMock := domain.Warehouse{
		Id:            1,
		Address:       "Rua do Teste Nova",
		Telephone:     "911111111",
		WarehouseCode: "TEST",
		LocalityId:    1,
	}

	rows := sqlmock.NewRows([]string{
		"id", "address", "telephone", "warehouse_code", "locality_id",
	}).AddRow(
		whMock.Id,
		whMock.Address,
		whMock.Telephone,
		whMock.WarehouseCode,
		whMock.LocalityId,
	)

	mock.ExpectQuery(sqlGetById).WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).WithArgs(
		newWhMock.Address,
		newWhMock.Telephone,
		newWhMock.WarehouseCode,
		newWhMock.LocalityId,
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	whRepo := NewRepository(db)

	result, err := whRepo.Update(
		context.Background(),
		1,
		"Rua do Teste Nova",
		"911111111",
		"TEST",
		1,
	)

	assert.NoError(t, err)
	assert.Equal(t, newWhMock, result)
}
