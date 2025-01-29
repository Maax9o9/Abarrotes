package infraestructure

import (
	"Abarrotes/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}


func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}


func (mysql *MySQL) Save(name string, price float64) {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, price)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Producto agregado correctamente")
	}
}


func (mysql *MySQL) GetAll() {
	query := "SELECT idproduct, name, price FROM product"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	for rows.Next() {
		var idproduct int
		var name string
		var price float32
		if err := rows.Scan(&idproduct, &name, &price); err != nil {
			log.Fatalf("Error al escanear la fila: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", idproduct, name, price)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterando sobre las filas: %v", err)
	}
}
