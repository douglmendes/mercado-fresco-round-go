package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/stretchr/testify/assert"
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
