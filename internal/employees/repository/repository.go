package repository

import (
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/employees/domain"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll() ([]domain.Employee, error) {
	getAllSql := "SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employees"
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

func (r *repository) GetById(id int64) (*domain.Employee, error) {
	getByIdSql := "SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employees where id = ?"
	row := r.db.QueryRow(getByIdSql, id)
	var e domain.Employee
	err := row.Scan(&e.Id, &e.CardNumberId, &e.FirstName, &e.LastName, &e.WarehouseId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee %d not found", id)
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, fmt.Errorf("unexpected database error")
		}
	}
	return &e, nil
}

/*
func (r *repository) LastID() (int, error) {
	var emp []domain.Employee
	if err := r.db.Read(&emp); err != nil {
		return 0, err
	}
	if len(emp) == 0 {
		return 0, nil
	}
	return emp[len(emp)-1].Id, nil
}
*/
func (r *repository) Create(cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	createSql := "insert into employees (id_card_number , first_name, last_name, warehouse_id) values (?,?,?,?)"
	result, err := r.db.Exec(createSql, cardNumberId, firstName, lastName, warehouseId)
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

func (r *repository) Update(id int64, cardNumberId string, firstName string, lastName string, warehouseId int) (*domain.Employee, error) {
	updateSql := "UPDATE employees  SET id_card_number  =  ? , first_name= ? , last_name = ? , warehouse_id  = ? WHERE id=?"
	//_, err := r.db.Exec(updateSql, cardNumberId, firstName, lastName, warehouseId, id)

	emp, err := r.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("employee %d not found", id)
	}
	if cardNumberId != "" {
		emp.CardNumberId = cardNumberId
	}
	if firstName == " " {
		emp.FirstName = firstName
	}
	if lastName != "" {
		emp.LastName = lastName
	}
	if warehouseId != 0 {
		emp.WarehouseId = warehouseId
	}

	result, err := r.db.Exec(updateSql, emp.CardNumberId, emp.FirstName, emp.LastName, emp.WarehouseId, id)
	log.Println(result.RowsAffected())

	return emp, nil

}

/*
func (r *repository) Delete(id int) error {
	var emp []domain.Employee
	if err := r.db.Read(&emp); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range emp {
		if emp[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("employee %d not found", id)
	}

	emp = append(emp[:index], emp[index+1:]...)
	if err := r.db.Write(emp); err != nil {
		return err
	}

	return nil
}*/

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}

}
