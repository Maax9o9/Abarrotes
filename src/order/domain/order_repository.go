package domain

import (
	"Abarrotes/src/order/domain/entities"
)

type OrderRepository interface {
	Create(order entities.Order) (entities.Order, error)
	Update(order entities.Order) (entities.Order, error)
	Delete(id int) error
	List() ([]entities.Order, error)
	GetByID(id int) (entities.Order, error)
}