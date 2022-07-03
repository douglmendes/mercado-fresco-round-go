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
	row := r.db.QueryRow(GetByIdQuery, id)

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

func (r *repository) Create(arg domain.Product) (domain.Product, error) {
	result, err := r.db.Exec(
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

func (r *repository) Update(arg domain.Product) (domain.Product, error) {
	_, err := r.db.Exec(
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

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec(DeleteQuery, id)

	return err
}
