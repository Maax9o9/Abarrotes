package routes

import (
	controllers "Abarrotes/src/order/infraestructureO/controller"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(
	router *gin.Engine,
	createOrderController *controllers.CreateOrderController,
	deleteOrderController *controllers.DeleteOrderController,
	listOrderController *controllers.ListOrderController,
	updateOrderController *controllers.UpdateOrderController,
	getOrderStatusController *controllers.GetOrderStatusController,
	waitForOrderUpdateController *controllers.WaitForOrderUpdateController,
	
) {
	router.POST("/orders", createOrderController.Handle)
	router.DELETE("/orders/:id", deleteOrderController.Handle)
	router.GET("/orders", listOrderController.Handle)
	router.PATCH("/orders/:id", updateOrderController.Handle)
	router.GET("/orders/:id/status", getOrderStatusController.Handle)
	router.GET("/orders/:id/wait", waitForOrderUpdateController.Handle)
}
