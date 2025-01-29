package controllers

import (
	"Abarrotes/src/products/application"
	"Abarrotes/src/products/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ModifyProductController struct {
	useCase *application.ModifyProduct
}

func NewModifyProductController(useCase *application.ModifyProduct) *ModifyProductController {
	return &ModifyProductController{useCase: useCase}
}

func (ctrl *ModifyProductController) Execute(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = ctrl.useCase.Execute(productID, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto modificado correctamente"})
}
