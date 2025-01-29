package routes

import (
	controllers "Abarrotes/src/products/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, showController *controllers.ShowProductController, createController *controllers.CreateProductController) {
	router.GET("/products", showController.Handle)
	router.POST("/products", createController.Handle)
}
