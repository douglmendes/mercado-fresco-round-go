package sellers

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

var sl []Seller = []Seller{}

type Repository interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Store(id, cid int, commpanyName, address, telephone string) (Seller, error)
	LastID() (int, error)
	Update(id, cid int, commpanyName, address, telephone string) (Seller, error)
	// UpdateAdress(id int, address string) (Seller, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Seller, error) {
	var sl []Seller
	if err := r.db.Read(&sl); err != nil {
		return []Seller{}, nil
	}
	return sl, nil
}

func (r *repository) GetById(id int) (Seller, error) {
	var sl []Seller
	if err := r.db.Read(&sl); err != nil {
		return Seller{}, nil
	}

	for i := range sl {
		if sl[i].ID == id {
				return sl[i], nil
		}
	}

	
	return Seller{}, fmt.Errorf("seller %d not found", id)
}

func (r *repository) LastID() (int, error) {
	var sl []Seller
	if err := r.db.Read(&sl); err != nil {
		return 0, err
	}
	
	if len(sl) == 0 {
		return 0, nil
	}

	return sl[len(sl)-1].ID, nil
}

func (r *repository) Store(id, cid int, commpanyName, address, telephone string) (Seller, error) {
	var sl []Seller
	if err := r.db.Read(&sl); err != nil {
		return Seller{}, err
	}
	s := Seller{id, cid, commpanyName, address, telephone}
	sl = append(sl, s)
	if err := r.db.Write(sl); err != nil {
		return Seller{}, err
	}
	return s, nil
}

func (r *repository) Update(id, cid int, commpanyName, address, telephone string) (Seller, error) {
	var sl []Seller
	if err := r.db.Read(&sl); err != nil {
		return Seller{}, nil
	}

	s := Seller{}

	updated := false
	for i := range sl {
		if sl[i].ID == id {
			// s.ID = sl[i].ID
			// sl[i] = s
			s = sl[i]
			if cid != 0 {
				fmt.Printf("cid é %v", cid)
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

			fmt.Printf("seller é %v", s)
			sl[i] = s
			updated = true

			// r.Store()
			if err := r.db.Write(sl); err != nil {
				return Seller{}, err
			}
		}
	}

	if !updated {
		return Seller{}, fmt.Errorf("seller %d not found", id)
	}
	return s, nil
	
}

// func (r *repository) Update(id, cid int, commpanyName, address, telephone string) (Seller, error) {
// 	var sl []Seller
// 	if err := r.db.Read(&sl); err != nil {
// 		return Seller{}, nil
// 	}

// 	s := Seller{Cid: cid, CompanyName: commpanyName, Address: address, Telephone: telephone}

// 	updated := false
// 	for i := range sl {
// 		if sl[i].ID == id {
// 			s.ID = id
// 			sl[i] = s
// 			updated = true
// 		}
// 	}
// 	if err := r.db.Write(sl); err != nil {
// 		return Seller{}, err
// 	}
// 	if !updated {
// 		return Seller{}, fmt.Errorf("seller %d not found", id)
// 	}
// 	return s, nil
	
// }

func (r *repository) Delete(id int) error {
	var sl []Seller
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

func NewRepository(db store.Store) Repository {
	return &repository {
		db: db,
	}
	
}