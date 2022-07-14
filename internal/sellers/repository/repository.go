package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	var sellers []domain.Seller
	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
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
			&seller.LocalityId,
		); err != nil {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			return sellers, err
		}

		sellers = append(sellers, seller)
	}
	return sellers, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Seller, error) {

	row := r.db.QueryRowContext(ctx, queryGetById, id)

	seller := domain.Seller{}

	err := row.Scan(
		&seller.ID,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityId,
	)

	if errors.Is(err, sql.ErrNoRows) {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return seller, fmt.Errorf("seller %d not found", id)
	}

	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return seller, err
	}

	return seller, nil
}

func (r *repository) Create(ctx context.Context, cid int, commpanyName, address, telephone string, localityId int) (domain.Seller, error) {

	seller := domain.Seller{
		Cid:         cid,
		CompanyName: commpanyName,
		Address:     address,
		Telephone:   telephone,
		LocalityId:  localityId,
	}

	result, err := r.db.ExecContext(
		ctx,
		queryCreate,
		cid,
		commpanyName,
		address,
		telephone,
		localityId,
	)

	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Seller{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Seller{}, err
	}

	seller.ID = int(lastID)

	return seller, nil
}

func (r *repository) Update(ctx context.Context, id, cid int, commpanyName, address, telephone string, localityId int) (domain.Seller, error) {

	seller, err := r.GetById(ctx, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
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
	if localityId != 0 {
		seller.LocalityId = localityId
	}

	result, err := r.db.ExecContext(
		ctx,
		queryUpdate,
		seller.Cid,
		seller.CompanyName,
		seller.Address,
		seller.Telephone,
		seller.LocalityId,
		id,
	)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Seller{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return domain.Seller{}, err
	}
	log.Println(affectedRows)

	return seller, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	result, err := r.db.ExecContext(ctx, queryDelete, id)
	if err != nil {
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		logger.Error(ctx, store.GetPathWithLine(), "seller not found")
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
