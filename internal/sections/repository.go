package sections

var database []Section

type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	LastID() (int, error)
	Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId,
		currentTemperature, minimumTemperature int,
	) (Section, error)
	Exists(id int) error
	Update(id int, args map[string]int) (Section, error)
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

	return Section{}, &ErrorNotFound{id}
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
	maximumCapacity, warehouseId, productTypeId,
	currentTemperature, minimumTemperature int,
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

	return &ErrorNotFound{id}
}

func (r *repository) Exists(id int) error {
	for _, section := range database {
		if section.Id == id {
			return nil
		}
	}

	return &ErrorNotFound{id}
}

func (r *repository) Update(id int, args map[string]int) (Section, error) {
	var selectedSection *Section
	for i, section := range database {
		if section.Id == id {
			selectedSection = &database[i]
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

	return *selectedSection, nil
}

func NewRepository() Repository {
	return &repository{}
}
