package infraestructure

import (
	"Abarrotes/src/core"
	"Abarrotes/src/order/domain"
	"Abarrotes/src/order/domain/entities"
	"log"
	"time"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func (mysql *MySQL) List() ([]entities.Order, error) {
	return mysql.GetAll()
}

func NewMySQL() domain.OrderRepository {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Create(order entities.Order) (entities.Order, error) {
	query := "INSERT INTO orders (order_date, status) VALUES (?, ?)"
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

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM orders WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar order con ID %d: %v", id, err)
		return err
	}
	return nil
}

func (mysql *MySQL) GetByID(id int) (entities.Order, error) {
	query := "SELECT id, order_date, status FROM orders WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var order entities.Order
	var rawDate []byte

	for rows.Next() {
		if err := rows.Scan(&order.ID, &rawDate, &order.Status); err != nil {
			log.Printf("Error al escanear order: %v", err)
			return entities.Order{}, err
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(rawDate))
		if err != nil {
			log.Printf("Error al parsear order_date: %v", err)
			return entities.Order{}, err
		}
		order.OrderDate = parsedTime
	}

	if err := rows.Err(); err != nil {
		return entities.Order{}, err
	}

	return order, nil
}

func (mysql *MySQL) GetAll() ([]entities.Order, error) {
	query := "SELECT id, order_date, status FROM orders"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var orders []entities.Order

	for rows.Next() {
		var order entities.Order
		var rawDate []byte

		if err := rows.Scan(&order.ID, &rawDate, &order.Status); err != nil {
			log.Printf("Error al escanear order: %v", err)
			return nil, err
		}

		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(rawDate))
		if err != nil {
			log.Printf("Error al parsear order_date: %v", err)
			return nil, err
		}
		order.OrderDate = parsedTime

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (mysql *MySQL) Update(order entities.Order) (entities.Order, error) {
	query := "UPDATE orders SET order_date = ?, status = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, order.OrderDate, order.Status, order.ID)
	if err != nil {
		log.Printf("Error al modificar order: %v", err)
		return entities.Order{}, err
	}

	return order, nil
}
