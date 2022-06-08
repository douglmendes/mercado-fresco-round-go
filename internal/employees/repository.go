package employees

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	LastID() (int, error)
	Create(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error)
	Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Employee, error) {
	var emp []Employee
	if err := r.db.Read(&emp); err != nil {
		return []Employee{}, nil
	}
	return emp, nil
}

func (r *repository) GetById(id int) (Employee, error) {
	var emp []Employee
	if err := r.db.Read(&emp); err != nil {
		return Employee{}, nil
	}

	for i := range emp {
		if emp[i].Id == id {
			return emp[i], nil
		}
	}
	return Employee{}, fmt.Errorf("Employee %d no found", id)
}

func (r *repository) LastID() (int, error) {
	var emp []Employee
	if err := r.db.Read(&emp); err != nil {
		return 0, err
	}
	if len(emp) == 0 {
		return 0, nil
	}
	return emp[len(emp)-1].Id, nil
}

func (r *repository) Create(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error) {
	var emp []Employee
	if err := r.db.Read(&emp); err != nil {
		return Employee{}, err
	}
	s := Employee{id, cardNumberId, firstName, lastName, warehouseId}
	emp = append(emp, s)
	if err := r.db.Write(emp); err != nil {
		return Employee{}, err
	}
	return s, nil
}

func (r *repository) Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error) {
	var emp []Employee
	if err := r.db.Read(&emp); err != nil {
		return Employee{}, nil
	}

	e := Employee{}

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

			fmt.Printf("employee is %v", e)
			emp[i] = e
			updated = true
			if err := r.db.Write(emp); err != nil {
				return Employee{}, err
			}
		}
	}

	if !updated {
		return Employee{}, fmt.Errorf("employee %d not found", id)
	}
	return e, nil

}

func (r *repository) Delete(id int) error {
	var emp []Employee
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
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}

}
