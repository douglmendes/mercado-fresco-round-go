package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	var sellers []domain.Seller
	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return []domain.Seller{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var seller domain.Seller

		if err := rows.Scan(
			&seller.ID,
			&seller.Cid,
			&seller.CompanyName,
			&seller.Address,
			&seller.Telephone,
		); err != nil {
			return sellers, err
		}

		sellers = append(sellers, seller)
	}
	return sellers, nil
}

func (r *repository) GetById(ctx context.Context,id int) (domain.Seller, error) {

	row := r.db.QueryRowContext(ctx, queryGetById, id)

	seller := domain.Seller{}

	err := row.Scan(
		&seller.ID,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return seller, fmt.Errorf("seller %d not found", id)
	}

	if err != nil {
		return seller, err
	}

	return seller, nil
}

func (r *repository) Create(ctx context.Context, cid int, commpanyName, address, telephone string) (domain.Seller, error) {

	seller := domain.Seller{
		Cid: cid,
		CompanyName: commpanyName,
		Address: address,
		Telephone: telephone,
	}

	result, err := r.db.ExecContext(
		ctx,
		queryCreate,
		cid,
		commpanyName,
		address,
		telephone,
	)

	if err != nil {
		return domain.Seller{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Seller{}, err
	}

	seller.ID = int(lastID)

	return seller, nil
}

func (r *repository) Update(ctx context.Context, id, cid int, commpanyName, address, telephone string) (domain.Seller, error) {

	seller, err := r.GetById(ctx, id)
	if err != nil {
		return domain.Seller{}, fmt.Errorf("seller %d not found", id)
	}

	if cid != 0 {
		seller.Cid = cid
	}
	if commpanyName != "" {
		seller.CompanyName = commpanyName
	}
	if address != "" {
		seller.Address = address
	}
	if telephone != "" {
		seller.Telephone = telephone
	}

	result, err := r.db.ExecContext(
		ctx,
		queryUpdate,
		seller.Cid,
		seller.CompanyName,
		seller.Address,
		seller.Telephone,
		id,
	)
	if err != nil {
		return domain.Seller{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Seller{}, err
	}
	log.Println(affectedRows)

	return seller, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	result, err := r.db.ExecContext(ctx, queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return fmt.Errorf("seller %d not found", id)
	}

	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}