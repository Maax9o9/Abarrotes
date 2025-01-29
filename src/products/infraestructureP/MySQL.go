package infraestructure

import (
	"Abarrotes/src/core"
	"Abarrotes/src/products/domain"
	"Abarrotes/src/products/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func (mysql *MySQL) RemoveProduct(productID int) error {
	panic("unimplemented")
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

func (mysql *MySQL) Delete(productID int) error {
	query := "DELETE FROM product WHERE idproduct = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, productID)
	if err != nil {
		log.Printf("Error al eliminar producto con ID %d: %v", productID, err)
		return fmt.Errorf("error al eliminar producto: %v", err)
	}
	return nil
}

func (mysql *MySQL) GetByID(productID int) (entities.Product, error) {
	query := "SELECT idproduct, name, price FROM product WHERE idproduct = ?"
	rows := mysql.conn.FetchRows(query, productID)
	defer rows.Close()

	var product entities.Product
	if rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return entities.Product{}, fmt.Errorf("error al escanear producto: %v", err)
		}
		return product, nil
	}

	return entities.Product{}, fmt.Errorf("producto no encontrado")
}

func (mysql *MySQL) Update(productID int, product entities.Product) error {
	query := "UPDATE product SET name = ?, price = ? WHERE idproduct = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price, productID)
	if err != nil {
		log.Printf("Error al modificar producto con ID %d: %v", productID, err)
		return fmt.Errorf("error al modificar producto: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No se encontró producto con ID %d para modificar", productID)
		return fmt.Errorf("no se encontró el producto con ID %d", productID)
	}

	log.Printf("Producto con ID %d modificado correctamente", productID)
	return nil
}