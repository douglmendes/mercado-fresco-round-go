package repository

import (
	"database/sql"
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

/*
func (r *repository) GetById(id int) (domain.Employee, error) {
	var emp []domain.Employee
	if err := r.db.Read(&emp); err != nil {
		return domain.Employee{}, nil
	}

	for i := range emp {
		if emp[i].Id == id {
			return emp[i], nil
		}
	}
	return domain.Employee{}, fmt.Errorf("Employee %d not found", id)
}

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

func (r *repository) Create(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (domain.Employee, error) {
	var emp []domain.Employee
	if err := r.db.Read(&emp); err != nil {
		return domain.Employee{}, err
	}
	s := domain.Employee{id, cardNumberId, firstName, lastName, warehouseId}
	emp = append(emp, s)
	if err := r.db.Write(emp); err != nil {
		return domain.Employee{}, err
	}
	return s, nil
}

func (r *repository) Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (domain.Employee, error) {
	var emp []domain.Employee
	if err := r.db.Read(&emp); err != nil {
		return domain.Employee{}, nil
	}

	e := domain.Employee{}

	updated := false
	for i := range emp {
		if emp[i].Id == id {

			e = emp[i]
			if id != 0 {
				e.Id = id
			}
			if cardNumberId != "" {
				e.CardNumberId = cardNumberId
			}
			if firstName != "" {
				e.FirstName = firstName
			}
			if lastName != "" {
				e.LastName = lastName
			}
			if warehouseId != 0 {
				e.WarehouseId = warehouseId
			}

			fmt.Printf("employee is %v", e)
			emp[i] = e
			updated = true
			if err := r.db.Write(emp); err != nil {
				return domain.Employee{}, err
			}
		}
	}

	if !updated {
		return domain.Employee{}, fmt.Errorf("employee %d not found", id)
	}
	return e, nil

}

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
