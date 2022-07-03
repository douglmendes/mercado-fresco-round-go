package mariadb

import (
	"database/sql"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.ProductRepository {
	return &repository{db}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	products := []domain.Product{}

	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		product := domain.Product{}

		err := rows.Scan(
			&product.Id,
			&product.ProductCode,
			&product.Description,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemperature,
			&product.FreezingRate,
			&product.ProductTypeId,
			&product.SellerId,
		)
		if err != nil {
			return []domain.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	// TODO: implementation
	return domain.Product{}, nil
}

func (r *repository) Create(arg domain.Product) (domain.Product, error) {
	// TODO: implementation
	return domain.Product{}, nil
}

func (r *repository) LastID() (int, error) {
	// TODO: implementation
	return 0, nil
}

func (r *repository) Update(arg domain.Product) (domain.Product, error) {
	// TODO: implementation
	return domain.Product{}, nil
}

func (r *repository) Delete(id int) error {
	// TODO: implementation
	return nil
}
