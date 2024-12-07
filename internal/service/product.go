package service

import (
	"github.com/maxwelbm/pod_example/internal/model"
	"github.com/maxwelbm/pod_example/internal/repository"
)

type ServiceProducts struct {
	Repo repository.Repository
}

func (s *ServiceProducts) Create(product model.Product) (model.Product, error) {
	product, err := s.Repo.Create(product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (s *ServiceProducts) GetAll() ([]*model.Product, error) {
	return s.Repo.GetAll()
}

func (s *ServiceProducts) GetID(id int) (*model.Product, error) {
	return s.Repo.GetID(id)
}

func (s *ServiceProducts) GetSearch(price float64) ([]*model.Product, error) {
	return s.Repo.GetSearch(price)
}

func NewServiceProducts(repo repository.Repository) ServiceProducts {
	return ServiceProducts{
		Repo: repo,
	}
}
