package controller

import (
	"Abarrotes/src/order/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteOrderController struct {
	useCase *application.DeleteOrder
}

func NewDeleteOrderController(uc *application.DeleteOrder) *DeleteOrderController {
	return &DeleteOrderController{useCase: uc}
}

func (c *DeleteOrderController) Handle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id"})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
