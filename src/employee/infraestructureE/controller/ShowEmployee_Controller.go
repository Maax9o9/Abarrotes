package controller

import (
	"Abarrotes/src/employee/application"
	"Abarrotes/src/employee/domain/entities"
)

type ShowEmployeeController struct {
	showEmployee application.ShowEmployee
}

func NewShowEmployeeController(se application.ShowEmployee) *ShowEmployeeController {
	return &ShowEmployeeController{showEmployee: se}
}

func (sec *ShowEmployeeController) Execute() (entities.Employee, error) {
	employees, err := sec.showEmployee.Execute()
	if err != nil {
		return entities.Employee{}, err
	}
	if len(employees) == 0 {
		return entities.Employee{}, nil
	}
	return employees[0], nil
}