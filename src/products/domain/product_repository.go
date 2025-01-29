package domain

import "Abarrotes/src/products/domain/entities"

type ProductRepository interface {
	GetAll() ([]entities.Product, error)
	Create(product entities.Product) (entities.Product, error)
}
