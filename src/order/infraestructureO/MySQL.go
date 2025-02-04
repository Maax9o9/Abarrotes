package infraestructureo

import (
	"Abarrotes/src/core"
	"Abarrotes/src/order/domain"
	"Abarrotes/src/order/domain/entities"
	"log"
	"time"
)

// MySQL es la implementación de domain.OrderRepository para MySQL.
type MySQL struct {
	conn *core.Conn_MySQL
}

// NewMySQL crea una nueva instancia del repositorio MySQL para Order.
func NewMySQL() domain.OrderRepository {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

// Create inserta una nueva orden en la base de datos.
func (mysql *MySQL) Create(order entities.Order) (entities.Order, error) {
	query := "INSERT INTO orders (order_date, status) VALUES (?, ?)"
	// Si no se estableció la fecha, asignamos la fecha actual.
	if order.OrderDate.IsZero() {
		order.OrderDate = time.Now()
	}
	result, err := mysql.conn.ExecutePreparedQuery(query, order.OrderDate, order.Status)
	if err != nil {
		log.Printf("Error al insertar order: %v", err)
		return entities.Order{}, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID insertado: %v", err)
		return entities.Order{}, err
	}
	order.ID = int(lastInsertID)
	return order, nil
}

// Delete elimina una orden de la base de datos dado su ID.
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM orders WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar order con ID %d: %v", id, err)
		return err
	}

	return nil
}

// GetByID obtiene una orden según su ID.
func (mysql *MySQL) GetByID(id int) (entities.Order, error) {
	query := "SELECT id, order_date, status FROM orders WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var order entities.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.OrderDate, &order.Status); err != nil {
			log.Printf("Error al escanear order: %v", err)
			return entities.Order{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return entities.Order{}, err
	}

	return order, nil
}

// GetAll retorna todas las órdenes almacenadas en la base de datos.
func (mysql *MySQL) GetAll() ([]entities.Order, error) {
	query := "SELECT id, order_date, status FROM orders"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(&order.ID, &order.OrderDate, &order.Status); err != nil {
			log.Printf("Error al escanear order: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}
// Update modifica una orden existente en la base de datos.
func (mysql *MySQL) Update(order entities.Order) (entities.Order, error) {
	query := "UPDATE orders SET order_date = ?, status = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, order.OrderDate, order.Status, order.ID)
	if err != nil {
		log.Printf("Error al modificar order: %v", err)
		return entities.Order{}, err
	}

	return order, nil
}

// List returns a list of orders based on the provided filter.
func (mysql *MySQL) List(filter domain.OrderFilter) ([]entities.Order, error) {
	query := "SELECT id, order_date, status FROM orders WHERE status = ?"
	rows := mysql.conn.FetchRows(query, filter.Status)
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(&order.ID, &order.OrderDate, &order.Status); err != nil {
			log.Printf("Error al escanear order: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
	_, err := mysql.conn.ExecutePreparedQuery(query, order.OrderDate, order.Status, order.ID)
	if err != nil {
		log.Printf("Error al modificar order: %v", err)
		return entities.Order{}, err
	}

	return order, nil
}
