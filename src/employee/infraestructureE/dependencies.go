package infraestructure

import (
	"Abarrotes/src/employee/application"
	"Abarrotes/src/employee/infraestructureE/controller"
)

func Init() (*controller.AddEmployeeController, *controller.DeleteEmployeeController, *controller.ShowEmployeeController, *controller.ModifyEmployeeController) {
	es := NewMySQL()

	createEmployeeUseCase := application.NewAddEmployee(es)
	deleteEmployeeUseCase := application.NewDeleteEmployee(es)
	getEmployeeUseCase := application.NewShowEmployee(es)
	modifyEmployeeUseCase := application.NewModifyEmployee(es)

	
	
	
	createEmployeeController := controller.NewAddEmployeeController(createEmployeeUseCase)
	deleteEmployeeController := controller.NewDeleteEmployeeController(deleteEmployeeUseCase)
	getEmployeeController := controller.NewShowEmployeeController(*getEmployeeUseCase)
	modifyEmployeeController := controller.NewModifyEmployeeController(*modifyEmployeeUseCase)
	

	return createEmployeeController, deleteEmployeeController, getEmployeeController, modifyEmployeeController
}