package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Locality, error) {
	var localities []domain.Locality
	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return []domain.Locality{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var locality domain.Locality

		if err := rows.Scan(
			&locality.Id,
			&locality.LocalityName,
			&locality.ProvinceName,
			&locality.CountryName,
		); err != nil {
			return localities, err
		}

		localities = append(localities, locality)
	}

	return localities, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Locality, error) {
	row := r.db.QueryRowContext(ctx, queryGetById, id)

	locality := domain.Locality{}

	err := row.Scan(
		&locality.Id,
		&locality.LocalityName,
		&locality.ProvinceName,
		&locality.CountryName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return locality, fmt.Errorf("locality %d not found", id)
	}

	if err != nil {
		return locality, err
	}

	return locality, nil
}

func (r *repository) GetBySellers(ctx context.Context, id int) ([]domain.SellersByLocality, error) {

	var sellersByLocality []domain.SellersByLocality

	if id != 0 {
		row := r.db.QueryRowContext(ctx, queryGetBySeller, id)
	
		var sellers domain.SellersByLocality
	
		err := row.Scan(
			&sellers.LocalityId,
			&sellers.LocalityName,
			&sellers.SellersCount,
		)
	
		if errors.Is(err, sql.ErrNoRows) {
			return sellersByLocality, fmt.Errorf("locality %d not found", id)
		}
	
		if err != nil {
			return sellersByLocality, err
		}
	
		sellersByLocality = append(sellersByLocality, sellers)
	} else {
		rows, err := r.db.QueryContext(ctx, queryGetBySellers)
		if err != nil {
			return []domain.SellersByLocality{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var sellers domain.SellersByLocality
			err := rows.Scan(
				&sellers.LocalityId,
				&sellers.LocalityName,
				&sellers.SellersCount,
			)
			if err != nil {
				return sellersByLocality, err
			}

			sellersByLocality = append(sellersByLocality, sellers)
		}
	}
	return sellersByLocality, nil
}

func (r *repository) Create(ctx context.Context, localityName, provinceName, countryName string) (domain.Locality, error) {
	locality := domain.Locality{
		LocalityName: localityName,
		ProvinceName: provinceName,
		CountryName:  countryName,
	}

	result, err := r.db.ExecContext(
		ctx,
		queryCreate,
		localityName,
		provinceName,
		countryName,
	)

	if err != nil {
		return domain.Locality{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Locality{}, err
	}

	locality.Id = int(lastId)

	return locality, nil
}

func (r *repository) Update(ctx context.Context, id int, localityName, provinceName, countryName string) (domain.Locality, error) {

	locality, err := r.GetById(ctx, id)
	if err != nil {
		return domain.Locality{}, fmt.Errorf("locality %d not found", id)
	}

	if localityName != "" {
		locality.LocalityName = localityName
	}
	if provinceName != "" {
		locality.ProvinceName = provinceName
	}
	if countryName != "" {
		locality.CountryName = countryName
	}

	result, err := r.db.ExecContext(
		ctx,
		queryUpdate,
		locality.LocalityName,
		locality.ProvinceName,
		locality.CountryName,
		id,
	)
	if err != nil {
		return domain.Locality{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Locality{}, err
	}

	return locality, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	result, err := r.db.ExecContext(ctx, queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if affectedRows == 0 {
		return fmt.Errorf("locality %d not found", id)
	}

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sql.DB) domain.LocalityRepository {
	return &repository{
		db: db,
	}
}
