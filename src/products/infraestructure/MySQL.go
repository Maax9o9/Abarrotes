package infraestructure

import (
	"Abarrotes/src/core"
	"Abarrotes/src/products/domain"
	"Abarrotes/src/products/domain/entities"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() domain.ProductRepository {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Create(product entities.Product) (entities.Product, error) {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price)
	if err != nil {
		log.Printf("Error al insertar producto: %v", err)
		return entities.Product{}, err
	}

	lastInsertID, _ := result.LastInsertId()
	product.ID = int(lastInsertID)

	return product, nil
}

func (mysql *MySQL) GetAll() ([]entities.Product, error) {
	query := "SELECT idproduct, name, price FROM product"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Printf("Error al escanear producto: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
