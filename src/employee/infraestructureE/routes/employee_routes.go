package routes

import (
	controllers "Abarrotes/src/employee/infraestructureE/controller"

	"github.com/gin-gonic/gin"
)


func RegisterEmployeeRoutes(router *gin.Engine, addEmployeeController *controllers.AddEmployeeController, deleteEmployeeController *controllers.DeleteEmployeeController) {
	router.POST("/employees", addEmployeeController.Handle)
	router.DELETE("/employees/:id", deleteEmployeeController.DeleteEmployee)
	router.GET("/employees/", deleteEmployeeController.DeleteEmployee)
	router.PUT("/employees/:id", deleteEmployeeController.DeleteEmployee)

}