package buyers

import (
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetById(id int) (Buyer, error)
	GetAll() ([]Buyer, error)
	LastID() (int, error)
	Create(id int, cardNumberId, firstName, lastName string) (Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (Buyer, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Buyer, error) {
	var by []Buyer
	if err := r.db.Read(&by); err != nil {
		return []Buyer{}, nil
	}
	return by, nil
}

func (r *repository) GetById(id int) (Buyer, error) {
	var by []Buyer

	if err := r.db.Read(&by); err != nil {
		return Buyer{}, nil
	}

	for i := range by {
		if by[i].Id == id {
			return by[i], nil
		}
	}

	return Buyer{}, fmt.Errorf("Buyer %d not found", id)
}

func (r *repository) LastID() (int, error) {
	var by []Buyer

	if err := r.db.Read(&by); err != nil {
		return 0, err
	}

	if len(by) == 0 {
		return 0, nil
	}

	return by[len(by)-1].Id, nil
}

func (r *repository) Create(id int, cardNumberId, firstName, lastName string) (Buyer, error) {
	var by []Buyer

	if err := r.db.Read(&by); err != nil {
		return Buyer{}, err
	}
	s := Buyer{id, cardNumberId, firstName, lastName}
	by = append(by, s)
	if err := r.db.Write(by); err != nil {
		return Buyer{}, err
	}
	return s, nil
}

func (r *repository) Update(id int, cardNumberId, firstName, lastName string) (Buyer, error) {
	var by []Buyer

	if err := r.db.Read(&by); err != nil {
		return Buyer{}, nil
	}

	s := Buyer{}

	updated := false
	for i := range by {
		if by[i].Id == id {
			s = by[i]
			if cardNumberId != "" {
				s.CardNumberId = cardNumberId
			}
			if firstName != "" {
				s.FirstName = firstName
			}
			if lastName != "" {
				s.LastName = lastName
			}

			by[i] = s
			updated = true

			if err := r.db.Write(by); err != nil {
				return Buyer{}, err
			}
		}
	}

	if !updated {
		return Buyer{}, fmt.Errorf("Buyer%d not found", id)
	}
	return s, nil

}

func (r *repository) Delete(id int) error {
	var by []Buyer

	if err := r.db.Read(&by); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range by {
		if by[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Buyer %d not found", id)
	}

	by = append(by[:index], by[index+1:]...)
	if err := r.db.Write(by); err != nil {
		return err
	}

	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
