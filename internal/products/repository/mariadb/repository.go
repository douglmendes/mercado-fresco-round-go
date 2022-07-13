package mariadb

import (
	"context"
	"database/sql"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.ProductRepository {
	return &repository{db}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	products := []domain.Product{}

	rows, err := r.db.QueryContext(ctx, GetAllQuery)
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

func (r *repository) GetById(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRowContext(ctx, GetByIdQuery, id)

	product := domain.Product{}

	err := row.Scan(
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
		return domain.Product{}, err
	}

	return product, nil
}

func (r *repository) Create(ctx context.Context, arg domain.Product) (domain.Product, error) {
	result, err := r.db.ExecContext(
		ctx,
		CreateQuery,
		arg.ProductCode,
		arg.Description,
		arg.Width,
		arg.Height,
		arg.Length,
		arg.NetWeight,
		arg.ExpirationRate,
		arg.RecommendedFreezingTemperature,
		arg.FreezingRate,
		arg.ProductTypeId,
		arg.SellerId,
	)
	if err != nil {
		return domain.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}

	product := arg
	product.Id = int(id)

	return product, nil
}

func (r *repository) Update(ctx context.Context, arg domain.Product) (domain.Product, error) {
	_, err := r.db.ExecContext(
		ctx,
		UpdateQuery,
		arg.ProductCode,
		arg.Description,
		arg.Width,
		arg.Height,
		arg.Length,
		arg.NetWeight,
		arg.ExpirationRate,
		arg.RecommendedFreezingTemperature,
		arg.FreezingRate,
		arg.ProductTypeId,
		arg.SellerId,
		arg.Id,
	)
	if err != nil {
		return domain.Product{}, err
	}

	return arg, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, DeleteQuery, id)

	return err
}
