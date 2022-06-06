package sections

import "fmt"

var database []Section

type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
}

type repository struct {
}

func (r *repository) GetAll() ([]Section, error) {
	return database, nil
}

func (r *repository) GetById(id int) (Section, error) {
	for _, section := range database {
		if section.Id == id {
			return section, nil
		}
	}

	return Section{}, fmt.Errorf("section %d not found in database", id)
}

func NewRepository() Repository {
	database = append(database, Section{Id: 1})
	return &repository{}
}
