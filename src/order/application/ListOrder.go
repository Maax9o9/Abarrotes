package application

import (
	"Abarrotes/src/order/domain"
	"Abarrotes/src/order/domain/entities"
)

type ListOrder struct {
	repository domain.OrderRepository
}

func NewListOrder(repo domain.OrderRepository) *ListOrder {
	return &ListOrder{repository: repo}
}

func (uc *ListOrder) Execute() ([]entities.Order, error) {
	return uc.repository.List()
}
