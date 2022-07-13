package jsondb

import (
	"context"
	"fmt"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) domain.ProductRepository {
	return &repository{db}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.Read(&products)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Product, error) {
	var products []domain.Product

	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, fmt.Errorf("product (%d) not found", id)
}

func (r *repository) Create(ctx context.Context, arg domain.Product) (domain.Product, error) {
	var products []domain.Product

	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	lastId, err := r.lastID()
	if err != nil {
		return domain.Product{}, err
	}

	arg.Id = lastId + 1

	products = append(products, arg)

	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}

	return arg, nil
}

func (r *repository) lastID() (int, error) {
	var products []domain.Product

	if err := r.db.Read(&products); err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

func (r *repository) Update(ctx context.Context, arg domain.Product) (domain.Product, error) {
	var products []domain.Product
	updated := false

	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	for index, product := range products {
		if product.Id == arg.Id {
			products[index] = arg
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product (%d) not found", arg.Id)
	}

	r.db.Write(products)

	return arg, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	deleted := false
	foundIndex := 0
	var products []domain.Product

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
