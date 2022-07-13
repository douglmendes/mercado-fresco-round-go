package repository

import (
	"context"
	"database/sql"

	"github.com/douglmendes/mercado-fresco-round-go/internal/product_batches/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.ProductBatchesRepository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAll(ctx context.Context) ([]domain.ProductBatch, error) {
	var productBatches []domain.ProductBatch

	rows, err := r.db.QueryContext(ctx, getQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product_batch domain.ProductBatch

		if err := rows.Scan(
			&product_batch.Id,
			&product_batch.BatchNumber,
			&product_batch.CurrentQuantity,
			&product_batch.CurrentTemperature,
			&product_batch.DueDate,
			&product_batch.InitialQuantity,
			&product_batch.ManufacturingDate,
			&product_batch.ManufacturingHour,
			&product_batch.MinimumTemperature,
			&product_batch.ProductId,
			&product_batch.SectionId,
		); err != nil {
			return productBatches, err
		}

		productBatches = append(productBatches, product_batch)
	}

	return productBatches, nil
}

func (r repository) Create(ctx context.Context, batchNumber, currentQuantity, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour, minimumTemperature, productId, sectionId int) (*domain.ProductBatch, error) {
	product_batch := domain.ProductBatch{
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: minimumTemperature,
		ProductId:          productId,
		SectionId:          sectionId,
	}

	result, err := r.db.ExecContext(ctx, createQuery, &batchNumber, &currentQuantity, &currentTemperature, &dueDate, &initialQuantity, &manufacturingDate, &manufacturingHour, &minimumTemperature, &productId, &sectionId)
	if err != nil {
		return nil, err
	}

	incrementId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product_batch.Id = int(incrementId)

	return &product_batch, nil
}

func (r repository) GetBySectionId(ctx context.Context, sectionId int) ([]domain.SectionRecords, error) {
	records := []domain.SectionRecords{}
	if sectionId == 0 {
		mappedRecords := map[int]domain.SectionRecords{}

		rows, err := r.db.QueryContext(ctx, allSectionsReportQuery)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var section_id int
			var current_quantity int
			var section_number int

			if err := rows.Scan(
				&current_quantity,
				&section_id,
				&section_number,
			); err != nil {
				return nil, err
			}

			if _, ok := mappedRecords[section_id]; !ok {
				mappedRecords[section_id] = domain.SectionRecords{
					SectionId:     section_id,
					SectionNumber: section_number,
					ProductsCount: current_quantity,
				}
			} else {
				updatedRecord := mappedRecords[section_id]
				updatedRecord.ProductsCount += current_quantity
				mappedRecords[section_id] = updatedRecord
			}
		}

		for _, record := range mappedRecords {
			records = append(records, record)
		}

		return records, nil
	}

	rows, err := r.db.QueryContext(ctx, singleSectionReportQuery, sectionId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	total := 0
	sectionNumber := 0
	for rows.Next() {
		var record domain.SectionRecords
		err := rows.Scan(&record.ProductsCount, &record.SectionId, &record.SectionNumber)
		if err != nil {
			return nil, err
		}

		total += record.ProductsCount
		sectionNumber = record.SectionNumber
	}

	records = append(records, domain.SectionRecords{
		ProductsCount: total,
		SectionId:     sectionId,
		SectionNumber: sectionNumber,
	})

	return records, nil
}
