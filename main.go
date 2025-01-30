package main

import (
	infr "Abarrotes/src/employee/infraestructureE"
	infraestructure "Abarrotes/src/products/infraestructureP"
	routesE "Abarrotes/src/employee/infraestructureE/routes"
	routesP "Abarrotes/src/products/infraestructureP/routes"
	

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	showController, createController, removeController, modifyController := infraestructure.Init()
	routesP.RegisterProductRoutes(router, showController, createController, removeController, modifyController)

	addEmployeeController, deleteEmployeeController, ShowEmployeeController, modifyEmployeeController := infr.Init()
	routesE.RegisterEmployeeRoutes(router, addEmployeeController, deleteEmployeeController, ShowEmployeeController, modifyEmployeeController)

	router.Run(":8080")
}
