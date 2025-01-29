package controllers

import (
	"Abarrotes/src/products/application"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemoveProductController struct {
	useCase *application.RemoveProduct
}

func NewRemoveProductController(useCase *application.RemoveProduct) *RemoveProductController {
	return &RemoveProductController{useCase: useCase}
}

func (ctrl *RemoveProductController) Execute(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	fmt.Println("Intentando eliminar producto con ID:", productID) 

	err = ctrl.useCase.Execute(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}

