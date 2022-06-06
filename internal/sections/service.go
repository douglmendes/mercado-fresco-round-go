package sections

type Service interface {
	GetAll() ([]Section, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Section, error) {
	return s.repository.GetAll()
}

func NewService(r Repository) Service {
	return &service{r}
}
