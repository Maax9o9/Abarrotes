package controller

import (
	"Abarrotes/src/order/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetOrderStatusController struct {
	useCase *application.GetOrderStatus
}

func NewGetOrderStatusController(uc *application.GetOrderStatus) *GetOrderStatusController {
	return &GetOrderStatusController{useCase: uc}
}

func (c *GetOrderStatusController) Handle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id"})
		return
	}

	status, err := c.useCase.Execute(orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": status})
}
