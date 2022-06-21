package sections

import "github.com/douglmendes/mercado-fresco-round-go/pkg/store"

//go:generate mockgen -source=./repository.go -destination=./mock/repository_mock.go
type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (*Section, error)
	LastID() (int, error)
	Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId,
		currentTemperature, minimumTemperature int,
	) (*Section, error)
	Exists(id int) error
	Update(id int, args map[string]int) (*Section, error)
	Delete(id int) (*Section, error)
}

type repository struct {
	database store.Store
}

func (r *repository) GetAll() ([]Section, error) {
	var data []Section

	if err := r.database.Read(&data); err != nil {
		return []Section{}, err
	}

	return data, nil
}

func (r *repository) GetById(id int) (*Section, error) {
	data, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, section := range data {
		if section.Id == id {
			return &section, nil
		}
	}

	return nil, &ErrorNotFound{id}
}

func (r *repository) LastID() (int, error) {
	data, err := r.GetAll()
	if err != nil {
		return 0, err
	}

	count := len(data)
	if count == 0 {
		return 0, nil
	}

	return data[count-1].Id, nil
}

func (r *repository) Create(
	sectionNumber, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseId, productTypeId,
	currentTemperature, minimumTemperature int,
) (*Section, error) {
	lastID, err := r.LastID()
	if err != nil {
		return nil, err
	}

	data, err := r.GetAll()
	if err != nil {
		return nil, err
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

	data = append(data, section)

	if err := r.database.Write(data); err != nil {
		return nil, err
	}

	return &section, nil
}

func (r *repository) Delete(id int) (*Section, error) {
	data, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for i, section := range data {
		if section.Id == id {
			data = append(data[:i], data[i+1:]...)
			err := r.database.Write(data)
			if err != nil {
				return nil, err
			}
			return &section, nil
		}
	}

	return nil, &ErrorNotFound{id}
}

func (r *repository) Exists(id int) error {
	_, err := r.GetById(id)
	return err
}

func (r *repository) Update(id int, args map[string]int) (*Section, error) {
	data, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	var selectedSection *Section
	for i, section := range data {
		if section.Id == id {
			selectedSection = &data[i]
			break
		}
	}

	for key, value := range args {
		switch key {
		case "section_number":
			selectedSection.SectionNumber = value
		case "current_temperature":
			selectedSection.CurrentTemperature = value
		case "minimum_temperature":
			selectedSection.MinimumTemperature = value
		case "current_capacity":
			selectedSection.CurrentCapacity = value
		case "minimum_capacity":
			selectedSection.MinimumCapacity = value
		case "maximum_capacity":
			selectedSection.MaximumCapacity = value
		case "warehouse_id":
			selectedSection.WarehouseId = value
		case "product_type_id":
			selectedSection.ProductTypeId = value
		}
	}

	if err := r.database.Write(data); err != nil {
		return nil, err
	}

	return selectedSection, nil
}

func NewRepository(s store.Store) Repository {
	return &repository{s}
}
