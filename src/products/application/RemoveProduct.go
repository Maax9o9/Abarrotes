package application

import (
	"Abarrotes/src/products/domain"
	"fmt"
)

type RemoveProduct struct {
	repository domain.ProductRepository
}

func NewRemoveProduct(repo domain.ProductRepository) *RemoveProduct {
	return &RemoveProduct{repository: repo}
}

func (uc *RemoveProduct) Execute(productID int) error {
	fmt.Println("Ejecutando eliminación del producto con ID:", productID) 

	err := uc.repository.RemoveProduct(productID)
	if err != nil {
		fmt.Println("Error al eliminar producto:", err)
		return err
	}

	fmt.Println("Producto eliminado correctamente")
	return nil
}
