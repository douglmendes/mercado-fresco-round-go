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
	// TODO: implementation
	return domain.ProductRecord{}, nil
}
