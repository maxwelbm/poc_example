package repository

import (
	"github.com/maxwelbm/pod_example/internal/model"
	pkgErr "github.com/maxwelbm/pod_example/pkg/error"
)

type RepositoryDB struct {
	DB map[int]*model.Product
}

func (r *RepositoryDB) Create(product model.Product) (model.Product, error) {
	id := len(r.DB) + 1

	product.ID = id

	r.DB[int(id)] = &product
	return product, nil
}

func (r *RepositoryDB) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	for _, v := range r.DB {
		products = append(products, v)

	}

	return products, nil
}

func (r *RepositoryDB) GetID(id int) (*model.Product, error) {
	product, ok := r.DB[id]
	if !ok {
		return nil, pkgErr.ErrorNotFound
	}

	return product, nil
}

func (r *RepositoryDB) GetSearch(price float64) ([]*model.Product, error) {
	var products []*model.Product
	for _, v := range r.DB {
		if v.Price > price {
			products = append(products, v)
		}
	}

	if len(products) == 0 {
		return nil, pkgErr.ErrorNotFound
	}

	return products, nil
}

func NewMeliDB() RepositoryDB {
	return RepositoryDB{
		DB: make(map[int]*model.Product),
	}
}
