package controller

import (
	"Abarrotes/src/order/application"
	"Abarrotes/src/order/domain/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateOrderController struct {
	useCase *application.UpdateOrder
}

func NewUpdateOrderController(uc *application.UpdateOrder) *UpdateOrderController {
	return &UpdateOrderController{useCase: uc}
}

func (c *UpdateOrderController) Handle(ctx *gin.Context) {
	var orderToUpdate entities.Order

	if err := ctx.ShouldBindJSON(&orderToUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Printf("Error al bindear JSON: %v", err)
		return
	}

	log.Printf("Datos de la orden recibidos en el controlador: %+v", orderToUpdate)

	order, err := c.useCase.Execute(orderToUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("Error al ejecutar el caso de uso: %v", err)
		return
	}

	log.Printf("Orden actualizada: %+v", order)

	ctx.JSON(http.StatusOK, order)
}
