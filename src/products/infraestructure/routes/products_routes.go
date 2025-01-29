package routes

import (
	controllers "Abarrotes/src/products/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, showController *controllers.ShowProductController, createController *controllers.CreateProductController, removeController *controllers.RemoveProductController, modifyController *controllers.ModifyProductController) {
	router.GET("/products", showController.Handle)
	router.POST("/products", createController.Handle)
	router.DELETE("/products/:id", removeController.Execute)
	router.PUT("/:id", modifyController.Execute)
}
