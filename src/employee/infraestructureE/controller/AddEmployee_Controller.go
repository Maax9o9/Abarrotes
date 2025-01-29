package controller

import (
	"Abarrotes/src/employee/application"
	"Abarrotes/src/employee/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddEmployeeController struct {
	useCase *application.AddEmployee
}

func (c *AddEmployeeController) Execute() {
	panic("unimplemented")
}

func NewAddEmployeeController(uc *application.AddEmployee) *AddEmployeeController {
	return &AddEmployeeController{useCase: uc}
}

func (c *AddEmployeeController) Handle(ctx *gin.Context) {
	var newEmployee entities.Employee
	if err := ctx.ShouldBindJSON(&newEmployee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	employee, err := c.useCase.Execute(newEmployee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, employee)
}