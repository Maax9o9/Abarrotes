package controllers

import (
	"Abarrotes/src/products/application"
	"Abarrotes/src/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase *application.CreateProduct
}

func (c *CreateProductController) Execute() {
	panic("unimplemented")
}

func NewCreateProductController(uc *application.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: uc}
}

func (c *CreateProductController) Handle(ctx *gin.Context) {
	var newProduct entities.Product
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product, err := c.useCase.Execute(newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
