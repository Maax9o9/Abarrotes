package application

import (
	"Abarrotes/src/employee/domain"
	"Abarrotes/src/employee/domain/entities"
)

type ShowEmployee struct {
	employeeStorage domain.EmployeeRepository
}

func NewShowEmployee(es domain.EmployeeRepository) *ShowEmployee {
	return &ShowEmployee{employeeStorage: es}
}

func (se *ShowEmployee) Execute() ([]entities.Employee, error) {
	return se.employeeStorage.GetAll()
}