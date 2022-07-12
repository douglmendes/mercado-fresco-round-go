package repository

import (
	"database/sql"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"
)

type repository struct {
	database *sql.DB
}

func (r *repository) GetAll() ([]domain.Section, error) {
	var data []domain.Section

	rows, err := r.database.Query(GetAllQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var section domain.Section

		if err := rows.Scan(&section.Id, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseId, &section.ProductTypeId); err != nil {
			return nil, err
		}

		data = append(data, section)
	}

	return data, nil
}

func (r *repository) GetById(id int) (*domain.Section, error) {
	row := r.database.QueryRow(GetByIdQuery, id)

	var section domain.Section

	if err := row.Scan(&section.Id, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseId, &section.ProductTypeId); err != nil {
		return nil, err
	}

	return &section, nil
}

func (r *repository) Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (*domain.Section, error) {
	result, err := r.database.Exec(CreateQuery, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	section := domain.Section{
		Id:                 int(id),
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}

	return &section, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.database.Exec(DeleteQuery, id)
	return err
}

func (r *repository) Exists(id int) error {
	_, err := r.GetById(id)
	return err
}

func (r *repository) Update(id int, args map[string]int) (*domain.Section, error) {
	section, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	for key, value := range args {
		switch key {
		case "section_number":
			section.SectionNumber = value
		case "current_temperature":
			section.CurrentTemperature = value
		case "minimum_temperature":
			section.MinimumTemperature = value
		case "current_capacity":
			section.CurrentCapacity = value
		case "minimum_capacity":
			section.MinimumCapacity = value
		case "maximum_capacity":
			section.MaximumCapacity = value
		case "warehouse_id":
			section.WarehouseId = value
		case "product_type_id":
			section.ProductTypeId = value
		}
	}

	_, err = r.database.Exec(UpdateQuery, section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.MaximumCapacity, section.WarehouseId, section.ProductTypeId, id)
	if err != nil {
		return nil, err
	}

	return section, nil
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{db}
}
