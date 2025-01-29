package controllers

import (
	"Abarrotes/src/products/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ShowProductController struct {
	useCase *application.ShowProduct
}

func NewShowProductController(uc *application.ShowProduct) *ShowProductController {
	return &ShowProductController{useCase: uc}
}

func (c *ShowProductController) Handle(ctx *gin.Context) {
	products, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
