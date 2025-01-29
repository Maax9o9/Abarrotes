package controller

import (
	"Abarrotes/src/employee/application"
	"Abarrotes/src/employee/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ModifyEmployeeController struct {
	modifyEmployee application.ModifyEmployee
}

func NewModifyEmployeeController(er application.ModifyEmployee) *ModifyEmployeeController {
	return &ModifyEmployeeController{modifyEmployee: er}
}

func (mec *ModifyEmployeeController) ModifyEmployee(c *gin.Context) {
	var employee entities.Employee
	employee.ID, _ = strconv.Atoi(c.Param("id"))
	c.BindJSON(&employee)
	err := mec.modifyEmployee.ModifyEmployee(employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee modified successfully"})
}