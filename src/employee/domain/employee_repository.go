package domain

import "Abarrotes/src/employee/domain/entities"

type EmployeeRepository interface {
	Add(employee entities.Employee) (entities.Employee, error)
	DeleteEmployee(employeeID int) error
	GetByID(id int) (entities.Employee, error)
	GetAll() ([]entities.Employee, error)
	ModifyEmployee(employee entities.Employee) error
}

