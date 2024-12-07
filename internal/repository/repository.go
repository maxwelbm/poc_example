package repository

import "github.com/maxwelbm/pod_example/internal/model"

type Repository interface {
	Create(product model.Product) (model.Product, error)
	GetAll() ([]*model.Product, error)
	GetID(id int) (*model.Product, error)
	GetSearch(price float64) ([]*model.Product, error)
}
