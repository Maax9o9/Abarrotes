package application

import (
	"Abarrotes/src/employee/domain"
)

type DeleteEmployee struct {
	employeeStorage domain.EmployeeRepository
}

func NewDeleteEmployee(es domain.EmployeeRepository) *DeleteEmployee {
	return &DeleteEmployee{employeeStorage: es}
}

func (de *DeleteEmployee) Execute(id int) error {
	employee, err := de.employeeStorage.GetByID(id)
	if err != nil {
		return err
	}	
	return de.employeeStorage.DeleteEmployee(employee.ID)
}