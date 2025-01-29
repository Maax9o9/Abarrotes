package main

import (
	"Abarrotes/src/products/infraestructure/routes"
	"Abarrotes/src/products/infraestructure/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	showController, createController := products.InitializeDependencies()
	routes.RegisterProductRoutes(router, showController, createController)

	router.Run(":8080")
}
