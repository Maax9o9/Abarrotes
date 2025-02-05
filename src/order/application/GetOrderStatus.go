package application

import (
	"Abarrotes/src/order/domain"
)

type GetOrderStatus struct {
	repository domain.OrderRepository
}

func NewGetOrderStatus(repo domain.OrderRepository) *GetOrderStatus {
	return &GetOrderStatus{repository: repo}
}

func (uc *GetOrderStatus) Execute(orderID int) (string, error) {
	order, err := uc.repository.GetByID(orderID)
	if err != nil {
		return "", err
	}
	return order.Status, nil
}
