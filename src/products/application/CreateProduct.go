package application

import (
	"Abarrotes/src/products/domain"
	"Abarrotes/src/products/domain/entities"
)

type CreateProduct struct {
	repository domain.ProductRepository
}

func NewCreateProduct(repo domain.ProductRepository) *CreateProduct {
	return &CreateProduct{repository: repo}
}

func (uc *CreateProduct) Execute(product entities.Product) (entities.Product, error) {
	return uc.repository.Create(product)
}
