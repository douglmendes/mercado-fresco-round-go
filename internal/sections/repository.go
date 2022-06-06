package sections

import "fmt"

var database []Section

type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	LastID() (int, error)
	Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId int,
		currentTemperature, minimumTemperature float64,
	) (Section, error)
	Delete(id int) error
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

func (r *repository) LastID() (int, error) {
	count := len(database)
	if count == 0 {
		return 0, nil
	}

	return database[count-1].Id, nil
}

func (r *repository) Create(
	sectionNumber, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseId, productTypeId int,
	currentTemperature, minimumTemperature float64,
) (Section, error) {
	lastID, err := r.LastID()
	if err != nil {
		return Section{}, err
	}

	section := Section{
		Id:                 lastID + 1,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}

	database = append(database, section)
	return section, nil
}

func (r *repository) Delete(id int) error {
	for i, section := range database {
		if section.Id == id {
			database = append(database[:i], database[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("section %d not found in database", id)
}

func NewRepository() Repository {
	return &repository{}
}
