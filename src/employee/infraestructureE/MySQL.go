package infraestructure

import (
	"Abarrotes/src/core"
	"Abarrotes/src/employee/domain"
	"Abarrotes/src/employee/domain/entities"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func (mysql *MySQL) DeleteEmployee(employeeID int) error {
	query := "DELETE FROM employee WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, employeeID)
	if err != nil {
		log.Printf("Error al eliminar empleado con ID %d: %v", employeeID, err)
		return err
	}

	return nil
}


func (mysql *MySQL) Add(employee entities.Employee) (entities.Employee, error) {
    return mysql.Create(employee)
}

func NewMySQL() domain.EmployeeRepository {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Create(employee entities.Employee) (entities.Employee, error) {
	query := "INSERT INTO employee (name, last_name, age, job_position) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, employee.Name, employee.LastName, employee.Age, employee.JobPosition)
	if err != nil {
		log.Printf("Error al insertar empleado: %v", err)
		return entities.Employee{}, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID insertado: %v", err)
		return entities.Employee{}, err
	}
	employee.ID = int(lastInsertID)

	return employee, nil
}


func (mysql *MySQL) GetByID(id int) (entities.Employee, error) {
	query := "SELECT id, name, last_name, age, job_position FROM employee WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var employee entities.Employee
	for rows.Next() {
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.LastName, &employee.Age, &employee.JobPosition); err != nil {
			log.Printf("Error al escanear empleado: %v", err)
			return entities.Employee{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return entities.Employee{}, err
	}

	return employee, nil
}

func (mysql *MySQL) GetAll() ([]entities.Employee, error) {
	query := "SELECT id, name, last_name, age, job_position FROM employee"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.LastName, &employee.Age, &employee.JobPosition); err != nil {
			log.Printf("Error al escanear empleado: %v", err)
			return []entities.Employee{}, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return []entities.Employee{}, err
	}

	return employees, nil
}

func (mysql *MySQL) ModifyEmployee(employee entities.Employee) error {
	query := "UPDATE employee SET name = ?, last_name = ?, age = ?, job_position = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, employee.Name, employee.LastName, employee.Age, employee.JobPosition, employee.ID)
	if err != nil {
		log.Printf("Error al modificar empleado: %v", err)
		return err
	}

	return nil
}
