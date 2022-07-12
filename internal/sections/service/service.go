package service

import "github.com/douglmendes/mercado-fresco-round-go/internal/sections/domain"

type service struct {
	repository domain.Repository
}

func (s *service) GetAll() ([]domain.Section, error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (*domain.Section, error) {
	return s.repository.GetById(id)
}

func (s *service) Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (*domain.Section, error) {
	sections, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, section := range sections {
		if section.SectionNumber == sectionNumber {
			return nil, &domain.ErrorConflict{SectionNumber: sectionNumber}
		}
	}

	return s.repository.Create(
		sectionNumber, currentTemperature, minimumTemperature,
		currentCapacity, minimumCapacity, maximumCapacity,
		warehouseId, productTypeId,
	)
}

func (s *service) Update(id int, args map[string]int) (*domain.Section, error) {
	err := s.repository.Exists(id)
	if err != nil {
		return nil, err
	}

	if sectionNumber := args["section_number"]; sectionNumber != 0 {
		sections, err := s.repository.GetAll()
		if err != nil {
			return nil, err
		}

		for _, section := range sections {
			if section.SectionNumber == sectionNumber {
				return nil, &domain.ErrorConflict{SectionNumber: sectionNumber}
			}
		}
	}

	return s.repository.Update(id, args)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r domain.Repository) domain.Service {
	return &service{r}
}
