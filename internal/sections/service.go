package sections

type Service interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId,
		currentTemperature, minimumTemperature int,
	) (Section, error)
	Update(id int, args map[string]int) (Section, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Section, error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (Section, error) {
	return s.repository.GetById(id)
}

func (s *service) Create(
	sectionNumber, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseId, productTypeId,
	currentTemperature, minimumTemperature int,
) (Section, error) {
	sections, err := s.repository.GetAll()
	if err != nil {
		return Section{}, err
	}

	for _, section := range sections {
		if section.SectionNumber == sectionNumber {
			return Section{}, &ErrorConflict{sectionNumber}
		}
	}

	return s.repository.Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId,
		currentTemperature, minimumTemperature,
	)
}

func (s *service) Update(id int, args map[string]int) (Section, error) {
	err := s.repository.Exists(id)
	if err != nil {
		return Section{}, err
	}

	if sectionNumber := args["section_number"]; sectionNumber != 0 {
		sections, err := s.repository.GetAll()
		if err != nil {
			return Section{}, err
		}

		for _, section := range sections {
			if section.SectionNumber == sectionNumber {
				return Section{}, &ErrorConflict{sectionNumber}
			}
		}
	}

	return s.repository.Update(id, args)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{r}
}
