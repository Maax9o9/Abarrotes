package controller

import (
	"Abarrotes/src/order/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListOrderController struct {
	useCase *application.ListOrder
}

func NewListOrderController(uc *application.ListOrder) *ListOrderController {
	return &ListOrderController{useCase: uc}
}

func (c *ListOrderController) Handle(ctx *gin.Context) {
	orders, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
