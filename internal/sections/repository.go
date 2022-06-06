package sections

var database []Section

type Repository interface {
	GetAll() ([]Section, error)
}

type repository struct {
}

func (r *repository) GetAll() ([]Section, error) {
	return database, nil
}

func NewRepository() Repository {
	return &repository{}
}
