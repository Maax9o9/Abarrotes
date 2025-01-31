package domain

import "Abarrotes/src/products/domain/entities"

type ProductRepository interface {
	Create(product entities.Product) (entities.Product, error)
	GetAll() ([]entities.Product, error)
	GetByID(id int) (entities.Product, error)
	Update(id int, product entities.Product) error
	Delete(id int) error
}
