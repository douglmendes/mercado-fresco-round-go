package buyers

import "github.com/douglmendes/mercado-fresco-round-go/pkg/store"

var emp []Employee = []Employee{}

type Repository interface {
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	Store(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (Employee, error)
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


func (r *repository) GetById(id int) (Employee, error){
	var emp []Employee
	if err := r.db.Read(&emp) != nil {
		
	}
}