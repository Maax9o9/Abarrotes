package application

import (
	"Abarrotes/src/products/domain"
	"Abarrotes/src/products/domain/entities"
)

type ShowProduct struct {
	repository domain.ProductRepository
}

func NewShowProduct(repo domain.ProductRepository) *ShowProduct {
	return &ShowProduct{repository: repo}
}

func (uc *ShowProduct) Execute() ([]entities.Product, error) {
	return uc.repository.GetAll()
}
