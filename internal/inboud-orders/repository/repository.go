package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll(ctx context.Context) ([]domain.InboudOrder, error) {

	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		log.Println("Error while querying inboud orders table" + err.Error())
		return nil, err
	}
	io := make([]domain.InboudOrder, 0)
	for rows.Next() {
		var i domain.InboudOrder
		err := rows.Scan(&i.Id, &i.OrderDate, &i.OrderNumber, &i.EmployeeId, &i.ProductBatchId, &i.WarehouseId)
		if err != nil {
			log.Println("Error while scanning inbound orders " + err.Error())
			return nil, err
		}
		io = append(io, i)
	}
	return io, nil

}

func (r *repository) GetByEmployee(ctx context.Context, employee int64) ([]domain.EmployeeInboudOrder, error) {
	io := make([]domain.EmployeeInboudOrder, 0)
	if employee != 0 {
		row := r.db.QueryRow(queryGetByEmplyee, employee)
		var i domain.EmployeeInboudOrder
		err := row.Scan(&i.Id, &i.CardNumberId, &i.FirstName, &i.LastName, &i.WarehouseId, &i.InboudOrderCount)
		if err != nil {
			log.Println("Error while scanning inbound orders " + err.Error())
			return nil, err
		}
		io = append(io, i)
	} else {
		rows, err := r.db.Query(queryGetByEmplyee)
		if err != nil {
			log.Println("Error while querying inboud orders table" + err.Error())
			return nil, err
		}
		for rows.Next() {
			var i domain.EmployeeInboudOrder
			err := rows.Scan(&i.Id, &i.CardNumberId, &i.FirstName, &i.LastName, &i.WarehouseId, &i.InboudOrderCount)
			if err != nil {
				log.Println("Error while scanning inbound orders " + err.Error())
				return nil, err
			}
			io = append(io, i)
		}
	}
	return io, nil
}

func (r *repository) Create(ctx context.Context, orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (*domain.InboudOrder, error) {
	result, err := r.db.Exec(queryCreate, orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error retrieving id %d", id)
	}
	e := domain.InboudOrder{int(id), orderDate, orderNumber, employeeId, productBatchId, warehouseId}
	return &e, nil

}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}

}
