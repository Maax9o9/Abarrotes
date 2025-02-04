package infraestructureo

import (
	"Abarrotes/src/order/application"
	"Abarrotes/src/order/infraestructureO/controller"
)


func Init() (*controller.CreateOrderController, *controller.DeleteOrderController, *controller.ListOrderController,*controller.UpdateOrderController) {
	
	or := NewMySQL()

	createOrderUseCase := application.NewCreateOrder(or)
	deleteOrderUseCase := application.NewDeleteOrder(or)
	listOrderUseCase := application.NewListOrder(or)
	updateOrderUseCase := application.NewUpdateOrder(or)

	createOrderController := controller.NewCreateOrderController(createOrderUseCase)
	deleteOrderController := controller.NewDeleteOrderController(deleteOrderUseCase)
	listOrderController := controller.NewListOrderController(listOrderUseCase)
	updateOrderController := controller.NewUpdateOrderController(updateOrderUseCase)

	return createOrderController, deleteOrderController, listOrderController, updateOrderController
}
