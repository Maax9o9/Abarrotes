package application

import (
	"Abarrotes/src/order/domain"
	"Abarrotes/src/order/domain/entities"
)

type CreateOrder struct {
	repository domain.OrderRepository
}

func NewCreateOrder(repo domain.OrderRepository) *CreateOrder {
	return &CreateOrder{repository: repo}
}

func (uc *CreateOrder) Execute(order entities.Order) (entities.Order, error) {
	return uc.repository.Create(order)
}
