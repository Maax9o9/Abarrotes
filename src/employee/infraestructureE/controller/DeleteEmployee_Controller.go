package controller

import (
	"Abarrotes/src/employee/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteEmployeeController struct {
	deleteEmployeeUseCase *application.DeleteEmployee
}

func NewDeleteEmployeeController(deleteEmployeeUseCase *application.DeleteEmployee) *DeleteEmployeeController {
	return &DeleteEmployeeController{deleteEmployeeUseCase: deleteEmployeeUseCase}
}

func (dec *DeleteEmployeeController) DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := dec.deleteEmployeeUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}