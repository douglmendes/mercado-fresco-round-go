package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	locality "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
)

const (
	sqlCreateCarrier  = "INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	sqlGetAllCarriers = "SELECT id, cid, company_name, address, telephone, locality_id FROM carries"
	sqlLocality       = "SELECT carries.id, carries.cid, carries.company_name, carries.address, carries.telephone, carries.locality_id " +
		"FROM carries INNER JOIN localities local ON carries.locality_id = local.id WHERE carries.id = ?"
	sqlGetLocality = "SELECT id, locality_name, province_name, country_name FROM localities WHERE id = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.CarrierRepository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAll(ctx context.Context) ([]domain.Carrier, error) {
	var carrierList []domain.Carrier

	rows, err := r.db.QueryContext(ctx, sqlGetAllCarriers)
	if err != nil {
		return []domain.Carrier{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var carrier domain.Carrier
		if err := rows.Scan(
			&carrier.Id,
			&carrier.Cid,
			&carrier.CompanyName,
			&carrier.Address,
			&carrier.Telephone,
			&carrier.LocalityId,
		); err != nil {
			return carrierList, err
		}
		carrierList = append(carrierList, carrier)
	}
	return carrierList, nil
}

func (r repository) GetLocal(id int) (locality.Locality, error) {
	row := r.db.QueryRow(sqlGetLocality, id)

	local := locality.Locality{}

	err := row.Scan(
		&local.Id,
		&local.LocalityName,
		&local.ProvinceName,
		&local.CountryName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return local, fmt.Errorf("local %d not found", id)
	}

	if err != nil {
		return local, err
	}

	return local, nil
}

func (r repository) Create(ctx context.Context, cid, companyName, address, telephone string, localityId int) (domain.Carrier, error) {
	carrier := domain.Carrier{
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityId:  localityId,
	}

	result, err := r.db.ExecContext(ctx, sqlCreateCarrier, &cid, &companyName, &address, &telephone, &localityId)
	if err != nil {
		return domain.Carrier{}, err
	}

	incrementId, err := result.LastInsertId()
	if err != nil {
		return domain.Carrier{}, err
	}

	carrier.Id = int(incrementId)

	return carrier, nil
}
