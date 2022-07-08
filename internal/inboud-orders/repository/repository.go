package repository

import (
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/inboud-orders/domain"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll() ([]domain.InboudOrder, error) {
	getAllSql := "SELECT id,order_date,order_number,employee_id,product_batch_id, warehouse_id FROM inbound_orders"
	rows, err := r.db.Query(getAllSql)
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

func (r *repository) GetByEmployee(employee int64) ([]domain.EmployeeInboudOrder, error) {
	log.Println(employee)
	io := make([]domain.EmployeeInboudOrder, 0)
	if employee != 0 {
		getByIdSql := "Select e.id , " +
			"e.id_card_number , " +
			"e.first_name , " +
			"e.last_name , " +
			"e.warehouse_id ," +
			"count(*) as inbound_orders_count " +
			"from inbound_orders io inner join employees e on e.id = io.employee_id where employee_id = ?"
		row := r.db.QueryRow(getByIdSql, employee)
		var i domain.EmployeeInboudOrder
		err := row.Scan(&i.Id, &i.CardNumberId, &i.FirstName, &i.LastName, &i.WarehouseId, &i.InboudOrderCount)
		if err != nil {
			log.Println("Error while scanning inbound orders " + err.Error())
			return nil, err
		}
		io = append(io, i)
	} else {
		getByIdSql := "Select e.id , " +
			"e.id_card_number , " +
			"e.first_name , " +
			"e.last_name , " +
			"e.warehouse_id ," +
			"count(*) as inbound_orders_count " +
			"from inbound_orders io inner join employees e on e.id = io.employee_id group by e.id"
		rows, err := r.db.Query(getByIdSql)
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

func (r *repository) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (*domain.InboudOrder, error) {
	createSql := "insert into inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) values(?,?,?,?,?)"
	result, err := r.db.Exec(createSql, orderDate, orderNumber, employeeId, productBatchId, warehouseId)
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
