package sellers

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]domain.Seller, error) {
	var sl []domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return []domain.Seller{}, nil
	}
	return sl, nil
}

func (r *repository) GetById(id int) (domain.Seller, error) {
	var sl []domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return domain.Seller{}, nil
	}

	for i := range sl {
		if sl[i].ID == id {
			return sl[i], nil
		}
	}

	return domain.Seller{}, fmt.Errorf("seller %d not found", id)
}

func (r *repository) LastID() (int, error) {
	var sl []domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return 0, err
	}

	if len(sl) == 0 {
		return 0, nil
	}

	return sl[len(sl)-1].ID, nil
}

func (r *repository) Create(id, cid int, commpanyName, address, telephone string) (domain.Seller, error) {
	var sl []domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return domain.Seller{}, err
	}
	s := domain.Seller{id, cid, commpanyName, address, telephone}
	sl = append(sl, s)
	if err := r.db.Write(sl); err != nil {
		return domain.Seller{}, err
	}
	return s, nil
}

func (r *repository) Update(id, cid int, commpanyName, address, telephone string) (domain.Seller, error) {
	var sl []domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return domain.Seller{}, err
	}

	s := domain.Seller{}

	updated := false
	for i := range sl {
		if sl[i].ID == id {
			s = sl[i]
			if cid != 0 {
				s.Cid = cid
			}
			if commpanyName != "" {
				s.CompanyName = commpanyName
			}
			if address != "" {
				s.Address = address
			}
			if telephone != "" {
				s.Telephone = telephone
			}

			sl[i] = s
			updated = true
			if err := r.db.Write(sl); err != nil {
				return domain.Seller{}, err
			}
		}
	}

	if !updated {
		return domain.Seller{}, fmt.Errorf("seller %d not found", id)
	}
	return s, nil

}

func (r *repository) Delete(id int) error {
	var sl []domain.Seller
	if err := r.db.Read(&sl); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range sl {
		if sl[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("seller %d not found", id)
	}

	sl = append(sl[:index], sl[index+1:]...)
	if err := r.db.Write(sl); err != nil {
		return err
	}

	return nil
}

func NewRepository(db store.Store) domain.Repository {
	return &repository{
		db: db,
	}

}
