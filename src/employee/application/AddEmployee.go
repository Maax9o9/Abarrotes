package application

import (
	"Abarrotes/src/employee/domain"
	"Abarrotes/src/employee/domain/entities"
)

type AddEmployee struct {
	repository domain.EmployeeRepository
}

func NewAddEmployee(repo domain.EmployeeRepository) *AddEmployee {
	return &AddEmployee{repository: repo}
}

func (uc *AddEmployee) Execute(employee entities.Employee) (entities.Employee, error) {
	return uc.repository.Add(employee)
}