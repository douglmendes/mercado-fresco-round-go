package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	getAllSql := queryGetAll
	rows, err := r.db.Query(getAllSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, err
	}
	employees := make([]domain.Employee, 0)
	for rows.Next() {
		var e domain.Employee
		err := rows.Scan(&e.Id, &e.CardNumberId, &e.FirstName, &e.LastName, &e.WarehouseId)
		if err != nil {
			log.Println("Error while scanning employees " + err.Error())
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil

}

func (r *repository) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	row := r.db.QueryRow(queryGetById, id)
	var e domain.Employee
	err := row.Scan(&e.Id, &e.CardNumberId, &e.FirstName, &e.LastName, &e.WarehouseId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			return nil, fmt.Errorf("Employee %d not found", id)
		} else {
			logger.Error(ctx, store.GetPathWithLine(), err.Error())
			return nil, fmt.Errorf("unexpected database error")
		}
	}
	return &e, nil
}

func (r *repository) Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	result, err := r.db.Exec(queryCreate, cardNumberId, firstName, lastName, warehouseId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error retrieving id %d ", id)
	}
	e := domain.Employee{id, cardNumberId, firstName, lastName, warehouseId}
	return &e, nil

}

func (r *repository) Update(ctx context.Context, id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {

	emp, err := r.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("employee %d not found", id)
	}
	if cardNumberId != "" {
		emp.CardNumberId = cardNumberId
	}
	//TODO validar firstName
	if firstName != "" {
		emp.FirstName = firstName
	}
	if lastName != "" {
		emp.LastName = lastName
	}
	if warehouseId != 0 {
		emp.WarehouseId = warehouseId
	}

	_, err = r.db.Exec(queryUpdate, emp.CardNumberId, emp.FirstName, emp.LastName, emp.WarehouseId, id)

	return emp, nil

}

func (r *repository) Delete(ctx context.Context, id int64) error {

	_, err := r.db.Exec(queryDelete, id)
	if err != nil {
		return fmt.Errorf("error when deleting employee %d", id)
	}

	return nil
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}

}
