package application

import (
	"Abarrotes/src/order/domain"
)

type DeleteOrder struct {
	repository domain.OrderRepository
}

func NewDeleteOrder(repo domain.OrderRepository) *DeleteOrder {
	return &DeleteOrder{repository: repo}
}

func (uc *DeleteOrder) Execute(id int) error {
	return uc.repository.Delete(id)
}
