package mariadb

import (
	"database/sql"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.ProductRecordRepository {
	return &repository{db}
}

func (r repository) GetById(id int) ([]domain.ProductRecord, error) {
	// TODO: implementation
	return []domain.ProductRecord{}, nil
}

func (r repository) Create(arg domain.ProductRecord) (domain.ProductRecord, error) {
	result, err := r.db.Exec(
		CreateQuery,
		arg.LastUpdateDate,
		arg.PurchasePrice,
		arg.SalePrice,
		arg.ProductId,
	)
	if err != nil {
		return domain.ProductRecord{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.ProductRecord{}, err
	}

	productRecord := arg
	productRecord.Id = int(id)

	return productRecord, nil
}
