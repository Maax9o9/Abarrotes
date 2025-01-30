package controller

import (
	"Abarrotes/src/employee/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowEmployeeController struct {
	showEmployee application.ShowEmployee
}

func NewShowEmployeeController(se application.ShowEmployee) *ShowEmployeeController {
	return &ShowEmployeeController{showEmployee: se}
}

func (sec *ShowEmployeeController) Handle(ctx *gin.Context) {
	products, err := sec.showEmployee.Execute()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}	

