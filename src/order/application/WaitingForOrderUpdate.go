package application

import (
	"Abarrotes/src/order/domain"
	"time"
)

type WaitForOrderUpdate struct {
	repository domain.OrderRepository
}

func NewWaitForOrderUpdate(repo domain.OrderRepository) *WaitForOrderUpdate {
	return &WaitForOrderUpdate{repository: repo}
}

func (uc *WaitForOrderUpdate) Execute(orderID int, currentStatus string, timeoutSeconds int) (string, error) {
	timeout := time.After(time.Duration(timeoutSeconds) * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return currentStatus, nil
		case <-ticker.C:
			order, err := uc.repository.GetByID(orderID)
			if err != nil {
				return "", err
			}
			if order.Status != currentStatus {
				return order.Status, nil
			}
		}
	}
}
