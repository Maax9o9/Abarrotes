package infraestructure

import (
	"Abarrotes/src/order/application"
	"Abarrotes/src/order/infraestructureO/controller"
)

func Init() (
	*controller.CreateOrderController,
	*controller.DeleteOrderController,
	*controller.ListOrderController,
	*controller.UpdateOrderController,
	*controller.GetOrderStatusController,
	*controller.WaitForOrderUpdateController,
) {
	orderRepository := NewMySQL()

	createOrderUseCase := application.NewCreateOrder(orderRepository)
	deleteOrderUseCase := application.NewDeleteOrder(orderRepository)
	listOrderUseCase := application.NewListOrder(orderRepository)
	updateOrderUseCase := application.NewUpdateOrder(orderRepository)
	getOrderStatusUseCase := application.NewGetOrderStatus(orderRepository)
	waitForOrderUpdateUseCase := application.NewWaitForOrderUpdate(orderRepository)

	createOrderController := controller.NewCreateOrderController(createOrderUseCase)
	deleteOrderController := controller.NewDeleteOrderController(deleteOrderUseCase)
	listOrderController := controller.NewListOrderController(listOrderUseCase)
	updateOrderController := controller.NewUpdateOrderController(updateOrderUseCase)
	getOrderStatusController := controller.NewGetOrderStatusController(getOrderStatusUseCase)
	waitForOrderUpdateController := controller.NewWaitForOrderUpdateController(waitForOrderUpdateUseCase)

	return createOrderController, deleteOrderController, listOrderController, updateOrderController, getOrderStatusController, waitForOrderUpdateController
}
