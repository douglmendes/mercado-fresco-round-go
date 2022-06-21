package products

import (
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Create(arg Product) (Product, error)
	LastID() (int, error)
	Update(arg Product) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product

	err := r.db.Read(&products)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var products []Product

	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}

	return Product{}, fmt.Errorf("product (%d) not found", id)
}

func (r *repository) Create(arg Product) (Product, error) {
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return Product{}, err
	}

	products = append(products, arg)

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return arg, nil
}

func (r *repository) LastID() (int, error) {
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

func (r *repository) Update(arg Product) (Product, error) {
	var products []Product
	updated := false

	if err := r.db.Read(&products); err != nil {
		return Product{}, err
	}

	for index, product := range products {
		if product.Id == arg.Id {
			products[index] = arg
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("product (%d) not found", arg.Id)
	}

	r.db.Write(products)

	return arg, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	foundIndex := 0
	var products []Product

	if err := r.db.Read(&products); err != nil {
		return err
	}

	for index, product := range products {
		if product.Id == id {
			deleted = true
			foundIndex = index
		}
	}

	if !deleted {
		return fmt.Errorf("product (%d) not found", id)
	}

	products = append(products[:foundIndex], products[foundIndex+1:]...)

	r.db.Write(products)

	return nil
}
