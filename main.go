package main

import (
	"Abarrotes/src/products/infraestructure"
	"Abarrotes/src/products/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	showController, createController := infraestructure.Init()
	routes.RegisterProductRoutes(router, showController, createController)
	router.Run(":8080")
}
