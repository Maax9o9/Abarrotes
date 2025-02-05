package application

import (
	"Abarrotes/src/order/domain"
	"Abarrotes/src/order/domain/entities"
)

type UpdateOrder struct {
	repository domain.OrderRepository
}

func NewUpdateOrder(repo domain.OrderRepository) *UpdateOrder {
	return &UpdateOrder{repository: repo}
}

func (uc *UpdateOrder) Execute(order entities.Order) (entities.Order, error) {
	return uc.repository.Update(order)
}
