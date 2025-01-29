package main

import (
	"Abarrotes/src/product/infrastructureP"
	"Abarrotes/src/product/infrastructureP/routes"
	"Abarrotes/src/employee/infrastructureE"
	"Abarrotes/src/employee/infrastructureE/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	showController, createController,removeController,modifyController := infrastructureP.Init()
	routes.RegisterProductRoutes(router, showController, createController,removeController,modifyController)
	
	addEmployeeController, deleteEmployeeController, ShowEmployeeController := infrastructureE.Init()
	routes.RegisterEmployeeRoutes(router, addEmployeeController, deleteEmployeeController, ShowEmployeeController)

	router.Run(":8080")
}
