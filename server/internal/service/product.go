package service

import (
	"are.moe/testtask/internal/domain"
	"are.moe/testtask/internal/repository"
)

type ProductService interface {
	Create(input domain.ProductInput) (domain.Product, error)
	GetByID(id domain.ProductID) (domain.Product, error)
	Update(id domain.ProductID, input domain.ProductInput) (domain.Product, error)
	Delete(id domain.ProductID) error
	List(offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, domain.ProductID, error)
	FindByName(name string, offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, domain.ProductID, error)
}

type Products struct {
	repo repository.ProductRepo
}

func NewProductService(repo repository.ProductRepo) *Products {
	return &Products{repo}
}

func (s *Products) Create(input domain.ProductInput) (domain.Product, error) {
	return s.repo.Create(input)
}

func (s *Products) GetByID(id domain.ProductID) (domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *Products) Update(id domain.ProductID, input domain.ProductInput) (domain.Product, error) {
	return s.repo.Update(id, input)
}

func (s *Products) Delete(id domain.ProductID) error {
	return s.repo.Delete(id)
}

func (s *Products) List(offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, domain.ProductID, error) {
	return s.repo.List(offset, limit)
}

func (s *Products) FindByName(name string, offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, domain.ProductID, error) {
	return s.repo.FindByName(name, offset, limit)
}
