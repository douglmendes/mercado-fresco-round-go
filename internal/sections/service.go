package sections

type Service interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
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

func NewService(r Repository) Service {
	return &service{r}
}
