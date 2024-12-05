package repository

import (
	"github.com/google/uuid"
	"github.com/maxwelbm/pod_example/internal/service/model"
)

type RepositoryDB struct {
	DB map[string]*model.Product
}

func (r *RepositoryDB) Create(product model.Product) (model.Product, error) {
	id := uuid.New()
	r.DB[id.String()] = &product
	product.Id = product.Id
	return product, nil
}

func NewMeliDB(ok bool) RepositoryDB {
	return RepositoryDB{
		DB: make(map[string]*model.Product),
	}
}
