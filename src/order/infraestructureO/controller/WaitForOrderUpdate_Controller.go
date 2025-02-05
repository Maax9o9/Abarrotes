package controller

import (
	"Abarrotes/src/order/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WaitForOrderUpdateController struct {
	useCase *application.WaitForOrderUpdate
}

func NewWaitForOrderUpdateController(uc *application.WaitForOrderUpdate) *WaitForOrderUpdateController {
	return &WaitForOrderUpdateController{useCase: uc}
}

func (c *WaitForOrderUpdateController) Handle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order id"})
		return
	}

	currentStatus := ctx.Query("current")
	if currentStatus == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'current' (current status) is required"})
		return
	}

	timeoutSeconds := 30 
	if timeoutQuery := ctx.Query("timeout"); timeoutQuery != "" {
		if t, err := strconv.Atoi(timeoutQuery); err == nil {
			timeoutSeconds = t
		}
	}

	newStatus, err := c.useCase.Execute(orderID, currentStatus, timeoutSeconds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": newStatus})
}
