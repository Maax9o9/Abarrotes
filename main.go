package main

import (
	"Abarrotes/src/products/infraestructure"
	"Abarrotes/src/products/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	showController, createController,removeController,modifyController := infraestructure.Init()
	routes.RegisterProductRoutes(router, showController, createController,removeController,modifyController)
	router.Run(":8080")
}
