package sections

import "fmt"

type Service interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId int,
		currentTemperature, minimumTemperature float64,
	) (Section, error)
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
	maximumCapacity, warehouseId, productTypeId int,
	currentTemperature, minimumTemperature float64,
) (Section, error) {
	sections, err := s.repository.GetAll()
	if err != nil {
		return Section{}, err
	}

	for _, section := range sections {
		if section.SectionNumber == sectionNumber {
			return Section{}, fmt.Errorf("a section with number %d already exists", sectionNumber)
		}
	}

	return s.repository.Create(
		sectionNumber, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseId, productTypeId,
		currentTemperature, minimumTemperature,
	)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{r}
}
