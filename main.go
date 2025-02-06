package main

import (
	infraE "Abarrotes/src/employee/infraestructureE"
	infraP "Abarrotes/src/products/infraestructureP"
	infraO "Abarrotes/src/order/infraestructureO"
	routesE "Abarrotes/src/employee/infraestructureE/routes"
	routesP "Abarrotes/src/products/infraestructureP/routes"
	routesO "Abarrotes/src/order/infraestructureO/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	
	showController, createController, removeController, modifyController := infraP.Init()
	routesP.RegisterProductRoutes(router, showController, createController, removeController, modifyController)

	addEmployeeController, deleteEmployeeController, ShowEmployeeController, modifyEmployeeController := infraE.Init()
	routesE.RegisterEmployeeRoutes(router, addEmployeeController, deleteEmployeeController, ShowEmployeeController, modifyEmployeeController)

	createOrderController, deleteOrderController, listOrderController, updateOrderController, getOrderStatusController, waitForOrderUpdateController := infraO.Init()
	routesO.RegisterOrderRoutes(router, createOrderController, deleteOrderController, listOrderController, updateOrderController, getOrderStatusController, waitForOrderUpdateController)


	router.Run(":8080")
}
