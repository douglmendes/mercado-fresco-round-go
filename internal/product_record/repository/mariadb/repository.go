package mariadb

import (
	"context"
	"database/sql"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.ProductRecordRepository {
	return &repository{db}
}

func (r repository) GetByProductId(ctx context.Context, productId int) ([]domain.ProductRecordCount, error) {
	productRecords := []domain.ProductRecordCount{}

	if productId != 0 {
		row := r.db.QueryRowContext(ctx, GetAllGroupByProductIdWhereIdQuery, productId)

		productRecordCount := domain.ProductRecordCount{}

		err := row.Scan(
			&productRecordCount.ProductId,
			&productRecordCount.Description,
			&productRecordCount.RecordsCount,
		)
		if err != nil {
			return productRecords, err
		}

		productRecords = append(productRecords, productRecordCount)
	} else {
		rows, err := r.db.Query(GetAllGroupByProductIdQuery)
		if err != nil {
			return productRecords, err
		}

		for rows.Next() {
			productRecordCount := domain.ProductRecordCount{}

			err := rows.Scan(
				&productRecordCount.ProductId,
				&productRecordCount.Description,
				&productRecordCount.RecordsCount,
			)
			if err != nil {
				return productRecords, err
			}

			productRecords = append(productRecords, productRecordCount)
		}
	}

	return productRecords, nil
}

func (r repository) Create(ctx context.Context, arg domain.ProductRecord) (domain.ProductRecord, error) {
	result, err := r.db.ExecContext(
		ctx,
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
