package application

import (
	"Abarrotes/src/employee/domain"
	"Abarrotes/src/employee/domain/entities"
)


type ModifyEmployee struct {
	employeeRepository domain.EmployeeRepository
}

func NewModifyEmployee(er domain.EmployeeRepository) *ModifyEmployee {
	return &ModifyEmployee{
		employeeRepository: er,
	}
}

func (me *ModifyEmployee) ModifyEmployee(e entities.Employee) error {
	return me.employeeRepository.ModifyEmployee(e)
}
